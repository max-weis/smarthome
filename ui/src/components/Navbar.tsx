import { Home } from 'lucide-react'
import { Button } from './ui/button'

export function Navbar() {
  return (
    <nav className="bg-secondary">
      <div className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center">
            <Button variant="ghost" className="text-primary">
              <Home className="h-5 w-5 mr-2" />
              Home
            </Button>
          </div>
        </div>
      </div>
    </nav>
  )
}