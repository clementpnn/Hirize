import { UserService } from "@/services/user"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute("/(user)/dashboard")({
  component: Dashboard,
  loader: async () => {
    const response = await UserService.checkSession()

    try {
      if (!response.data) {
        return {
          redirect: "/login"
        };
      }
    } catch {
      return {
        redirect: "/login"
      }
    }
    return null
  },
})

export default function Dashboard() {
  return <div>Dashboard</div>
}