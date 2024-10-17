import React, { useState } from 'react'
import { Configuration } from '../../../api/device-api.schemas'
import { getSmartHomeDeviceAPI } from '../../../api/device-api'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"
import { Textarea } from "@/components/ui/textarea"
import { Save } from 'lucide-react'

interface UpdateConfigurationModalProps {
    isOpen: boolean
    onClose: () => void
    configuration: Configuration
    deviceId: string
    onUpdateSuccess: (updatedConfig: Configuration) => void
}

export function UpdateModal({
    isOpen,
    onClose,
    configuration,
    deviceId,
    onUpdateSuccess
}: UpdateConfigurationModalProps) {
    const [editedConfig, setEditedConfig] = useState<Configuration>(configuration)
    const [error, setError] = useState<string | null>(null)

    const handleSwitchChange = (checked: boolean) => {
        setEditedConfig(prev => ({ ...prev, active: checked }))
    }

    const handleDataChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        try {
            const parsedData = JSON.parse(event.target.value)
            setEditedConfig(prev => ({ ...prev, data: parsedData }))
            setError(null)
        } catch (err) {
            setError('Invalid JSON format')
        }
    }

    const handleSave = async () => {
        try {
            const api = getSmartHomeDeviceAPI()
            await api.updateConfiguration(deviceId, configuration.id, editedConfig)
            onUpdateSuccess(editedConfig)
        } catch (err) {
            console.error('Error updating configuration:', err)
            setError('Failed to update configuration. Please try again.')
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onClose}>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Edit Configuration</DialogTitle>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="name" className="text-right">
                            Name
                        </Label>
                        <input
                            id="name"
                            value={editedConfig.name}
                            onChange={(e) => setEditedConfig(prev => ({ ...prev, name: e.target.value }))}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="status" className="text-right">
                            Status
                        </Label>
                        <Switch
                            id="status"
                            checked={editedConfig.active}
                            onCheckedChange={handleSwitchChange}
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="data" className="text-right">
                            Data
                        </Label>
                        <Textarea
                            id="data"
                            value={JSON.stringify(editedConfig.data, null, 2)}
                            onChange={handleDataChange}
                            className="col-span-3"
                            rows={10}
                        />
                    </div>
                    {error && (
                        <div className="text-red-500 text-sm">{error}</div>
                    )}
                </div>
                <DialogFooter>
                    <Button onClick={handleSave} disabled={!!error}>
                        Save Changes
                        <Save className="ml-2 h-4 w-4" />
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}