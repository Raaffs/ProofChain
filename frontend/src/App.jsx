import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import { ColorModeContext, useMode } from './themes';
import { Css } from '@mui/icons-material';
import { CssBaseline , ThemeProvider} from '@mui/material';
import {Route,Routes} from "react-router-dom"
import Topbar from './scenes/global/Topbar';
import Sidebar from './scenes/global/Sidebar';
import Dashboard from './scenes/dashboard';
import RegisterUser from './scenes/auth/register';
// import Documents from './scenes/documents';
// import Line from './scenes/line';
// import Pie from './scenes/pie';
// import Form from './scenes/form';

function App() {
   const [theme,colorMode]=useMode();
   return(
        <ColorModeContext.Provider value={colorMode}>
            <ThemeProvider theme={theme}>
                <CssBaseline/>
                 <div className='app'>
                    <Sidebar/>
                    <main className="content">
                        <Topbar/>
                        <Routes>
                             <Route path="/" element={<Dashboard/>}/> 
                             <Route path="/register" element={<RegisterUser/>}/>
                            {/* 
                            <Route path="/documents/approved" element={<Dashboard/>}/>
                            <Route path="/documents/pending" element={<Dashboard/>}/>
                            <Route path="/documents/rejected" element={<Dashboard/>}/>
                            <Route path="/transactions" element={<Dashboard/>}/>
                            <Route path="/transactions/approved" element={<Dashboard/>}/>
                            <Route path="/transactions/rejected" element={<Dashboard/>}/>
                            <Route path="/transactions/pending" element={<Dashboard/>}/>
                            <Route path="/upload" element={<Dashboard/>}/>
                            <Route path="/pie" element={<Pie/>}/> */}
                        </Routes>
                    </main>
                 </div>
            </ThemeProvider>
        </ColorModeContext.Provider>
   )
}

export default App
