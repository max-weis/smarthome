import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import { Navbar } from './components/Navbar'
import { DeviceList } from './components/device/DeviceList'

function HomePage() {
  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Hello World</h1>
    </div>
  )
}

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-surface text-foreground">
        <Navbar />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/devices" element={<DeviceList />} />
        </Routes>
      </div>
    </Router>
  )
}

export default App