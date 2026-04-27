# Sign-up production readiness

The handler logic itself is solid (validation, error mapping, response shape, logging, password hashing). What's listed here is the system around it.

## Hard blockers

- [ ] **Issue a session/token in the sign-up response.** Currently a successful 201 leaves the client unauthenticated with nowhere to go but `/auth/sign-in` (which is a stub). Either return a JWT, set a session cookie, or document that sign-up + sign-in is intentionally a two-step flow.
- [ ] **Rate limit `/auth/sign-up`.** bcrypt at cost 12 is ~250–500ms of CPU per call. With no rate limit, a few hundred concurrent requests pin a core — trivial DoS vector. Per-IP token bucket at the middleware layer (or at the LB/reverse proxy).
- [ ] **Cap request body size.** Gin doesn't limit body size by default. A 100MB JSON payload will be read, decoded, and validated. Wrap `r.Body` with `http.MaxBytesReader` (typically 1MB for JSON APIs).
- [ ] **Graceful shutdown.** `r.Run()` doesn't handle SIGTERM. In-flight requests get killed mid-bcrypt or mid-DB-write when the container terminates. Use `http.Server` with `Shutdown(ctx)` driven by a signal handler.
- [ ] **TLS.** Confirm TLS termination is happening upstream (LB / reverse proxy). Passwords cross the wire — anything plaintext between client and server breaks signup.

## Important (could run internally, not publicly)

- [ ] **Email verification flow.** Anyone can sign up as anyone. No proof of email ownership = no reliable password reset, account squatting, spam accounts. Create users in `unverified` state, send a confirmation link, gate sign-in until verified.
- [ ] **Mitigate email enumeration via 409.** "User with email exists" is a perfect "is this person registered?" oracle. Standard mitigation: always respond 202 with "check your email" and let the email flow handle the existing-user case.
- [ ] **Tests.** None today. Minimum: table-driven tests for each validator branch → expected `ProblemDetailError`, service mapping (`ErrEmailExists` → conflict), and the password 72-byte boundary.
- [ ] **Real `/health` check.** Currently returns "healthy" unconditionally — won't fail when the DB is down, so the orchestrator won't take the pod out of rotation. Should ping the DB.

## Production hygiene

- [ ] **Password breach check.** 12-char minimum is fine, but NIST SP 800-63B recommends checking against breach corpora (HIBP) and common-password lists rather than complexity rules.
- [ ] **Metrics.** Logs aren't a substitute. Can't graph signup rate, p99 latency, or error rate from logs alone. Prometheus or OpenTelemetry.
- [ ] **Tune DB connection pool.** `sqlx.Connect` uses defaults for `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`. Audit against the DB's connection limit.
- [ ] **CAPTCHA / bot defense.** Sign-up is anonymous and unauthenticated — first-line bot target. Hcaptcha / Cloudflare Turnstile or similar.

## Scope decision

If this is **a personal finance app, single user, behind Cloudflare** — the handler is fine to ship. Real gaps that still matter at that scale: session/token, body size cap, graceful shutdown, real health check. Everything else is overkill for one user.

If this is **public signup, real users, real money** — none of it is ready. Focus order: session/token, rate limit, email verification, tests.
