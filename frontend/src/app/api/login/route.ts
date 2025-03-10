import { ACCESS_TOKEN } from "@/lib/constants"
import httpRequest from "@/lib/httpRequest"
import { errorMessage } from "@/lib/utils"
import { cookies } from "next/headers"
import { NextRequest, NextResponse } from "next/server"

export async function POST(request: NextRequest) {
  const data = await request.json()

  try {
    const response = await httpRequest.post<{ access_token: string }>(
      `/api/auth/login`,
      data,
      {
        headers: {
          Authorization: `Bearer ${request.cookies.get(ACCESS_TOKEN)?.value}`,
        },
      }
    )

    const cookie = await cookies()
    cookie.set(ACCESS_TOKEN, response.data.access_token)

    return NextResponse.json(response.data)
  } catch (error: unknown) {
    const errorResponse = errorMessage(error).error
    return NextResponse.json(errorResponse, { status: errorResponse.status })
  }
}
