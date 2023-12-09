import { Button } from "@nextui-org/react"
import clsx from "clsx"
import { ReactNode } from "react"
import { BsX } from "react-icons/bs"

import Flex from "@/components/ui/layout/Flex"

export default function ModalContainerNG({
  title,
  children,
  endContent,
  onClose,
  className,
}: {
  title: string | ReactNode
  children: ReactNode
  endContent?: ReactNode
  onClose?: () => void
  className?: string
}) {
  return (
    <Flex
      col
      className={clsx(
        className ? className : "w-[40rem]",
        "rounded-md border border-neutral-700",
      )}
    >
      {/* Header */}
      <div className={clsx("bg-neutral-900 px-4 py-3")}>
        <Flex justify="between">
          {typeof title === "string" ? (
            <h1 className="text-lg font-semibold">{title}</h1>
          ) : (
            title
          )}
          {onClose && (
            <Button
              isIconOnly
              startContent={<BsX />}
              onClick={onClose}
              size="sm"
              variant="light"
            />
          )}
        </Flex>
      </div>
      <div
        className={clsx(
          "space-y-4 border border-neutral-700 bg-neutral-950 p-4",
          { "rounded-b-md": !endContent },
        )}
      >
        {children}
      </div>
      {endContent && (
        <div
          className={clsx(
            "border-b border-l border-r border-neutral-700",
            "bg-neutral-900 px-4 py-3",
          )}
        >
          {endContent}
        </div>
      )}
    </Flex>
  )
}
