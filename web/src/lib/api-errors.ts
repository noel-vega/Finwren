import type { FieldValues, Path, UseFormReturn } from "react-hook-form"
import z from "zod"

export const ProblemDetailErrorSchema = z.object({
  code: z.string(),
  detail: z.string(),
  pointer: z.string(),
  params: z.record(z.string(), z.string()).optional(),
})

export type ProblemDetailError = z.infer<typeof ProblemDetailErrorSchema>

export const ProblemDetailSchema = z.object({
  type: z.string(),
  status: z.number(),
  title: z.string(),
  detail: z.string(),
  errors: z.array(ProblemDetailErrorSchema).optional()
})

export type ProblemDetail = z.infer<typeof ProblemDetailSchema>

export class ApiError extends Error {
  constructor(public problem: ProblemDetail) {
    super(problem.detail)
    this.name = "ApiError"
  }
}


export function setFormFieldErrors<T extends FieldValues>(form: UseFormReturn<T>, errors: ProblemDetailError[]) {
  for (const err of errors) {
    const name = err.pointer.replace(/^\//, "") as Path<T>
    form.setError(name, { message: err.detail })
  }
}
