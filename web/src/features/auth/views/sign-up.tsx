import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import z from "zod"
import type { SubmitEvent } from "react"
import { Field, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

const SignUpFormSchema = z.object({
  email: z.string(),
  firstName: z.string(),
  lastName: z.string(),
  password: z.string(),
  confirmPassword: z.string()
})

export function SignUpView() {
  const form = useForm({
    resolver: zodResolver(SignUpFormSchema)
  })

  const handleSubmit = (e: SubmitEvent) => {
    return form.handleSubmit((data) => {
      console.log(data)
    })(e)
  }

  return (
    <div>
      <h1>Sign Up</h1>

      <form onSubmit={handleSubmit}>
        <Field>
          <FieldLabel>Email</FieldLabel>
          <Input type="email" />
        </Field>

        <Field>
          <FieldLabel>First Name</FieldLabel>
          <Input type="text" />
        </Field>

        <Field>
          <FieldLabel>Last Name</FieldLabel>
          <Input type="text" />
        </Field>

        <Field>
          <FieldLabel>Password</FieldLabel>
          <Input type="password" />
        </Field>

        <Field>
          <FieldLabel>Confirm Password</FieldLabel>
          <Input type="password" />
        </Field>

        <Button type="submit">
          Sign Up
        </Button>
      </form>
    </div>
  )
}
