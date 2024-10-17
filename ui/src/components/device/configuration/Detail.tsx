import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getSmartHomeDeviceAPI } from '../../../api/device-api'
import { Configuration, Device } from '../../../api/device-api.schemas'
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Edit } from 'lucide-react'
import { JsonView } from './DataView'
import { Breadcrumbs } from './Breadcrumb'
import { UpdateModal } from './UpdateModal'

export function ConfigurationDetail() {
    const { deviceId, configId } = useParams<{ deviceId: string, configId: string }>()
    const [configuration, setConfiguration] = useState<Configuration>({} as Configuration)
    const [device, setDevice] = useState<Device>({} as Device)
    const [error, setError] = useState<string>('')
    const [isModalOpen, setIsModalOpen] = useState(false)

    useEffect(() => {
        const fetchConfigurationDetails = async () => {
            if (!deviceId || !configId) {
                setError('Invalid device or configuration ID')
                return
            }
            try {
                const api = getSmartHomeDeviceAPI()
                const config = await api.getConfiguration(deviceId, configId)
                if (!config) {
                    throw new Error('Configuration not found')
                }
                setConfiguration(config)

                const deviceResponse = await api.getDevice(deviceId)
                setDevice(deviceResponse)
            } catch (err) {
                console.error('Error fetching configuration details:', err)
                setError('Failed to fetch configuration details. Please try again later.')
            }
        }

        fetchConfigurationDetails()
    }, [deviceId, configId])

    const handleUpdateSuccess = (updatedConfig: Configuration) => {
        setConfiguration(updatedConfig)
        setIsModalOpen(false)
    }

    if (error || !configuration) {
        return (
            <div className="container mx-auto px-4 py-8">
                <div className="text-center text-red-500 p-4">
                    <p>{error || 'Configuration not found'}</p>
                </div>
            </div>
        )
    }

    if (error || !configuration) {
        return (
            <div className="container mx-auto px-4 py-8">
                <div className="text-center text-red-500 p-4">
                    <p>{error || 'Configuration not found'}</p>
                </div>
            </div>
        )
    }

    return (
        <div className="container mx-auto px-4 py-8">
            <Breadcrumbs
                deviceId={device.id}
                deviceName={device.name}
                configurationName={configuration.name}
            />

            <Card className="bg-surface-mixed mt-4">
                <CardHeader className="flex flex-row items-center justify-between">
                    <CardTitle className="text-2xl font-bold">{configuration.name}</CardTitle>
                    <Button onClick={() => setIsModalOpen(true)} variant="outline">
                        Edit
                        <Edit className="ml-2 h-4 w-4" />
                    </Button>
                </CardHeader>
                <CardContent>
                    <div className="flex flex-col lg:flex-row gap-8">
                        <div className="flex-1 space-y-4">
                            <div>
                                <Label htmlFor="id">ID</Label>
                                <p className="mt-1">{configuration.id}</p>
                            </div>
                            <div>
                                <Label htmlFor="status">Status</Label>
                                <div className="flex items-center space-x-2 mt-1">
                                    <Badge variant={configuration.active ? "default" : "secondary"}>
                                        {configuration.active ? "Active" : "Inactive"}
                                    </Badge>
                                </div>
                            </div>
                        </div>
                        <div className="flex-1">
                            <Label>Configuration Data</Label>
                            <div className="mt-1">
                                <JsonView data={configuration.data} />
                            </div>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <UpdateModal
                isOpen={isModalOpen}
                onClose={() => setIsModalOpen(false)}
                configuration={configuration}
                deviceId={deviceId!}
                onUpdateSuccess={handleUpdateSuccess}
            />
        </div>
    )
}