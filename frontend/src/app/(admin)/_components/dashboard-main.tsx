"use client"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { useCustomerDetail } from "@/services/query/customer"
import { zodResolver } from "@hookform/resolvers/zod"
import React, { Fragment, useState } from "react"
import { useForm } from "react-hook-form"
import { toast } from "react-toastify"
import { z } from "zod"
import CustomerDetails from "./customer-details"
import BankAccountInfo from "./bank-account-info"
import PocketInfo from "./pocket-info"
import TermDeposits from "./term-deposits"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"

const formSchema = z.object({
  q: z.string().min(1, {
    message: "Field is required!",
  }),
})

type FormValues = z.infer<typeof formSchema>

export default function DashboardMain() {
  const [query, setQuery] = useState("")

  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      q: "",
    },
  })

  const { data: customer, isFetching } = useCustomerDetail({ q: query })

  function onSubmit(values: FormValues) {
    if (!values.q.trim()) {
      toast.error("Please enter a customer ID or name")
      return
    }
    setQuery(values.q)
  }

  return (
    <Fragment>
      <Card className="mb-8">
        <CardHeader>
          <CardTitle>Customer Search</CardTitle>
          <CardDescription>
            Search by customer ID, name, or email
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex gap-4">
              <FormField
                control={form.control}
                name="q"
                render={({ field }) => (
                  <FormItem className="flex-1">
                    <FormControl>
                      <Input
                        placeholder="Enter customer details..."
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <Button type="submit" className="min-w-24" loading={isFetching}>
                Submit
              </Button>
            </form>
          </Form>
        </CardContent>
      </Card>

      {customer && (
        <>
          <div className="mb-6 flex items-center gap-4">
            <Avatar className="h-16 w-16">
              <AvatarImage src={customer.profileImage} alt={customer.name} />
              <AvatarFallback>
                {customer.name
                  .split(" ")
                  .map((n) => n[0])
                  .join("")}
              </AvatarFallback>
            </Avatar>
            <div>
              <h2 className="text-2xl font-bold">{customer.name}</h2>
              <p className="text-muted-foreground">ID: {customer.id}</p>
            </div>
          </div>

          <Tabs defaultValue="details" className="mb-10">
            <TabsList className="grid h-auto w-full max-w-xl grid-cols-2 md:grid-cols-4">
              <TabsTrigger value="details">Details</TabsTrigger>
              <TabsTrigger value="accounts">Bank Accounts</TabsTrigger>
              <TabsTrigger value="pockets">Pockets</TabsTrigger>
              <TabsTrigger value="deposits">Term Deposits</TabsTrigger>
            </TabsList>

            <TabsContent value="details">
              <CustomerDetails customer={customer} />
            </TabsContent>

            <TabsContent value="accounts">
              <BankAccountInfo accounts={customer.bankAccounts} />
            </TabsContent>

            <TabsContent value="pockets">
              <PocketInfo pockets={customer.pockets} />
            </TabsContent>

            <TabsContent value="deposits">
              <TermDeposits deposits={customer.termDeposits} />
            </TabsContent>
          </Tabs>
        </>
      )}
    </Fragment>
  )
}
