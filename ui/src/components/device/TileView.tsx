import { Device } from '../../api/device-api.schemas'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Lightbulb, Thermometer, HelpCircle, MoreVertical, Eye, Edit, Trash2 } from 'lucide-react'
import { cn } from '@/lib/utils'
import { useNavigate } from 'react-router-dom'

interface TileViewProps {
    devices: Device[]
    onEdit: (device: Device) => void
}

export function TileView({ devices, onEdit }: TileViewProps) {
    const navigate = useNavigate()

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

    const getStatusColor = (status: string) => {
        switch (status.toLowerCase()) {
            case 'on':
                return 'bg-green-500'
            case 'off':
                return 'bg-red-500'
            case 'idle':
                return 'bg-yellow-500'
            default:
                return 'bg-gray-500'
        }
    }

    return (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            {devices.map((device) => (
                <Card key={device.id} className="bg-surface-mixed">
                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                            {device.name}
                        </CardTitle>
                        <div className="flex items-center space-x-2">
                            {getDeviceIcon(device.type)}
                            <DropdownMenu>
                                <DropdownMenuTrigger asChild>
                                    <Button variant="ghost" className="h-8 w-8 p-0">
                                        <span className="sr-only">Open menu</span>
                                        <MoreVertical className="h-4 w-4" />
                                    </Button>
                                </DropdownMenuTrigger>
                                <DropdownMenuContent align="end">
                                    <DropdownMenuItem onClick={() => navigate(`/devices/${device.id}`)}>
                                        <Eye className="mr-2 h-4 w-4" />
                                        <span>View</span>
                                    </DropdownMenuItem>
                                    <DropdownMenuItem onClick={() => onEdit(device)}>
                                        <Edit className="mr-2 h-4 w-4" />
                                        <span>Edit</span>
                                    </DropdownMenuItem>
                                    <DropdownMenuItem disabled>
                                        <Trash2 className="mr-2 h-4 w-4" />
                                        <span>Delete</span>
                                    </DropdownMenuItem>
                                </DropdownMenuContent>
                            </DropdownMenu>
                        </div>
                    </CardHeader>
                    <CardContent>
                        <div className="flex justify-between items-center">
                            <Badge variant="outline" className="text-xs">
                                {device.type}
                            </Badge>
                            <div className={cn(
                                "h-3 w-3 rounded-full",
                                getStatusColor(device.status)
                            )} />
                        </div>
                        <p className="text-xs text-muted-foreground mt-2">
                            Status: {device.status}
                        </p>
                    </CardContent>
                </Card>
            ))}
        </div>
    )
}