import React from 'react';
import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';
import { tokens } from '../themes';
import { useTheme } from '@emotion/react';
ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const InstituteStatusChart = ({ data }) => {
  const theme=useTheme();
  const colors=tokens(theme.palette.mode)
  const statuses = { 0: 'Approved', 1: 'Rejected', 2: 'Pending' };

  // Aggregate data by Institute and status
  const aggregatedData = data.reduce((acc, doc) => {
    const institute = doc.Institute;
    const status = statuses[doc.Stats] || 'Unknown';
    
    if (!acc[institute]) {
      acc[institute] = { Approved: 0, Rejected: 0, Pending: 0 };
    }
    acc[institute][status] += 1;

    return acc;
  }, {});

  // Prepare data for Chart.js
  const chartData = {
    labels: Object.keys(aggregatedData), // Institute names
    datasets: [
      {
        label: 'Approved',
        data: Object.values(aggregatedData).map((entry) => entry.Approved),
        backgroundColor: colors.greenAccent[500], // Green for Approved
        categoryPercentage:1,
        barPercentage:0.5,
        borderWidth:1,
      },
      {
        label: 'Rejected',
        data: Object.values(aggregatedData).map((entry) => entry.Rejected),
        backgroundColor: 'red', // Red for Rejected
        categoryPercentage:1,
        barPercentage:0.5,
        borderWidth:1,
      },
      {
        label: 'Pending',
        data: Object.values(aggregatedData).map((entry) => entry.Pending),
        backgroundColor: colors.blueAccent[400], // Orange for Pending
        categoryPercentage:1,
        barPercentage:0.5,
        borderWidth:1,
      },
    ],
  };

  const options = {
    responsive: true,
    indexAxis: 'y', // Switch to horizontal layout
    plugins: {
      legend: { position: 'top' },
      tooltip: {
        callbacks: {
          label: (context) => `${context.dataset.label}: ${context.raw}`,
        },
      },
    },
    scales: {
      x: {
        stacked: true,
        title: { display: true, text: 'Document Count' },
        beginAtZero: true  ,
        categoryPercentage: 1, 
        barPercentage: 1,      // Adjust bar width to fit the reduced category width
 
      },
      y: {
        stacked: true,
        title: { display: true, text: 'Institute' },
        categoryPercentage: 1, 
        barPercentage: 1,      // Adjust bar width to fit the reduced category width
      },
    },
    // barThickness: 30, // Adjust this value for desired bar width
  };

  return (
    <div style={{ width: '100%', height: '400px' }}>
      {/* Horizontal Stacked Bar Chart */}
      <Bar data={chartData} options={options} />
    </div>
  );
};

export default InstituteStatusChart;
