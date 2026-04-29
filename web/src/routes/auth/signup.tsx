import { SignUpView } from '@/features/auth/views/signup'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/signup')({
  head: () => ({ meta: [{ title: "Sign Up" }] }),
  component: SignUpView,
})
