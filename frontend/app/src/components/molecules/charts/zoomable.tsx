import { useState, useMemo, useRef } from 'react';
import {
  CartesianGrid,
  XAxis,
  YAxis,
  ReferenceArea,
  ResponsiveContainer,
  Bar,
  BarChart,
} from 'recharts';
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from '@/components/ui/chart';
import { capitalize } from '@/lib/utils';

export type DataPoint<T extends string> = Record<T, number> & {
  date: string;
};

const getNextActiveLabel = (activeLabel: string, data: DataPoint<string>[]) => {
  const currentIndex = data.findIndex((d) => d.date === activeLabel);
  if (currentIndex === -1) {
    return null;
  }

  // if we're at the end of the data, determine the time between the last two data points and add that to the last date
  if (currentIndex === data.length - 1) {
    const lastDate = new Date(data[currentIndex].date);
    const secondLastDate = new Date(data[currentIndex - 1].date);
    const diff = lastDate.getTime() - secondLastDate.getTime();
    return new Date(lastDate.getTime() + diff).toISOString();
  }

  return data[currentIndex + 1]?.date || activeLabel;
};

const getPrevActiveLabel = (activeLabel: string, data: DataPoint<string>[]) => {
  const currentIndex = data.findIndex((d) => d.date === activeLabel);
  if (currentIndex === -1) {
    return activeLabel;
  }

  // if we're at the start of the data, determine the time between the first two data points and subtract that from the first date
  if (currentIndex === 0) {
    const firstDate = new Date(data[currentIndex].date);
    const secondDate = new Date(data[currentIndex + 1].date);
    const diff = secondDate.getTime() - firstDate.getTime();
    return new Date(firstDate.getTime() - diff).toISOString();
  }

  return data[currentIndex - 1]?.date || activeLabel;
};

type ZoomableChartProps<T extends string> = {
  data: DataPoint<T>[];
  colors?: Record<string, string>;
  zoom?: (startTime: string, endTime: string) => void;
  showYAxis?: boolean;
};

export function ZoomableChart<T extends string>({
  data,
  colors,
  zoom,
  showYAxis = true,
}: ZoomableChartProps<T>) {
  const [refAreaLeft, setRefAreaLeft] = useState<string | null>(null);
  const [refAreaRight, setRefAreaRight] = useState<string | null>(null);
  const [actualRefAreaLeft, setActualRefAreaLeft] = useState<string | null>(
    null,
  );
  const [actualRefAreaRight, setActualRefAreaRight] = useState<string | null>(
    null,
  );
  const [isSelecting, setIsSelecting] = useState(false);
  const chartRef = useRef<HTMLDivElement>(null);

  const chartConfig = useMemo<ChartConfig>(() => {
    const keys = Object.keys(data[0] || {}).filter((key) => key !== 'date');
    return keys.reduce<ChartConfig>((acc, key, index) => {
      acc[key] = {
        label: capitalize(key),
        color: colors?.[key] || `hsl(${(index * 360) / keys.length}, 70%, 50%)`,
      };
      return acc;
    }, {});
  }, [data, colors]);

  const handleMouseDown = (e: any) => {
    if (e.activeLabel) {
      setRefAreaLeft(e.activeLabel);
      setActualRefAreaLeft(getPrevActiveLabel(e.activeLabel, data));
      setIsSelecting(true);
    }
  };

  const handleMouseMove = (e: any) => {
    if (isSelecting && e.activeLabel) {
      setRefAreaRight(e.activeLabel);
      setActualRefAreaRight(getNextActiveLabel(e.activeLabel, data));
    }
  };

  const handleMouseUp = () => {
    if (actualRefAreaLeft && actualRefAreaRight) {
      const [left, right] = [actualRefAreaLeft, actualRefAreaRight].sort();
      zoom?.(left, right);
    }
    setRefAreaLeft(null);
    setActualRefAreaLeft(null);
    setRefAreaRight(null);
    setActualRefAreaRight(null);
    setIsSelecting(false);
  };

  const minDate = new Date(
    Math.min(...data.map((d) => new Date(d.date).getTime())),
  );
  const maxDate = new Date(
    Math.max(...data.map((d) => new Date(d.date).getTime())),
  );

  const formatXAxis = (tickItem: string) => {
    const date = new Date(tickItem);
    const timeDiff = maxDate.getTime() - minDate.getTime();
    const oneDay = 24 * 60 * 60 * 1000;
    const sevenDays = 7 * oneDay;

    if (timeDiff > sevenDays) {
      return date.toLocaleDateString([], { month: 'short', day: 'numeric' });
    } else if (timeDiff > oneDay) {
      return `${date.toLocaleDateString([], { month: 'short', day: 'numeric' })} ${date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`;
    } else {
      return date.toLocaleTimeString([], {
        hour: '2-digit',
        minute: '2-digit',
      });
    }
  };

  // remove date from dataKeys
  const dataKeys = Object.keys(data[0] || {}).filter((key) => key !== 'date');

  return (
    <ChartContainer
      config={chartConfig}
      className="w-full h-[200px] min-h-[200px]"
    >
      <div className="h-full" ref={chartRef} style={{ touchAction: 'none' }}>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={data}
            margin={{
              left: 0,
              right: 0,
              top: 0,
              bottom: 0,
            }}
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="date"
              tickFormatter={formatXAxis}
              tickLine={false}
              axisLine={false}
              tickMargin={4}
              minTickGap={16}
              style={{ fontSize: '10px', userSelect: 'none' }}
            />
            {showYAxis && (
              <YAxis
                tickLine={false}
                axisLine={false}
                tickMargin={4}
                style={{ fontSize: '10px', userSelect: 'none' }}
              />
            )}
            <ChartTooltip
              content={
                <ChartTooltipContent
                  className="w-[150px] sm:w-[200px] font-mono text-xs sm:text-xs"
                  labelFormatter={(value) => new Date(value).toLocaleString()}
                />
              }
            />
            {dataKeys.map((key) => (
              <Bar
                key={key}
                type="monotone"
                dataKey={key}
                stroke={chartConfig[key].color}
                fillOpacity={1}
                fill={chartConfig[key].color}
                isAnimationActive={false}
              />
            ))}

            {refAreaLeft && refAreaRight && (
              <ReferenceArea
                x1={refAreaLeft}
                x2={refAreaRight}
                strokeOpacity={0.3}
                fill="hsl(var(--foreground))"
                fillOpacity={0.1}
              />
            )}
          </BarChart>
        </ResponsiveContainer>
      </div>
    </ChartContainer>
  );
}
