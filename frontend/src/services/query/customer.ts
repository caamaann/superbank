import { useQuery } from "@tanstack/react-query"
import { toast } from "react-toastify"
import { apiCustomerDetail } from "../api/customer"

export const useCustomerDetail = (params: { q: string }) => {
  return useQuery({
    queryKey: ["customer", params.q],
    queryFn: async () => {
      const response = await apiCustomerDetail(params)

      if ("error" in response) {
        toast.error(response.error.message)
        throw new Error(response.error.message)
      }

      return response.data
    },
    enabled: !!params.q,
  })
}
