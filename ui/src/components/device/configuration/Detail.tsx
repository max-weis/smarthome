import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getSmartHomeDeviceAPI } from '../../../api/device-api'
import { Configuration, Device } from '../../../api/device-api.schemas'
import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Textarea } from "@/components/ui/textarea"
import { Switch } from "@/components/ui/switch"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Edit, Save } from 'lucide-react'
import {
    Breadcrumb,
    BreadcrumbList,
    BreadcrumbItem,
    BreadcrumbLink,
    BreadcrumbPage,
    BreadcrumbSeparator
} from '../../ui/breadcrumb'
import { JsonView } from './DataView'

export function ConfigurationDetail() {
    const { deviceId, configId } = useParams<{ deviceId: string, configId: string }>()
    const [configuration, setConfiguration] = useState<Configuration | null>(null)
    const [device, setDevice] = useState<Device | null>(null)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)
    const [isEditing, setIsEditing] = useState(false)
    const [editedConfig, setEditedConfig] = useState<Configuration>(configuration as Configuration)

    useEffect(() => {
        const fetchConfigurationDetails = async () => {
            if (!deviceId || !configId) {
                setError('Invalid device or configuration ID')
                setIsLoading(false)
                return
            }
            try {
                const api = getSmartHomeDeviceAPI()
                const config = await api.getConfiguration(deviceId, configId)
                if (!config) {
                    throw new Error('Configuration not found')
                }
                setConfiguration(config)
                setEditedConfig(config)
                setIsLoading(false)

                const deviceResponse = await api.getDevice(deviceId)
                setDevice(deviceResponse)
            } catch (err) {
                console.error('Error fetching configuration details:', err)
                setError('Failed to fetch configuration details. Please try again later.')
                setIsLoading(false)
            }
        }

        fetchConfigurationDetails()
    }, [deviceId, configId])

    const handleSwitchChange = (checked: boolean) => {
        setEditedConfig(prev => ({ ...prev, active: checked }))
    }

    const handleSave = async () => {
        try {
            const api = getSmartHomeDeviceAPI()
            await api.updateConfiguration(deviceId!, configId!, editedConfig)
            setConfiguration(editedConfig as Configuration)
            setIsEditing(false)
        } catch (err) {
            console.error('Error updating configuration:', err)
            setError('Failed to update configuration. Please try again.')
        }
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
            <Breadcrumb>
                <BreadcrumbList>
                    <BreadcrumbItem>
                        <BreadcrumbLink href="/">Home</BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbLink href="/devices">Devices</BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbLink href={`/devices/${deviceId}`}>{device?.name}</BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbPage>{configuration.name}</BreadcrumbPage>
                    </BreadcrumbItem>
                </BreadcrumbList>
            </Breadcrumb>

            <Card className="bg-surface-mixed mt-4">
                <CardHeader className="flex flex-row items-center justify-between">
                    <CardTitle className="text-2xl font-bold">{configuration.name}</CardTitle>
                    <Button onClick={() => setIsEditing(!isEditing)} variant="outline">
                        {isEditing ? (
                            <>
                                Cancel
                                <Edit className="ml-2 h-4 w-4" />
                            </>
                        ) : (
                            <>
                                Edit
                                <Edit className="ml-2 h-4 w-4" />
                            </>
                        )}
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
                                    {isEditing ? (
                                        <Switch
                                            id="status"
                                            checked={editedConfig.active}
                                            onCheckedChange={handleSwitchChange}
                                        />
                                    ) : null}
                                    <Badge
                                        variant={(isEditing ? editedConfig.active : configuration.active) ? "default" : "secondary"}
                                    >
                                        {(isEditing ? editedConfig.active : configuration.active) ? "Active" : "Inactive"}
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
                {isEditing && (
                    <CardFooter>
                        <Button onClick={handleSave}>
                            Save Changes
                            <Save className="ml-2 h-4 w-4" />
                        </Button>
                    </CardFooter>
                )}
            </Card>
        </div>
    )
}