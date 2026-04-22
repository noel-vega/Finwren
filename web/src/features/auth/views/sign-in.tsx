import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import z from "zod"
import type { SubmitEvent } from "react"
import { Field, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

const SignInFormSchema = z.object({
  email: z.string(),
  password: z.string()
})

export function SignInView() {
  const form = useForm({
    resolver: zodResolver(SignInFormSchema)
  })

  const handleSubmit = (e: SubmitEvent) => {
    return form.handleSubmit((data) => {
      console.log(data)
    })(e)
  }

  return (
    <div>
      <h1>Sign In</h1>

      <form onSubmit={handleSubmit}>
        <Field>
          <FieldLabel>Email</FieldLabel>
          <Input type="email" />
        </Field>

        <Field>
          <FieldLabel>Password</FieldLabel>
          <Input type="password" />
        </Field>

        <Button type="submit">
          Sign In
        </Button>
      </form>
    </div>
  )
}
