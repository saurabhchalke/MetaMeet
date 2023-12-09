import { CSSProperties, PropsWithChildren, ReactNode } from "react"

import { Priority, Tag as TagType } from "@/api/types"
import { extractErrorMessage } from "@/api/util"
import { useAuth } from "@/contexts/AuthContext"

const tagStyles = {
  plain: "border border-neutral-500 text-neutral-500",
  color: "text-black font-medium",
  "color-border": "text-white border",
} as const

export function TagContainer({
  children,
  style,
  htmlStyle,
  className = "",
  color,
}: PropsWithChildren<{
  style: keyof typeof tagStyles
  htmlStyle?: CSSProperties
  className?: string
  color?: string
}>) {
  const styleClassName = tagStyles[style]
  let styleAttr = htmlStyle

  if (color) {
    if (style === "color-border") {
      styleAttr = { ...styleAttr, borderColor: color }
    } else {
      styleAttr = { ...styleAttr, backgroundColor: color }
    }
  }

  return (
    <span
      style={styleAttr}
      className={`${className} ${styleClassName} flex w-fit flex-row items-center space-x-1 whitespace-nowrap rounded-full px-3 py-1 text-sm`}
    >
      {children}
    </span>
  )
}

export function PriorityTag({ priority }: { priority: Priority }) {
  const htmlStyle = priority.color ? { color: priority.color } : undefined
  return (
    <TagContainer key={priority.ID} style="plain" htmlStyle={htmlStyle}>
      {priority.title}
    </TagContainer>
  )
}

export function FetchPriorityTag({
  projectID,
  priorityID,
}: {
  projectID: number
  priorityID: number
  className?: string
}) {
  const { priorities } = useAuth()

  const findQuery = priorities!.useFind(projectID, priorityID)
  if (findQuery.isLoading) {
    return <>Loading...</>
  }
  if (findQuery.isError) {
    return <>Error loading priority: {extractErrorMessage(findQuery.error)}</>
  }

  return <PriorityTag priority={findQuery.data.data} />
}

export function TagFromType({
  tag,
  className = "",
}: {
  tag: TagType
  className?: string
}) {
  return (
    <TagContainer style="color-border" color={tag.color} className={className}>
      {tag.title}
    </TagContainer>
  )
}

export default function Tag({
  className = "bg-green-600 text-white",
  style,
  children,
}: {
  className?: string
  children: ReactNode
  style?: CSSProperties
}) {
  return (
    <div
      style={style}
      className={`rounded-full text-sm ${className} flex flex-row items-center space-x-1 px-3 py-1`}
    >
      {children}
    </div>
  )
}
