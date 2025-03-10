import { ACCESS_TOKEN } from "@/lib/constants"
import httpRequest from "@/lib/httpRequest"
import { errorMessage } from "@/lib/utils"
import { NextRequest, NextResponse } from "next/server"

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams
  const query = searchParams.get("q")

  if (!query) {
    return NextResponse.json(
      { error: "Search query is required" },
      { status: 400 }
    )
  }

  try {
    const response = await httpRequest.get(
      `/api/customers/search?q=${encodeURIComponent(query)}`,
      {
        params: { q: query },
        headers: {
          Authorization: `Bearer ${request.cookies.get(ACCESS_TOKEN)?.value}`,
        },
      }
    )

    return NextResponse.json(response.data)
  } catch (error) {
    const errorResponse = errorMessage(error).error
    return NextResponse.json(errorResponse, { status: errorResponse.status })
  }
}
