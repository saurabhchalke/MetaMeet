import { Link, ScrollShadow } from "@nextui-org/react"
import NextLink from "next/link"

import { Action } from "@/api/types"
import ActionTag from "@/components/action/ActionTag"
import CommentSuite from "@/components/comment/CommentSuite"
import ResolveUserName from "@/components/resolve/ResolveUserName"
import { RelativeDate } from "@/components/ui/DateString"
import Flex from "@/components/ui/layout/Flex"
import ModalContainerNG from "@/components/ui/modal/ModalContainerNG"
import { PriorityTag, TagContainer } from "@/components/ui/tag/Tag"
import TagList from "@/components/ui/tag/TagList"
import RenderMarkdown from "@/components/ui/text/RenderMarkdown"

export default function ActionPeekModal({
  action,
  onClose,
}: {
  action: Action
  onClose: () => void
}) {
  return (
    <ModalContainerNG
      title={
        <Flex col gap={2}>
          <Flex gap={2}>
            <ActionTag action={action} />
            <h2 className="text-xl font-semibold">{action.title}</h2>
          </Flex>
          <span>
            Created by{" "}
            <span className="text-neutral-400">
              <ResolveUserName userID={action.creator_id} />
            </span>{" "}
            - <RelativeDate date={new Date(action.CreatedAt)} />
          </span>
          <TagList>
            {!!action.priority_id && (
              <span className="w-fit">
                <PriorityTag priority={action.priority!} />
              </span>
            )}
            {action.tags.map((tag) => (
              <TagContainer
                key={tag.ID}
                style="color"
                color={tag.color}
                className="text-xs"
              >
                {tag.title}
              </TagContainer>
            ))}
          </TagList>
        </Flex>
      }
      endContent={
        <Flex justify="end">
          <Link
            href={`/project/${action.project_id}/action/${action.ID}`}
            as={NextLink}
            showAnchorIcon
            color="secondary"
          >
            Open Action
          </Link>
        </Flex>
      }
      onClose={onClose}
    >
      <div className="w-full bg-neutral-950 p-3">
        <RenderMarkdown markdown={action.description || "*no description*"} />
      </div>

      <ScrollShadow className="max-h-64">
        <CommentSuite
          projectID={action.project_id}
          commentType="action"
          commentEntityID={action.ID}
        />
      </ScrollShadow>
    </ModalContainerNG>
  )
}
