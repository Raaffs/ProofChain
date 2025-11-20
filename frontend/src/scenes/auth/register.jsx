import { ChangeEvent,useEffect,useState } from "react";
import { 
    Grid,Box,TextField,
    Typography, Button,
    Avatar, Link,
    Card,CardActions, 
    InputLabel,Select,
    MenuItem
} from "@mui/material";
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { Login,Register2 } from "../../../wailsjs/go/main/App";
import {useNavigate} from "react-router-dom"
import { useTheme } from "@emotion/react";
import logo from '../../assets/images/bg7.jpg'
import bg from '../../assets/images/Untitled.png'
function RegisterUser({setAuthStatus}) {
    const theme=useTheme()
    const btnstyle={margin:'50px 0',width:'200px'}
    let registerAsVerifier=false
    const navigate =useNavigate()
    const [input, setInput] = useState({
        privateKey: "",
        username: "",
        password: ""
    });
    const [error, setError] = useState(null);
    const handleClick = () => {
        Register2(input.privateKey, input.username, input.password,registerAsVerifier)
        .then(() => {
            setAuthStatus(true)
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
    const setUserType=(Event)=>{
        Event.target.value==='verifier'?registerAsVerifier=true:registerAsVerifier=false
        console.log(registerAsVerifier)
    }
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
                {error && (
                  <Typography color="error" align="center" style={{ marginBottom: '16px' }}>
                    {error}
                  </Typography>
                )}
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
              <InputLabel >Register As</InputLabel>    
                <Select
                    onChange={setUserType}
                >
                    <MenuItem value='user'>User</MenuItem>
                    <MenuItem value='verifier'>Verifier</MenuItem>
                </Select>
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