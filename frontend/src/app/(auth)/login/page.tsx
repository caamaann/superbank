import Image from "next/image"
import React from "react"
import LoginForm from "./_components/login-form"

export default function Page() {
  return (
    <section className="grid h-screen lg:grid-cols-2">
      <LoginForm />
      <div className="relative hidden lg:block">
        <Image
          src={"/images/login.jpg"}
          className="object-contain"
          alt="background"
          fill
        />
      </div>
    </section>
  )
}
