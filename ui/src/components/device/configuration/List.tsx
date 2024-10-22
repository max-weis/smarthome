import { useState } from 'react'
import { ConfigurationListItem as ConfigItem } from '../../../api/device-api.schemas'
import {
    Table,
    TableBody,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { ConfigurationListItem } from './ListItem'
import { getSmartHomeDeviceAPI } from '@/api/device-api'

interface ConfigurationListProps {
    deviceId: string
    configurations: ConfigItem[]
    onUpdateConfiguration: (configId: string) => void
    onDeleteConfiguration: (configId: string) => void
}

export function ConfigurationList({
    deviceId,
    configurations,
    onUpdateConfiguration,
    onDeleteConfiguration,
}: ConfigurationListProps) {
    const [activeConfigId, setActiveConfigId] = useState<string | null>(
        configurations.find(config => config.active)?.id || null
    )

    const handleToggle = async (configId: string) => {
        try {
            await getSmartHomeDeviceAPI().toggleConfigurationStatus(deviceId, configId)
            setActiveConfigId(configId)
        } catch (error) {
            console.error("Failed to toggle configuration status:", error)
            // Optionally, add error handling here (e.g., show an error message to the user)
        }
    }

    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead>Name</TableHead>
                    <TableHead>Status</TableHead>
                    <TableHead className="text-right">Actions</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {configurations.map((config) => (
                    <ConfigurationListItem
                        key={config.id}
                        config={config}
                        deviceId={deviceId}
                        isActive={config.id === activeConfigId}
                        onToggle={() => handleToggle(config.id)}
                        onUpdateConfiguration={onUpdateConfiguration}
                        onDeleteConfiguration={onDeleteConfiguration}
                    />
                ))}
            </TableBody>
        </Table>
    )
}