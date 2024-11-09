import React from 'react';
import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';

// Register components
ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const InstituteStatusChart = ({ data }) => {
  // Define available statuses
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
        backgroundColor: '#4caf50', // Green for Approved
      },
      {
        label: 'Rejected',
        data: Object.values(aggregatedData).map((entry) => entry.Rejected),
        backgroundColor: 'red', // Red for Rejected
      },
      {
        label: 'Pending',
        data: Object.values(aggregatedData).map((entry) => entry.Pending),
        backgroundColor: '#ff9800', // Orange for Pending
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
        beginAtZero: true
      },
      y: {
        stacked: true,
        title: { display: true, text: 'Institute' }
      }
    },
  };

  return (
    <div style={{ width: '100%', height: '400px' }}>
      {/* Horizontal Stacked Bar Chart */}
      <Bar data={chartData} options={options} />
    </div>
  );
};

export default InstituteStatusChart;
