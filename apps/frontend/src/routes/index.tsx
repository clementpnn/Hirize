import { Button } from "@/components/ui/button"
import { H1 } from "@/components/ui/typography/h1"
import { createFileRoute, Link } from "@tanstack/react-router"
import { Lead } from "@/components/ui/typography/lead"

export const Route = createFileRoute("/")({
  component: Home
})

function Home() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen gap-16">
      <header className="flex flex-col items-center gap-14 p-8">
        <H1>
          Hirize
        </H1>
        <Lead>
          Suivez les étapes clé de vos candidatures et prévoyez facilement vos relances.
        </Lead>
        <div className="flex gap-8 justify-center">
          <Link to="/signin"><Button>Commencer</Button></Link>
          <Button variant="secondary">GitHub</Button>
        </div>
      </header>

      <img
        src="https://www.notion.so/front-static/pages/product/super-duper/carousel/calendar.png"
        alt="App mockup"
        className="max-w-screen-lg h-auto shadow-xl rounded-lg"
      />

    </main>
  )
}