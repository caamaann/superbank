import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { CustomerData } from "@/types/customer"

interface CustomerDetailsProps {
  customer: CustomerData
}

export default function CustomerDetails({ customer }: CustomerDetailsProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Customer Details</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="grid gap-4 md:grid-cols-2">
          <div>
            <p className="text-muted-foreground text-sm font-medium">
              Full Name
            </p>
            <p>{customer.name}</p>
          </div>
          <div>
            <p className="text-muted-foreground text-sm font-medium">
              Customer ID
            </p>
            <p>{customer.id}</p>
          </div>
          <div>
            <p className="text-muted-foreground text-sm font-medium">Email</p>
            <p>{customer.email}</p>
          </div>
          <div>
            <p className="text-muted-foreground text-sm font-medium">Phone</p>
            <p>{customer.phone}</p>
          </div>
          <div className="md:col-span-2">
            <p className="text-muted-foreground text-sm font-medium">Address</p>
            <p>{customer.address}</p>
          </div>
          <div>
            <p className="text-muted-foreground text-sm font-medium">
              Customer Since
            </p>
            <p>{new Date(customer.createdAt).toLocaleDateString()}</p>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
