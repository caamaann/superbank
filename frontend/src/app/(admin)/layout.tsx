import Navbar from "@/components/navbar"
import React, { Fragment } from "react"

interface LayoutProps {
  children: React.ReactNode
}

export default function Layout({ children }: LayoutProps) {
  return (
    <Fragment>
      <Navbar />
      <main>{children}</main>
    </Fragment>
  )
}
