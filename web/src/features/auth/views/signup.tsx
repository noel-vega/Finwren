import { SignupForm } from "../components/signup-form";

export function SignupView() {
  return (
    <div className="h-dvh flex items-center justify-center">
      <div className="max-w-xl w-full space-y-6">
        <h1 className="font-medium text-xl">Sign Up</h1>
        <SignupForm />
      </div>
    </div>
  )
}
