import { useNavigate } from 'react-router-dom'
import { ConfigurationListItem as ConfigItem } from '../../../api/device-api.schemas'
import { TableCell, TableRow } from "@/components/ui/table"
import { Button } from "@/components/ui/button"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal, Eye, Edit, Trash2 } from 'lucide-react'

interface ConfigurationListItemProps {
    config: ConfigItem
    deviceId: string
    isActive: boolean
    onToggle: () => void
    onUpdateConfiguration: (configId: string) => void
    onDeleteConfiguration: (configId: string) => void
}

export function ConfigurationListItem({
    config,
    deviceId,
    isActive,
    onToggle,
    onUpdateConfiguration,
    onDeleteConfiguration,
}: ConfigurationListItemProps) {
    const navigate = useNavigate()

    const handleToggle = (e: React.MouseEvent) => {
        e.stopPropagation()
        onToggle()
    }

    return (
        <TableRow
            className="bg-surface-mixed cursor-pointer"
            onClick={() => navigate(`/devices/${deviceId}/configurations/${config.id}`)}
        >
            <TableCell className="font-medium">{config.name}</TableCell>
            <TableCell>
                <Button
                    variant={isActive ? "default" : "secondary"}
                    size="sm"
                    onClick={handleToggle}
                >
                    {isActive ? "Active" : "Inactive"}
                </Button>
            </TableCell>
            <TableCell className="text-right">
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button variant="ghost" className="h-8 w-8 p-0">
                            <span className="sr-only">Open menu</span>
                            <MoreHorizontal className="h-4 w-4" />
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                        <DropdownMenuItem onClick={(e) => {
                            e.stopPropagation()
                            navigate(`/devices/${deviceId}/configurations/${config.id}`)
                        }}>
                            <Eye className="mr-2 h-4 w-4" />
                            <span>View</span>
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={(e) => {
                            e.stopPropagation()
                            onUpdateConfiguration(config.id)
                        }}>
                            <Edit className="mr-2 h-4 w-4" />
                            <span>Update</span>
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={(e) => {
                            e.stopPropagation()
                            onDeleteConfiguration(config.id)
                        }}>
                            <Trash2 className="mr-2 h-4 w-4" />
                            <span>Delete</span>
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </TableCell>
        </TableRow>
    )
}