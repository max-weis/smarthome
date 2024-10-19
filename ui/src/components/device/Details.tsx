import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getSmartHomeDeviceAPI } from '../../api/device-api'
import { Configuration, ConfigurationListItem, Device } from '../../api/device-api.schemas'
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Lightbulb, Thermometer, HelpCircle, Plus } from 'lucide-react'
import { cn } from '@/lib/utils'
import { ConfigurationList } from './configuration/List'
import { Breadcrumbs } from './Breadcrumbs'
import { CreateModal } from './configuration/CreateModal'

export function DeviceDetails() {
    const { id } = useParams<{ id: string }>()
    const [device, setDevice] = useState<Device | null>(null)
    const [configurations, setConfigurations] = useState<ConfigurationListItem[]>([])
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false)

    useEffect(() => {
        const fetchDeviceAndConfigurations = async () => {
            if (!id) return
            try {
                const api = getSmartHomeDeviceAPI()
                const deviceResponse = await api.getDevice(id)
                setDevice(deviceResponse)

                const configurations = await api.getDeviceConfigurations(id)
                setConfigurations(configurations)

                setIsLoading(false)
            } catch (err) {
                setError('Failed to fetch device details. Please try again later.')
                setIsLoading(false)
            }
        }

        fetchDeviceAndConfigurations()
    }, [id])

    const getDeviceIcon = (type: string) => {
        switch (type.toLowerCase()) {
            case 'light':
                return <Lightbulb className="h-6 w-6" />
            case 'thermostat':
                return <Thermometer className="h-6 w-6" />
            default:
                return <HelpCircle className="h-6 w-6" />
        }
    }

    const handleUpdateConfiguration = (configId: string) => {
        console.log(`Update configuration ${configId}`)
    }

    const handleDeleteConfiguration = (configId: string) => {
        console.log(`Delete configuration ${configId}`)
    }

    const handleCreateConfiguration = (newConfig: ConfigurationListItem) => {
        setConfigurations([...configurations, newConfig])
        setIsCreateModalOpen(false)
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

    if (error || !device) {
        return (
            <div className="container mx-auto px-4 py-8">
                <div className="text-center text-red-500 p-4">
                    <p>{error || 'Device not found'}</p>
                </div>
            </div>
        )
    }

    return (
        <div className="container mx-auto px-4 py-8">
            <Breadcrumbs deviceName={device.name} />

            <Card className="bg-surface-mixed mb-8 mt-4">
                <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle className="text-2xl font-bold">{device.name}</CardTitle>
                    {getDeviceIcon(device.type)}
                </CardHeader>
                <CardContent>
                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">Type</p>
                            <Badge variant="outline" className="mt-1">
                                {device.type}
                            </Badge>
                        </div>
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">Status</p>
                            <div className="flex items-center mt-1">
                                <div className={cn(
                                    "h-3 w-3 rounded-full mr-2",
                                )} />
                                {device.status}
                            </div>
                        </div>
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">ID</p>
                            <p className="mt-1">{device.id}</p>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <Card className="bg-surface-mixed">
                <CardHeader className="flex flex-row items-center justify-between space-y-0">
                    <CardTitle>Device Configurations</CardTitle>
                    <Button onClick={() => setIsCreateModalOpen(true)} variant="outline" size="sm">
                        <Plus className="mr-2 h-4 w-4" />
                        Create Configuration
                    </Button>
                </CardHeader>
                <CardContent>
                    <ConfigurationList
                        deviceId={device.id}
                        configurations={configurations}
                        onUpdateConfiguration={handleUpdateConfiguration}
                        onDeleteConfiguration={handleDeleteConfiguration}
                    />
                </CardContent>
            </Card>

            <CreateModal
                isOpen={isCreateModalOpen}
                onClose={() => setIsCreateModalOpen(false)}
                deviceId={device.id}
                onCreateSuccess={handleCreateConfiguration}
            />
        </div>
    )
}