import { ReactNode } from "react"

// TODO: Add more admonition types if needed
export type AdmonitionStyle = "danger"

const admonitionColors: Record<AdmonitionStyle, string> = {
  danger: "bg-red-500 bg-opacity-25 text-red-500",
} as const

export default function Admonition({
  children,
  style,
  fullWidth,
  className,
}: {
  children: ReactNode
  style: AdmonitionStyle
  fullWidth?: boolean
  className?: string
}) {
  const classNames = [
    "flex items-center space-x-2 rounded-md p-4", // base
    admonitionColors[style], // color
  ]
  // other styles
  fullWidth && classNames.push("w-full")
  // custom class
  className && classNames.push(className)

  return <div className={classNames.join(" ")}>{children}</div>
}
