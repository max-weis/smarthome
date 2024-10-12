import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getSmartHomeDeviceAPI } from '../../../api/device-api'
import { ConfigurationListItem } from '../../../api/device-api.schemas'
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import {
    Breadcrumb,
    BreadcrumbList,
    BreadcrumbItem,
    BreadcrumbLink,
    BreadcrumbPage,
    BreadcrumbSeparator
} from '../../ui/breadcrumb'

export function ConfigurationDetail() {
    const { deviceId, configId } = useParams<{ deviceId: string, configId: string }>()
    const [configuration, setConfiguration] = useState<ConfigurationListItem | null>(null)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)

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
                setIsLoading(false)
            } catch (err) {
                console.error('Error fetching configuration details:', err)
                setError('Failed to fetch configuration details. Please try again later.')
                setIsLoading(false)
            }
        }

        fetchConfigurationDetails()
    }, [deviceId, configId])

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
                        <BreadcrumbLink href={`/devices/${deviceId}`}>Device Details</BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbPage>{configuration.name}</BreadcrumbPage>
                    </BreadcrumbItem>
                </BreadcrumbList>
            </Breadcrumb>

            <Card className="bg-surface-mixed mt-4">
                <CardHeader>
                    <CardTitle className="text-2xl font-bold">{configuration.name}</CardTitle>
                </CardHeader>
                <CardContent>
                    <div className="grid gap-4">
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">ID</p>
                            <p className="mt-1">{configuration.id}</p>
                        </div>
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">Status</p>
                            <Badge
                                variant={configuration.active ? "default" : "secondary"}
                                className="mt-1"
                            >
                                {configuration.active ? "Active" : "Inactive"}
                            </Badge>
                        </div>
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">Value</p>
                            <p className="mt-1">{configuration.value}</p>
                        </div>
                        <div>
                            <p className="text-sm font-medium text-muted-foreground">Description</p>
                            <p className="mt-1">{configuration.description || 'No description available'}</p>
                        </div>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}