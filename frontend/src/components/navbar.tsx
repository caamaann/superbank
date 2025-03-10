"use client"

import Link from "next/link"
import React from "react"
import { Button } from "./ui/button"
import { useRouter } from "next/navigation"
import { deleteCookie } from "cookies-next/client"
import { ACCESS_TOKEN } from "@/lib/constants"
import Image from "next/image"

export default function Navbar() {
  const router = useRouter()

  const handleLogout = () => {
    deleteCookie(ACCESS_TOKEN)
    router.push("/login")
  }

  return (
    <header className="border-b border-dashed">
      <nav className="container flex items-center justify-between py-4">
        <Link href={"/"}>
          <Image
            src={"/next.svg"}
            alt="logo"
            width={100}
            height={36}
            className="object-contain"
          />
        </Link>
        <Button onClick={handleLogout} variant="destructive">
          Logout
        </Button>
      </nav>
    </header>
  )
}
