import axios from "axios"
import { API_URL } from "./constants"

const httpRequest = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
})

httpRequest.interceptors.request.use(
  function (config) {
    return config
  },
  function (error) {
    return Promise.reject(error)
  }
)

httpRequest.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    return Promise.reject(error)
  }
)

export default httpRequest
