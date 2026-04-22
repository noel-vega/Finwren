import { SignUpView } from '@/features/auth/views/sign-up'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(auth)/sign-up')({
  component: SignUpView,
})
