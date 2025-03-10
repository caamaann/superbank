import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Progress } from "@/components/ui/progress"
import { Pocket } from "@/types/customer"
import { calculateProgress, formatCurrency } from "@/lib/utils"

interface PocketInfoProps {
  pockets: Pocket[]
}

export default function PocketInfo({ pockets }: PocketInfoProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Pockets</CardTitle>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Balance</TableHead>
              <TableHead>Goal</TableHead>
              <TableHead>Progress</TableHead>
              <TableHead>Created</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {pockets.length === 0 ? (
              <TableRow>
                <TableCell colSpan={5} className="text-center">
                  No pockets found
                </TableCell>
              </TableRow>
            ) : (
              pockets.map((pocket) => (
                <TableRow key={pocket.id}>
                  <TableCell>{pocket.name}</TableCell>
                  <TableCell>
                    {formatCurrency(pocket.balance, pocket.currency)}
                  </TableCell>
                  <TableCell>
                    {pocket.goal
                      ? formatCurrency(pocket.goal, pocket.currency)
                      : "No goal set"}
                  </TableCell>
                  <TableCell className="w-64">
                    {pocket.goal ? (
                      <div className="space-y-2">
                        <Progress
                          value={calculateProgress(pocket.balance, pocket.goal)}
                        />
                        <p className="text-right text-xs">
                          {Math.round(
                            calculateProgress(pocket.balance, pocket.goal)
                          )}
                          %
                        </p>
                      </div>
                    ) : (
                      "N/A"
                    )}
                  </TableCell>
                  <TableCell>
                    {new Date(pocket.createdAt).toLocaleDateString()}
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
