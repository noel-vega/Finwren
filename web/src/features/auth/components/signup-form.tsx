import { zodResolver } from "@hookform/resolvers/zod"
import { Controller, useForm } from "react-hook-form"
import { SignUpRequestParamsSchema, useSignup } from "../api/signup"
import { ApiError, setFormFieldErrors } from "@/lib/api-errors"
import { Field, FieldError, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

export function SignupForm() {
  const form = useForm({
    resolver: zodResolver(SignUpRequestParamsSchema),
    defaultValues: {
      email: "",
      firstName: "",
      lastName: "",
      password: "",
      confirmPassword: ""
    }
  })

  const signup = useSignup({
    onError: (err) => {
      if (err instanceof ApiError) {
        if (err.problem.status === 400 && err.problem.errors) {
          setFormFieldErrors(form, err.problem.errors)
          return
        }

        if (err.problem.status === 409) {
          form.setError("email", { message: err.problem.detail })
          return
        }
        const message = err.problem.status >= 500 ? "Something went wrong. Please try again." : err.problem.detail
        form.setError("root", { message })
        return
      }

      form.setError("root", { message: `Something went wrong: ${err.message}` })
    }
  })


  const onSubmit = form.handleSubmit((data) => signup.mutate(data))

  const isDisabled = !form.formState.isValid || signup.isPending || signup.isSuccess
  return (
    <form onSubmit={onSubmit} className="space-y-4">
      {form.formState.errors.root && (
        <FieldError errors={[form.formState.errors.root]} />
      )}
      <Controller
        control={form.control}
        name="email"
        render={({ field, fieldState }) => (
          <Field data-invalid={fieldState.invalid}>
            <FieldLabel>Email</FieldLabel>
            <Input
              {...field}
              required
              type="email"
              placeholder="name@example.com"
              aria-invalid={fieldState.invalid}
              autoComplete="off"
            />
            {fieldState.invalid && (
              <FieldError errors={[fieldState.error]} />
            )}
          </Field>
        )}
      />

      <Field orientation="horizontal" className="flex items-start">
        <Controller
          control={form.control}
          name="firstName"
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel>First Name</FieldLabel>
              <Input
                {...field}
                required
                type="text"
                placeholder="John"
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && (
                <FieldError errors={[fieldState.error]} />
              )}
            </Field>
          )}
        />


        <Controller
          control={form.control}
          name="lastName"
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel>Last Name</FieldLabel>
              <Input
                {...field}
                required
                type="text"
                placeholder="Smith"
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && (
                <FieldError errors={[fieldState.error]} />
              )}
            </Field>
          )}
        />
      </Field>

      <Controller
        control={form.control}
        name="password"
        render={({ field, fieldState }) => (
          <Field data-invalid={fieldState.invalid}>
            <FieldLabel>Password</FieldLabel>
            <Input
              required
              {...field}
              type="password"
              placeholder="************"
              aria-invalid={fieldState.invalid}
            />
            {fieldState.invalid && (
              <FieldError errors={[fieldState.error]} />
            )}
          </Field>
        )}
      />

      <Controller
        control={form.control}
        name="confirmPassword"
        render={({ field, fieldState }) => (
          <Field data-invalid={fieldState.invalid}>
            <FieldLabel>Confirm Password</FieldLabel>
            <Input
              {...field}
              required
              type="password"
              placeholder="************"
              aria-invalid={fieldState.invalid}
            />
            {fieldState.invalid && (
              <FieldError errors={[fieldState.error]} />
            )}
          </Field>
        )}
      />

      <Button disabled={isDisabled} type="submit" size="lg" className="w-full">
        {signup.isPending ? "Signing up..." : "Sign up"}
      </Button>
    </form>
  )
}
