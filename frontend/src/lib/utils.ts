import { IApiError } from "@/types/api"
import { AxiosError } from "axios"
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const formatCurrency = (amount: number, currency: string) => {
  return new Intl.NumberFormat("de-DE", {
    style: "currency",
    currency: currency,
  }).format(amount)
}

export const calculateProgress = (current: number, goal?: number) => {
  if (!goal || goal <= 0) return 0
  const percentage = (current / goal) * 100
  return Math.min(percentage, 100)
}

export const errorMessage = (error: unknown) => {
  if (error instanceof AxiosError) {
    return {
      error: {
        message: error.response?.data?.message || "Something went wrong",
        status: error.response?.status || 500,
        data: error.response?.data?.data,
      },
    } as IApiError
  }

  return {
    error: {
      message: "Something went wrong",
      status: 500,
    },
  }
}
