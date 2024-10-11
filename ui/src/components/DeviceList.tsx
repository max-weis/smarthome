import React, { useEffect, useState } from 'react';
import { getDevices } from '../api/device-api';

export function DeviceList() {
    const [devices, setDevices] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchDevices = async () => {
            try {
                const response = await getDevices();
                setDevices(response.data);
                setIsLoading(false);
            } catch (err) {
                setError(err.message);
                setIsLoading(false);
            }
        };

        fetchDevices();
    }, []);

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div>
            <h2 className="text-xl font-bold mb-4">Devices</h2>
            <ul>
                {devices.map((device) => (
                    <li key={device.id} className="mb-2">
                        {device.name} - {device.status}
                    </li>
                ))}
            </ul>
        </div>
    );
}