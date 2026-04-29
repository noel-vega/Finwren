import { SignInView } from '@/features/auth/views/sign-in'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/sign-in')({
  component: SignInView,
})
