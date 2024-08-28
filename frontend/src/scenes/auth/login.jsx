import { ChangeEvent,useEffect,useState } from "react";
import { Grid,Box,Paper,TextField,Typography, Button,Avatar, Link,Card,CardContent,CardActions, CardMedia} from "@mui/material";
import {FormControlLabel} from "@mui/material";
import Header from "../../components/Header"
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { Login,Register } from "../../../wailsjs/go/main/App";
import useMediaQuery from "@mui/material/useMediaQuery";
import { tokens } from "../../themes";
import {Formik, Form} from 'formik'
import * as yup from "yup"
import { useTheme } from "@emotion/react";
import {useNavigate} from "react-router-dom"
import logo from '../../assets/images/bg8.jpg'
import bg from '../../assets/images/Untitled.png'
import PopUp from "../../components/PopUp";
const Redirect=({to})=>{
    return(
        <Link to={to}/>
    )
}
function LoginUser() {
    const theme=useTheme()
    const btnstyle={margin:'50px 0',width:'200px'}
    const navigate =useNavigate()
    const [input, setInput] = useState({
        username: "",
        password: ""
    });
    const [error, setError] = useState(null);

    const handleClick = () => {
        Login(input.username, input.password)
            .then(() => {
                navigate('/dashboard')
            })
            .catch((err) => {
                setError(err)
                console.log(err);
            });
    };

    const handleChange = (event) => {
        setInput({
            ...input,
            [event.target.name]: event.target.value
        });
    };
    return (
        <Box 
            display="flex" 
            padding={2}
            justifyContent="center" 
            alignItems="center"
            sx={{ backgroundColor: 'transparent', backgroundSize: '100% 100%', backgroundImage:`url(${bg})`, backgroundRepeat:'no-repeat',width:"dynamic",height:"86vh", borderRadius:"20px", margin:'5px'}}
        > 
        <Card sx={{width:"50%", height:"dynamic", minHeight:'60%', borderRadius:"16px",backgroundImage:`url(${bg})`,backgroundColor:'transparent'}} elevation="20">
            <Grid align='center' >
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main',top:'0px' }}>
                    <LockOutlinedIcon />
                </Avatar>
                <h1>LogIn</h1>
            </Grid>
        <CardActions sx={{backgroundColor:'transparent', }}>
            <Box display="flex"flexDirection="column" width="100%"  >
                {error && (
                  <Typography color="error" align="center" style={{ marginBottom: '10px' }}>
                    {error}
                  </Typography>
                )}
                <TextField 

                    label='Username' 
                    placeholder='Enter username' 
                    variant="outlined" 
                    name="username"
                    value={input.username}
                    onChange={handleChange}
                    fullWidth 
                    required 
                    style={{ padding:"5px" }}
                />
                <TextField 
                    label='Password' 
                    placeholder='Enter password' 
                    name="password"
                    value={input.password}
                    type='password' 
                    onChange={handleChange}
                    variant="outlined" 
                    fullWidth 
                    required 
                    style={{ padding:"5px" }}

                />
                    <Box display="flex" justifyContent="center" alignItems="center">
                        <Button 
                            type='submit' 
                            size="medium" 
                            variant="contained" 
                            borderRadius="100px"
                            style={btnstyle}
                            onClick={handleClick}
                        >
                            Log In
                        </Button>
                    </Box>
                    <Box>
                    <Typography variant="h5"> Don't have an account ?
            </Typography>
            <Redirect 
                to="/register" 
            />
                </Box>        
            </Box>
            
        </CardActions>    
    </Card>
    <Card sx={{padding:"10px" ,width:"100%", height:"dynamic",minHeight:"60%", backgroundImage:`url(${bg})`, borderRadius:"16px",marginLeft:"20px", backgroundRepeat:'no-repeat',backgroundColor:'transparent'}} 
        elevation="20" 
    
        justifyContent="right"
        alignItems="center"
    >
    </Card>
    </Box>
      );
}

export default LoginUser