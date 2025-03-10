import { useMutation } from "@tanstack/react-query"
import { apiLogin } from "../api/login"
import { ILogin } from "@/types/login"
import { toast } from "react-toastify"

export const useLogin = () => {
  return useMutation({
    mutationKey: ["login"],
    mutationFn: async (body: ILogin) => {
      const response = await apiLogin(body)

      if ("error" in response) {
        toast.error(response.error.message)
        throw new Error(response.error.message)
      }

      return response
    },
  })
}
