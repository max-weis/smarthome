import { useEffect, useState } from 'react'
import { getSmartHomeDeviceAPI } from '../../api/device-api'
import { Device } from '../../api/device-api.schemas'
import { TileView } from './TileView'
import { TableView } from './TableView'
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import { LayoutDashboard, Table } from 'lucide-react'

export function DeviceList() {
  const [devices, setDevices] = useState<Device[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [isTileView, setIsTileView] = useState(true)

  useEffect(() => {
    const fetchDevices = async () => {
      try {
        const api = getSmartHomeDeviceAPI()
        const response = await api.getDevices()
        setDevices(response)
        setIsLoading(false)
      } catch (err) {
        setError('Failed to fetch devices. Please try again later.')
        setIsLoading(false)
      }
    }

    fetchDevices()
  }, [])

  const handleView = (device: Device) => {
    console.log('View device:', device)
    // Implement view functionality
  }

  const handleEdit = (device: Device) => {
    console.log('Edit device:', device)
    // Implement edit functionality
  }

  if (isLoading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="flex justify-center items-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary"></div>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center text-red-500 p-4">
          <p>{error}</p>
        </div>
      </div>
    )
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">Devices</h1>
        <div className="flex items-center space-x-2">
          <Label htmlFor="view-toggle">
            <Table className="h-6 w-6" />
          </Label>
          <Switch
            id="view-toggle"
            checked={isTileView}
            onCheckedChange={setIsTileView}
          />
          <Label htmlFor="view-toggle">
            <LayoutDashboard className="h-6 w-6" />
          </Label>
        </div>
      </div>
      {isTileView ? (
        <TileView devices={devices} onView={handleView} onEdit={handleEdit} />
      ) : (
        <TableView devices={devices} onView={handleView} onEdit={handleEdit} />
      )}
    </div>
  )
}