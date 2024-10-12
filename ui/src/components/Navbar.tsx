import { Home, Server } from 'lucide-react'
import { Button } from './ui/button'
import { Link } from 'react-router-dom'

export function Navbar() {
  return (
    <nav className="bg-surface-mixed">
      <div className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center space-x-4">
            <Link to="/">
              <Button variant="ghost" className="text-primary hover:text-primary-foreground hover:bg-primary/90">
                <Home className="h-5 w-5 mr-2" />
                Home
              </Button>
            </Link>
            <Link to="/devices">
              <Button variant="ghost" className="text-primary hover:text-primary-foreground hover:bg-primary/90">
                <Server className="h-5 w-5 mr-2" />
                Devices
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </nav>
  )
}