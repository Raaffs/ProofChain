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
import logo from '../../assets/images/bg7.jpg'
import bg from '../../assets/images/Untitled.png'
function RegisterUser() {
    const theme=useTheme()
    const colored = tokens(theme.palette.mode);
    const btnstyle={margin:'50px 0',width:'200px'}
    const isNonMobile = useMediaQuery("(min-width:600px)")

    const [input, setInput] = useState({
        privateKey: "",
        username: "",
        password: ""
    });

    const handleClick = () => {
        Register(input.privateKey, input.username, input.password)
            .then(() => {
                // Registration successful, do something if needed
            })
            .catch((err) => {
                console.log(err);
                // Handle error if needed
            });
    };

    const handleChange = (event) => {
        setInput({
            ...input,
            [event.target.name]: event.target.value
        });
        console.log("data : ", input.privateKey,input.username,input.password)
    };
    return (
        <Box 
            display="flex" 
            padding={2}
            justifyContent="center" 
            alignItems="center"
            sx={{ backgroundColor: 'transparent', backgroundSize: '100% 100%', backgroundImage:`url(${bg})`, backgroundRepeat:'no-repeat',width:"85vw",height:"86vh", borderRadius:"20px", margin:'10px'}}
        > 
        <Card sx={{width:"50%", height:"70%",backgroundColor:`${theme.palette.mode==="dark"?'transparent':'transparent'}`, borderRadius:"16px",backgroundImage:`url(${bg})`}} elevation="20">
            <Grid align='center' >
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main',top:'0px' }}>
                    <LockOutlinedIcon />
                </Avatar>
                <h1>SignUp</h1>
            </Grid>
        <CardActions sx={{backgroundColor:'transparent', }}>
            <Box display="flex"flexDirection="column" width="100%"  >
                <TextField 
                    label='Private Key' 
                    placeholder='Enter private key' 
                    variant="outlined" 
                    fullWidth 
                    required 
                    name="privateKey"
                    value={input.privateKey}
                    onChange={handleChange}
                    InputProps={{
                        style:{
                            borderRadius:"7x",
                            marginBottom:'16px',
                            padding:'0px'
                        }
                    }}
                />
                <TextField 

                    label='Username' 
                    placeholder='Enter username' 
                    variant="outlined" 
                    fullWidth 
                    required 
                    name="username"
                    value={input.username}
                    onChange={handleChange}
                    InputProps={{
                        style:{
                            borderRadius:"7x",
                            marginBottom:'16px',
                            padding:'0px'
                        }
                    }}

                />
                <TextField 
                    label='Password' 
                    placeholder='Enter password' 
                    type='password' 
                    variant="outlined" 
                    fullWidth 
                    required 
                    name="password"
                    value={input.password}
                    onChange={handleChange}
                    InputProps={{
                        style:{
                            borderRadius:"7x",
                            marginBottom:'16px',
                            padding:'0px'
                        }
                    }}

                />
                    <Box display="flex" justifyContent="center" alignItems="center">
                        <Button 
                            type='submit' 
                            size="medium" 
                            variant="contained" 
                            borderRadius="100px"
                            onClick={handleClick}
                            style={btnstyle}
                        >
                            Create Account
                        </Button>
                    </Box>
                    <Box>
                    <Typography variant="h5"> Already have an account ?
                    <Link color="inherit" href="/login" >
                        Login
                </Link>
            </Typography>

                    </Box>
                    
            </Box>
            
        </CardActions>    
    </Card>
    <Card sx={{padding:"10px" ,width:"100%", height:"70%",backgroundColor:`${theme.palette.mode==='dark'?'transparent':'transparent'}`, backgroundImage:`url(${bg})`, borderRadius:"16px",marginLeft:"20px", backgroundRepeat:'no-repeat'}} 
        elevation="20" 
    
        justifyContent="right"
        alignItems="center"
    >
    </Card>
    </Box>
      );
}

export default RegisterUser