import { Box } from "@mui/material";
import IssueCard from "../../components/cards/certificate";
import Header from "../../components/Header";

const Issue = () => {
  return (
    <Box
      sx={{
        minHeight: "100vh", // Keeps children stacked vertically
        display: "flex",
        flexDirection: "column",
        justifyContent: "flex-start", // Aligns children to the top (main axis)
        
        // The fix is here:
        alignItems: "flex-start", // Aligns children to the left (cross axis)
        
        p: 2,
        width: "100%", // Ensure the Box takes full width to see the alignment
      }}
    >
      <Header title="Issue Certificate" subtitle="" />
      <IssueCard 
        data={null} 
        viewTitle="Issue Certificate"
      />
    </Box>
  );
};

export default Issue;