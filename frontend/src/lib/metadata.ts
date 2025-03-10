import { Metadata } from "next"
import { SITE_URL } from "./constants"

const title = "Superbank Test"
const description = "Website Test"

export const baseMetadata: Metadata = {
  metadataBase: new URL(`${SITE_URL}`),
  title: {
    default: `${title}`,
    template: `${title} | %s`,
  },
  openGraph: {
    type: "website",
    title: `${title}`,
    description: `${description}`,
    siteName: `${title}`,
    images: [
      {
        url: `/icons/richlink.jpg`,
      },
    ],
  },
  icons: [
    { rel: "icon", url: "/icons/favicon-32x32.png" },
    { rel: "apple-touch-icon", url: "/icons/apple-touch-icon.png" },
  ],
  twitter: {
    card: "summary_large_image",
    site: `${title}`,
    creator: "Salman",
    images: `/icons/richlink.jpg`,
  },
  description: `${description}`,
  authors: [{ name: "Salman", url: "https://github.com/caamaann" }],
}
