import { cn } from "@/libs/utils"

export function P({ children, className }: { children: React.ReactNode, className?: string }) {
  return (
    <p className={cn("leading-7 [&:not(:first-child)]:mt-6", className)} >
      {children}
    </p>
  )
}
