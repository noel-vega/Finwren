CREATE TABLE IF NOT EXISTS email_tokens (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash BYTEA NOT NULL UNIQUE,
    purpose TEXT NOT NULL CHECK(purpose IN ('verify_email', 'password_reset')),
    expires_at TIMESTAMPTZ NOT NULL,
    consumed_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX email_tokens_active_idx
ON email_tokens (user_id, purpose)
WHERE consumed_at IS NULL;
