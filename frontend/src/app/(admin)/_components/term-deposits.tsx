import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { TermDeposit } from "@/types/customer"
import { formatCurrency } from "@/lib/utils"

interface TermDepositsProps {
  deposits: TermDeposit[]
}

export default function TermDeposits({ deposits }: TermDepositsProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Term Deposits</CardTitle>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Amount</TableHead>
              <TableHead>Interest Rate</TableHead>
              <TableHead>Start Date</TableHead>
              <TableHead>Maturity Date</TableHead>
              <TableHead>Status</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {deposits.length === 0 ? (
              <TableRow>
                <TableCell colSpan={5} className="text-center">
                  No term deposits found
                </TableCell>
              </TableRow>
            ) : (
              deposits.map((deposit) => (
                <TableRow key={deposit.id}>
                  <TableCell>
                    {formatCurrency(deposit.amount, deposit.currency)}
                  </TableCell>
                  <TableCell>{deposit.interestRate.toFixed(2)}%</TableCell>
                  <TableCell>
                    {new Date(deposit.startDate).toLocaleDateString()}
                  </TableCell>
                  <TableCell>
                    {new Date(deposit.maturityDate).toLocaleDateString()}
                  </TableCell>
                  <TableCell>
                    <Badge variant={deposit.isActive ? "default" : "secondary"}>
                      {deposit.isActive ? "Active" : "Matured"}
                    </Badge>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}
