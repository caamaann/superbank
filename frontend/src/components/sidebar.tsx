import { Fragment } from "react"
import { Dialog, Transition } from "@headlessui/react"
import { XMarkIcon } from "@heroicons/react/24/outline"
import {
  HomeIcon,
  UserGroupIcon,
  BanknotesIcon,
  CreditCardIcon,
} from "@heroicons/react/24/solid"
import Link from "next/link"
import { usePathname } from "next/navigation"

function NavItem({ href, icon: Icon, children }) {
  const pathname = usePathname()
  const isActive = pathname === href || pathname.startsWith(`${href}/`)

  return (
    <Link
      href={href}
      className={`group flex items-center rounded-md px-2 py-2 text-sm font-medium ${
        isActive
          ? "bg-gray-900 text-white"
          : "text-gray-300 hover:bg-gray-700 hover:text-white"
      }`}
    >
      <Icon className="mr-3 h-6 w-6 flex-shrink-0" aria-hidden="true" />
      {children}
    </Link>
  )
}

export default function Sidebar({ sidebarOpen, setSidebarOpen }) {
  return (
    <>
      {/* Mobile sidebar */}
      <Transition.Root show={sidebarOpen} as={Fragment}>
        <Dialog
          as="div"
          className="relative z-40 md:hidden"
          onClose={setSidebarOpen}
        >
          <Transition.Child
            as={Fragment}
            enter="transition-opacity ease-linear duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="transition-opacity ease-linear duration-300"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="bg-opacity-75 fixed inset-0 bg-gray-600" />
          </Transition.Child>

          <div className="fixed inset-0 z-40 flex">
            <Transition.Child
              as={Fragment}
              enter="transition ease-in-out duration-300 transform"
              enterFrom="-translate-x-full"
              enterTo="translate-x-0"
              leave="transition ease-in-out duration-300 transform"
              leaveFrom="translate-x-0"
              leaveTo="-translate-x-full"
            >
              <Dialog.Panel className="relative flex w-full max-w-xs flex-1 flex-col bg-gray-800">
                <Transition.Child
                  as={Fragment}
                  enter="ease-in-out duration-300"
                  enterFrom="opacity-0"
                  enterTo="opacity-100"
                  leave="ease-in-out duration-300"
                  leaveFrom="opacity-100"
                  leaveTo="opacity-0"
                >
                  <div className="absolute top-0 right-0 -mr-12 pt-2">
                    <button
                      type="button"
                      className="ml-1 flex h-10 w-10 items-center justify-center rounded-full focus:ring-2 focus:ring-white focus:outline-none focus:ring-inset"
                      onClick={() => setSidebarOpen(false)}
                    >
                      <span className="sr-only">Close sidebar</span>
                      <XMarkIcon
                        className="h-6 w-6 text-white"
                        aria-hidden="true"
                      />
                    </button>
                  </div>
                </Transition.Child>
                <div className="h-0 flex-1 overflow-y-auto pt-5 pb-4">
                  <div className="flex flex-shrink-0 items-center px-4">
                    <h1 className="text-lg font-bold text-white">
                      Customer Dashboard
                    </h1>
                  </div>
                  <nav className="mt-5 space-y-1 px-2">
                    <NavItem href="/dashboard" icon={HomeIcon}>
                      Dashboard
                    </NavItem>
                    <NavItem href="/dashboard/customers" icon={UserGroupIcon}>
                      Customers
                    </NavItem>
                    <NavItem href="/dashboard/accounts" icon={BanknotesIcon}>
                      Accounts
                    </NavItem>
                    <NavItem
                      href="/dashboard/term-deposits"
                      icon={CreditCardIcon}
                    >
                      Term Deposits
                    </NavItem>
                  </nav>
                </div>
              </Dialog.Panel>
            </Transition.Child>
            <div className="w-14 flex-shrink-0">
              {/* Force sidebar to shrink to fit close icon */}
            </div>
          </div>
        </Dialog>
      </Transition.Root>

      {/* Static sidebar for desktop */}
      <div className="hidden md:fixed md:inset-y-0 md:flex md:w-64 md:flex-col">
        <div className="flex min-h-0 flex-1 flex-col bg-gray-800">
          <div className="flex flex-1 flex-col overflow-y-auto pt-5 pb-4">
            <div className="flex flex-shrink-0 items-center px-4">
              <h1 className="text-lg font-bold text-white">
                Customer Dashboard
              </h1>
            </div>
            <nav className="mt-5 flex-1 space-y-1 px-2">
              <NavItem href="/dashboard" icon={HomeIcon}>
                Dashboard
              </NavItem>
              <NavItem href="/dashboard/customers" icon={UserGroupIcon}>
                Customers
              </NavItem>
              <NavItem href="/dashboard/accounts" icon={BanknotesIcon}>
                Accounts
              </NavItem>
              <NavItem href="/dashboard/term-deposits" icon={CreditCardIcon}>
                Term Deposits
              </NavItem>
            </nav>
          </div>
        </div>
      </div>
    </>
  )
}
