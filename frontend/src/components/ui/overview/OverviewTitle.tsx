import { ReactNode } from "react"

import DateString from "@/components/ui/DateString"
import Hr from "@/components/ui/Hr"
import FetchUserTag from "@/components/user/FetchUserTag"

export default function OverviewTitle({
  creatorID,
  title,
  titleID,
  tag,
  createdAt,

  setEditTitle,

  isEdit,
  injectHeader,
  className = "",

  breadcrumbs,
  children,
}: {
  creatorID: string
  title: string
  titleID?: any
  tag?: ReactNode
  createdAt?: Date

  setEditTitle: (title: string) => void

  isEdit: boolean
  injectHeader?: ReactNode | false
  className?: string

  breadcrumbs?: ReactNode
  children?: ReactNode
}) {
  return (
    <div className={`w-full ${className}`}>
      {breadcrumbs}

      <div className="flex flex-row items-center">
        <h1 className="mt-1 w-full space-x-2 text-2xl">
          {isEdit ? (
            <input
              className="w-full border-b border-gray-700 bg-transparent font-bold focus:outline-none"
              defaultValue={title}
              onChange={(e) => setEditTitle(e.target.value)}
              autoComplete="off"
            />
          ) : (
            <span className="font-bold">{title}</span>
          )}
          {titleID && <span className="text-neutral-400">#{titleID}</span>}
        </h1>
        {injectHeader}
      </div>

      <div className="mt-2 flex flex-row items-center space-x-2 text-neutral-500">
        {tag && <div>{tag}</div>}
        <span>created by</span>
        <FetchUserTag userID={creatorID} size="sm" />
        <span>on</span>
        <span>
          <DateString value={createdAt} date time />
        </span>
      </div>

      {children}

      <Hr className="my-4" />
    </div>
  )
}
