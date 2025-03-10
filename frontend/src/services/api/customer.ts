import { IApiRes } from "@/types/api"
import { errorMessage } from "@/lib/utils"
import axios from "axios"
import { CustomerData } from "@/types/customer"

export const apiCustomerDetail = async (params: { q: string }) => {
  try {
    const response = await axios.get<IApiRes<CustomerData>>(
      "/api/customers/search",
      {
        params,
      }
    )

    return response.data
  } catch (error) {
    return errorMessage(error)
  }
}
