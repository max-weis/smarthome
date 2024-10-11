import { Navbar } from './components/Navbar'
import { Button } from './components/ui/button'

function App() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <Navbar />

      <div className="container mx-auto p-4">
        <h1 className="text-2xl font-bold mb-4">My Smarthome UI</h1>
        <Button>Click me</Button>
      </div>
    </div>
  )
}

export default App