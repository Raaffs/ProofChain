import { Box } from "@mui/material";
import Header from "../../components/Header";
import { useTheme } from "@emotion/react";
import { tokens } from "../../themes";
const Dashboard=()=>{
    const theme=useTheme()
    const colors=tokens(theme.palette.mode)
    return(
        <Box sx={{background:colors.primary[700], height:'100vh',borderRadius:'15px',margin:'10px' }}>
            <Box display="flex" justifyContent="space-between" alignItems="center">
            </Box>
        </Box>
    )
}

export default Dashboard