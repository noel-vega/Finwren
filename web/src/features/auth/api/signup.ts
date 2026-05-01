
import { ApiError, ProblemDetailSchema } from "@/lib/api-errors"
import { mutationOptions, useMutation, type UseMutationOptions } from "@tanstack/react-query"
import z from "zod"

export const SignUpRequestParamsSchema = z.object({
  email: z.string().min(1).max(254),
  firstName: z.string().min(1).max(100),
  lastName: z.string().min(1).max(100),
  password: z.string().min(12).max(72),
  confirmPassword: z.string().min(12)
}).refine((d) => d.password === d.confirmPassword, {
  path: ["confirmPassword"],
  message: "Passwords must match"
})

export type SignUpRequestParams = z.infer<typeof SignUpRequestParamsSchema>


export const UserSchema = z.object({
  id: z.number(),
  email: z.string(),
  firstName: z.string(),
  lastName: z.string(),
  avatar: z.string(),
  emailVerifiedAt: z.coerce.date(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
})

export type User = z.infer<typeof UserSchema>

export const SignUpResponseSchema = z.object({
  user: UserSchema,
  accessToken: z.string()
})

export type SignUpResponse = z.infer<typeof SignUpResponseSchema>


export async function signup(params: SignUpRequestParams) {
  const url = new URL("/auth/signup", import.meta.env.VITE_API_BASE_URL)
  const response = await fetch(url, {
    method: "POST",
    body: JSON.stringify(params),
    headers: {
      "Content-Type": "application/json"
    },
    credentials: "include"
  })
  const data = await response.json()
  if (response.ok) {
    return SignUpResponseSchema.parse(data)
  }

  const problem = ProblemDetailSchema.parse(data)
  throw new ApiError(problem)

}


function getUseSignupOptions() {
  return mutationOptions({
    mutationFn: signup,
  })
}

type Options = Omit<ReturnType<typeof getUseSignupOptions>, "mutationFn">

export function useSignup(options?: Options) {
  return useMutation({
    ...options,
    mutationFn: signup,
  })
}
