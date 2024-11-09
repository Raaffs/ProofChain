import React from 'react';
import { ResponsivePie } from '@nivo/pie';
import { tokens } from '../themes';
import { useTheme } from '@mui/material';
const StatusBreakdownChart = ({ data }) => {
  const theme=useTheme()
  const colors=tokens(theme.palette.mode)
  const statusCounts = data.reduce((acc, doc) => {
    const statusLabel =
      doc.Stats == 0 ? 'Approved' :
      doc.Stats == 1 ? 'Rejected' :
      doc.Stats == 2 ? 'Pending' : 'Unknown';

    acc[statusLabel] = (acc[statusLabel] || 0) + 1;
    return acc;
  }, {});

  // Transforming statusCounts into a format suitable for ResponsivePie
  const chartData = Object.entries(statusCounts).map(([status, count]) => ({
    id: status,
    label: status,
    value: count,
  }));

  // Custom color scheme based on status
  const colorMapping = {
    Approved: colors.greenAccent[500],  // Green
    Rejected: 'red',  // Red
    Pending: colors.blueAccent[300],   // Orange
    Unknown: '#9e9e9e',   // Grey for unknown status
  };

  return (
    <div style={{ height: 400 }}>
      <ResponsivePie
        data={chartData}
        margin={{ top: 40, right: 80, bottom: 80, left: 80 }}
        innerRadius={0.5} // Adjust for donut effect; set to 0 for solid pie
        padAngle={0.7}
        cornerRadius={3}
        colors={(datum) => colorMapping[datum.id]} // Apply custom colors
        borderWidth={1}
        // borderColor={{ from: 'color', modifiers: [['darker', 0.2]] }}
        arcLabel="value"
        arcLabelsSkipAngle={10}
        arcLabelsTextColor="#000000"
        legends={[
          {
            anchor: 'bottom',
            direction: 'row',
            justify: false,
            translateX: 0,
            translateY: 56,
            itemsSpacing: 0,
            itemWidth: 100,
            itemHeight: 18,
            itemTextColor: '#999',
            symbolSize: 18,
            symbolShape: 'circle',
            effects: [
              {
                on: 'hover',
                style: {
                  itemTextColor: '#000'
                }
              }
            ]
          }
        ]}
        tooltip={({ datum }) => (
          <div
            style={{
              padding: '6px 12px',
              color: '#ffffff', // Set text color for tooltip
              background: '#333333', // Set background for tooltip
              borderRadius: '4px'
            }}
          >
            <strong>{datum.id}</strong>: {datum.value}
          </div>
        )}
      />
    </div>
  );
};

export default StatusBreakdownChart;
