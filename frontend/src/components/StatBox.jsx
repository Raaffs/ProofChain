import { Box, Typography, useTheme, Button } from "@mui/material";
import { tokens } from "../themes";
import { useNavigate } from "react-router-dom";
const StatBox = ({ title, subtitle, icon, fontColor, path }) => {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const navigate=useNavigate()
  return (
    <Box width="100%" m="0 30px">
      <Box display="flex" justifyContent="center" alignItems="center">
        <Button
          onClick={()=>{navigate(path)}}
        >
          <Typography
            variant="h4"
            fontWeight="bold"
            sx={{ color: fontColor }}
          >
            {title}
          </Typography>
        </Button>
        <Box>
          <Box display="flex" justifyContent="flex-end">
            {icon}
          </Box>
        </Box>
      </Box>
      <Box display="flex" justifyContent="center" mt="2px">
        <Typography variant="h5" sx={{ color: fontColor }}>
          {subtitle}
        </Typography>
        <Typography
          variant="h5"
          fontStyle="italic"
          sx={{ color: colors.greenAccent[600] }}
        >
        </Typography>
      </Box>
    </Box>
  );
  
};

export default StatBox;