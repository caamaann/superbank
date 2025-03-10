export interface CustomerData {
  id: string
  name: string
  email: string
  phone: string
  address: string
  profileImage?: string
  createdAt: string
  bankAccounts: BankAccount[]
  pockets: Pocket[]
  termDeposits: TermDeposit[]
}

export interface BankAccount {
  id: string
  accountNumber: string
  accountType: string
  balance: number
  currency: string
  isActive: boolean
  createdAt: string
}

export interface Pocket {
  id: string
  name: string
  balance: number
  currency: string
  goal?: number
  createdAt: string
}

export interface TermDeposit {
  id: string
  amount: number
  currency: string
  interestRate: number
  startDate: string
  maturityDate: string
  isActive: boolean
}
