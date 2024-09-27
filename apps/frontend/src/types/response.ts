type APIResponse<T> = SuccessResponse<T> | ErrorResponse

type DefaultResponse = {
  message: string
}

type SuccessResponse<T> = DefaultResponse & {
  data: T
}

type ErrorResponse = DefaultResponse & {
  error: string
}