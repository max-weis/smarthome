import React, { useState } from 'react'
import { Configuration, ConfigurationListItem } from '../../../api/device-api.schemas'
import { getSmartHomeDeviceAPI } from '../../../api/device-api'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Switch } from "@/components/ui/switch"
import { Textarea } from "@/components/ui/textarea"
import { Save } from 'lucide-react'

interface CreateModalProps {
    isOpen: boolean
    onClose: () => void
    deviceId: string
    onCreateSuccess: (newConfig: ConfigurationListItem) => void
}

export function CreateModal({
    isOpen,
    onClose,
    deviceId,
    onCreateSuccess
}: CreateModalProps) {
    const [newConfig, setNewConfig] = useState<Partial<Configuration>>({
        name: '',
        active: true,
        data: {}
    })
    const [dataString, setDataString] = useState('{}')
    const [error, setError] = useState<string | null>(null)

    const handleSwitchChange = (checked: boolean) => {
        setNewConfig(prev => ({ ...prev, active: checked }))
    }

    const handleDataChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setDataString(event.target.value)
    }

    const handleSave = async () => {
        if (!newConfig.name) {
            setError('Name is required')
            return
        }

        try {
            const parsedData = JSON.parse(dataString)
            const configToSave = { ...newConfig, data: parsedData }

            const api = getSmartHomeDeviceAPI()
            const createdConfig = await api.createConfiguration(deviceId, configToSave as Configuration)
            onCreateSuccess(createdConfig)
            onClose()
        } catch (err) {
            if (err instanceof SyntaxError) {
                setError('Invalid JSON format in data field')
            } else {
                console.error('Error creating configuration:', err)
                setError('Failed to create configuration. Please try again.')
            }
        }
    }

    return (
        <Dialog open={isOpen} onOpenChange={onClose}>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Create New Configuration</DialogTitle>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="name" className="text-right">
                            Name
                        </Label>
                        <Input
                            id="name"
                            value={newConfig.name || ''}
                            onChange={(e) => setNewConfig(prev => ({ ...prev, name: e.target.value }))}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="status" className="text-right">
                            Active
                        </Label>
                        <Switch
                            id="status"
                            checked={newConfig.active}
                            onCheckedChange={handleSwitchChange}
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="data" className="text-right">
                            Data
                        </Label>
                        <Textarea
                            id="data"
                            value={dataString}
                            onChange={handleDataChange}
                            className="col-span-3"
                            rows={10}
                            placeholder="Enter JSON data here"
                        />
                    </div>
                    {error && (
                        <div className="text-red-500 text-sm">{error}</div>
                    )}
                </div>
                <DialogFooter>
                    <Button onClick={handleSave} disabled={!newConfig.name}>
                        Create Configuration
                        <Save className="ml-2 h-4 w-4" />
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}