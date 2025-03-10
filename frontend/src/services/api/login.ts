import { IApiRes } from "@/types/api"
import { ILogin } from "@/types/login"
import { errorMessage } from "@/lib/utils"
import axios from "axios"

export const apiLogin = async (data: ILogin) => {
  try {
    const response = await axios.post<IApiRes<{ access_token: string }>>(
      "/api/login",
      data
    )

    return response.data
  } catch (error) {
    return errorMessage(error)
  }
}
