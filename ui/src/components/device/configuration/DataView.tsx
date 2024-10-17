import { useState } from 'react'
import { cn } from "@/lib/utils"
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"

interface JsonViewProps {
  data: any
  level?: number
}

function JsonViewRaw({ data, level = 0 }: JsonViewProps) {
  if (data === null || data === undefined) {
    return <span className="text-red-500">null</span>;
  }

  if (typeof data !== 'object') {
    return <span className="text-green-500">{JSON.stringify(data)}</span>;
  }

  const isArray = Array.isArray(data);

  return (
    <div style={{ marginLeft: `${level * 20}px` }}>
      {isArray ? '[' : '{'}
      {Object.entries(data).map(([key, value], index) => (
        <div key={key} className={cn("my-1", level === 0 && "bg-muted p-2 rounded-md")}>
          {!isArray && <span className="text-blue-500">&quot;{key}&quot;</span>}
          {!isArray && ': '}
          <JsonViewRaw data={value} level={level + 1} />
          {index < Object.entries(data).length - 1 && ','}
        </div>
      ))}
      {isArray ? ']' : '}'}
    </div>
  );
}

function HumanReadableView({ data }: { data: any }) {
  if (data === null || data === undefined) {
    return <span className="text-red-500">null</span>;
  }

  if (typeof data !== 'object') {
    return <span className="text-green-500">{JSON.stringify(data)}</span>;
  }

  const renderValue = (value: any) => {
    if (value === null || value === undefined) {
      return <span className="text-red-500">null</span>;
    }
    if (typeof value === 'object') {
      return JSON.stringify(value);
    }
    return String(value);
  }

  return (
    <div className="space-y-2">
      {Object.entries(data).map(([key, value]) => (
        <div key={key} className="flex items-start">
          <span className="font-medium text-blue-500 min-w-[150px]">{key}:</span>
          <span className="text-green-500">{renderValue(value)}</span>
        </div>
      ))}
    </div>
  );
}

export function JsonView({ data }: JsonViewProps) {
  const [isJsonView, setIsJsonView] = useState(false);

  if (data === null || data === undefined) {
    return (
      <div className="bg-muted p-4 rounded-md overflow-x-auto">
        <span className="text-red-500">null</span>
      </div>
    );
  }

  return (
    <div>
      <div className="flex items-center justify-end space-x-2 mb-4">
        <Label htmlFor="json-view-toggle">JSON View</Label>
        <Switch
          id="json-view-toggle"
          checked={isJsonView}
          onCheckedChange={setIsJsonView}
        />
      </div>
      <div className="bg-muted p-4 rounded-md overflow-x-auto">
        {isJsonView ? <JsonViewRaw data={data} /> : <HumanReadableView data={data} />}
      </div>
    </div>
  );
}