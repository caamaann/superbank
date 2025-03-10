import type { NextRequest } from "next/server"
import { NextResponse } from "next/server"
import { ACCESS_TOKEN } from "./lib/constants"

export function middleware(request: NextRequest) {
  const accessToken = request.cookies.get(ACCESS_TOKEN)?.value

  if (!accessToken) {
    return NextResponse.redirect(new URL("/login", request.url))
  }

  if (request.nextUrl.pathname.startsWith("/login")) {
    return NextResponse.redirect(new URL("/", request.url))
  }

  return NextResponse.next()
}

export const config = {
  matcher: "/((?!api|login|_next/static|_next/image|favicon|images).*)",
}
