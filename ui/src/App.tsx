import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import { Navbar } from './components/Navbar'
import { DevicePage } from './components/device/Page'
import { DeviceDetails } from './components/device/Details'
import { ConfigurationDetail } from './components/device/configuration/Detail'

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
          <Route path="/devices" element={<DevicePage />} />
          <Route path="/devices/:id" element={<DeviceDetails />} />
          <Route path="/devices/:deviceId/configurations/:configId" element={<ConfigurationDetail />} />
        </Routes>
      </div>
    </Router>
  )
}

export default App