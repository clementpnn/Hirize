import { cn } from "@/libs/utils"

export function Lead({ children, className }: { children: React.ReactNode, className?: string }) {
  return (
    <p className={cn("text-xl text-muted-foreground", className)}>
      {children}
    </p>
  )
}
