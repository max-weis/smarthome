import { useNavigate } from 'react-router-dom'
import { ConfigurationListItem } from '../../../api/device-api.schemas'
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { Button } from "@/components/ui/button"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal, Eye, Edit, Trash2 } from 'lucide-react'

interface ConfigurationListProps {
    deviceId: string
    configurations: ConfigurationListItem[]
    toggleConfigurationActive: (configId: string) => void
    onUpdateConfiguration: (configId: string) => void
    onDeleteConfiguration: (configId: string) => void
}

export function ConfigurationList({
    deviceId,
    configurations,
    toggleConfigurationActive,
    onUpdateConfiguration,
    onDeleteConfiguration
}: ConfigurationListProps) {
    const navigate = useNavigate()

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
                    <TableRow key={config.id}>
                        <TableCell className="font-medium">{config.name}</TableCell>
                        <TableCell>
                            <Button
                                variant={config.active ? "default" : "secondary"}
                                size="sm"
                                onClick={() => toggleConfigurationActive(config.id)}
                            >
                                {config.active ? "Active" : "Inactive"}
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
                                    <DropdownMenuItem onClick={() => navigate(`/devices/${deviceId}/configurations/${config.id}`)}>
                                        <Eye className="mr-2 h-4 w-4" />
                                        <span>View</span>
                                    </DropdownMenuItem>
                                    <DropdownMenuItem onClick={() => onUpdateConfiguration(config.id)}>
                                        <Edit className="mr-2 h-4 w-4" />
                                        <span>Update</span>
                                    </DropdownMenuItem>
                                    <DropdownMenuItem onClick={() => onDeleteConfiguration(config.id)}>
                                        <Trash2 className="mr-2 h-4 w-4" />
                                        <span>Delete</span>
                                    </DropdownMenuItem>
                                </DropdownMenuContent>
                            </DropdownMenu>
                        </TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    )
}