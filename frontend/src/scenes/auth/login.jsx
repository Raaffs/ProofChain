import { useState } from "react";
import {
  Grid,
  Box,
  Paper,
  TextField,
  Typography,
  Button,
  Avatar,
  Link,
  Card,
  CardContent,
  CardActions,
  CardMedia,
} from "@mui/material";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import { Login } from "../../../wailsjs/go/main/App";
import { useTheme } from "@emotion/react";
import { useNavigate } from "react-router-dom";
import bg from "../../assets/images/Untitled.png";
import { btnstyle, textFieldSx } from "../../styles/styles";
import { tokens } from "../../themes";
import { PlatformOverviewCard } from "../../components/template/template";
const Redirect = ({ to }) => {
  return <Link to={to} />;
};
function LoginUser({ setAuthStatus }) {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);

  const navigate = useNavigate();
  const [input, setInput] = useState({
    username: "",
    password: "",
  });
  const [error, setError] = useState(null);

  const handleClick = () => {
    Login(input.username, input.password)
      .then(() => {
        setAuthStatus(true);
        navigate("/dashboard");
      })
      .catch((err) => {
        setError(err);
        console.log(err);
      });
  };

  const handleChange = (event) => {
    setInput({
      ...input,
      [event.target.name]: event.target.value,
    });
  };
  return (
    <Box
      display="flex"
      padding={2}
      justifyContent="center"
      alignItems="center"
      sx={{
        backgroundSize: "100% 100%",
        backgroundRepeat: "no-repeat",
        width: "dynamic",
        height: "86vh",
        borderRadius: "20px",
        margin: "5px",
      }}
    >
      <Card
        sx={{
          width: { xs: "90%", sm: "70%", md: "50%" }, // Responsive width
          minHeight: "450px",
          borderRadius: "16px",

          backgroundColor:
            theme.palette.mode == "light"
              ? "rgba(255, 255, 255, 0.9)"
              : "transparent",

          boxShadow: "0 8px 30px rgba(0, 0, 0, 0.2)",
          padding: "20px",
        }}
      >
        {/* Header Section */}
        <Grid align="center" sx={{ mb: 2 }}>
          <Avatar
            sx={{
              m: 1,
              bgcolor: "#FF6F61", // lighter, brighter red
              color: "white", // text/icon stays visible
              top: "0px",
              boxShadow: "0 4px 12px rgba(255, 111, 97, 0.4)", // subtle glow for depth
            }}
          >
            <LockOutlinedIcon />
          </Avatar>
          <Typography
            variant="h4"
            component="h1"
            sx={{
              fontWeight: 600,
              color: theme.palette.mode == "light" ? "black" : "white",
            }}
          >
            Log In
          </Typography>
        </Grid>

        {/* Form/Action Section */}
        <Box sx={{ p: { xs: 1, sm: 2 } }}>
          <Box display="flex" flexDirection="column" gap={2} width="100%">
            {/* Error Message */}
            {error && (
              <Typography color="error" align="center" sx={{ mb: 1 }}>
                {error}
              </Typography>
            )}
            {/* Username Field */}
            <TextField
              label="Username"
              placeholder="Enter username"
              variant="standard"
              name="username"
              value={input.username}
              onChange={handleChange}
              fullWidth
              required
              sx={textFieldSx} // Applying custom sleek style
            />
            {/* Password Field */}
            <TextField
              label="Password"
              placeholder="Enter password"
              name="password"
              value={input.password}
              type="password"
              onChange={handleChange}
              variant="standard" // Sleek underline variant
              fullWidth
              required
              sx={textFieldSx} // Applying custom sleek style
            />{" "}
            {/* Button */}
            <Box
              display="flex"
              justifyContent="center"
              alignItems="center"
              sx={{ mt: 3, mb: 2 }}
            >
              <Button
                type="submit"
                size="medium"
                variant="contained"
                borderRadius="100px"
                style={btnstyle} // Applying the red/orange gradient style
                onClick={handleClick}
              >
                Log In
              </Button>
            </Box>
            {/* Redirect Link */}
            <Box align="center" sx={{ mt: 2 }}>
              <Typography variant="body1" sx={{ color: "#555" }}>
                Don't have an account?
              </Typography>
              <Typography
                component="a"
                href="/register"
                sx={{
                  color: "#E94057", // Matching the button/field color
                  fontWeight: 600,
                  textDecoration: "none",
                  "&:hover": { textDecoration: "underline" },
                }}
              >
                Sign Up
              </Typography>
              {/* <Redirect to="/register" /> - Replace with the Typography link above */}
            </Box>
          </Box>
        </Box>
      </Card>
     <PlatformOverviewCard/>
    </Box>
  );
}

export default LoginUser;
