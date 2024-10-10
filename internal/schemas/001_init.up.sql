CREATE TABLE devices (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    status TEXT NOT NULL
);

INSERT INTO devices (id, name, type, status) VALUES
('1', 'Living Room Light', 'light', 'on'),
('2', 'Bedroom Light', 'light', 'off'),
('3', 'Kitchen Light', 'light', 'idle'),
('4', 'Bedroom Thermostat', 'thermostat', 'idle'),
('5', 'Kitchen Thermostat', 'thermostat', 'idle'); 

CREATE TABLE configurations (
    id TEXT PRIMARY KEY,
    device_id TEXT NOT NULL,
    name TEXT NOT NULL,
    active BOOLEAN NOT NULL,
    data JSON
);

INSERT INTO configurations (id, device_id, name, active, data) VALUES
('1', '1', 'Daymode', true, "{'brightness': 100}"),
('2', '1', 'Nightmode', false, "{'brightness': 0}"),
('3', '2', 'Evening', true, "{'brightness': 75, 'color_temp': 300}"),
('4', '2', 'Morning', false, "{'brightness': 50, 'color_temp': 200}"),
('5', '3', 'Nightmode', true, "{'brightness': 0}");