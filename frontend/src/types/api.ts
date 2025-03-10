export type IApiRes<T> = {
  status: boolean
  message: string
  data: T
}

export type IApiError<T = unknown> = {
  error: {
    message: string
    status: number
    data?: T
  }
}
