import type { NextConfig } from "next"

const nextConfig: NextConfig = {
  /* config options here */
  env: {
    SITE_URL: process.env.SITE_URL,
    API_URL: process.env.API_URL,
  },
}

export default nextConfig
