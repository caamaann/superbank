"use client"

import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useLogin } from "@/services/query/login"
import { zodResolver } from "@hookform/resolvers/zod"
import { useRouter } from "next/navigation"
import React from "react"
import { useForm } from "react-hook-form"
import { toast } from "react-toastify"
import { z } from "zod"

const formSchema = z.object({
  username: z.string().min(1, {
    message: "Username is required!",
  }),
  password: z.string().min(1, {
    message: "Password is required!",
  }),
})

type FormValues = z.infer<typeof formSchema>

export default function LoginForm() {
  const router = useRouter()
  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  })

  const { mutate, isPending } = useLogin()

  function onSubmit(values: FormValues) {
    mutate(values, {
      onSuccess: () => {
        toast.success("Login successful")
        router.push("/")
      },
    })
  }

  return (
    <div className="bg-accent flex items-center justify-center p-5">
      <div className="w-full max-w-md rounded-md border bg-white p-6">
        <h1 className="mb-2 text-2xl font-bold">Sign In to Your Account</h1>
        <p className="mb-6">Enter your credentials to access your account</p>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter username" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input
                      placeholder="Enter Password"
                      type="password"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit" className="w-full" loading={isPending}>
              Submit
            </Button>
          </form>
        </Form>
      </div>
    </div>
  )
}
