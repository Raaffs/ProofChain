import './App.css';
import { ColorModeContext, useMode } from './themes';
import { CssBaseline , ThemeProvider} from '@mui/material';
import {Route,Routes} from "react-router-dom"
import Topbar from './scenes/global/Topbar';
import Sidebar from './scenes/global/Sidebar';
import Dashboard from './scenes/dashboard';
import RegisterUser from './scenes/auth/register';
import LoginUser from './scenes/auth/login';
import LogoutUser from './scenes/auth/logout';
import ApprovedDocuments from './scenes/transactions/approved';
import RejectedDocuments from './scenes/transactions/rejected';
import PendingDocuments from './scenes/transactions/pending';
import UploadDocs from './scenes/upload'
import { useState } from 'react';
import Issue from './scenes/issue/issue';
// import Documents from './scenes/documents';
// import Line from './scenes/line';
// import Pie from './scenes/pie';
// import Form from './scenes/form';

function App() {
    const [theme,colorMode]=useMode();
    const [authStatus,setAuthStatus]=useState(false)
    return(
        <ColorModeContext.Provider value={colorMode}>
            <ThemeProvider theme={theme}>
                <CssBaseline/>
                 <div className='app'>
                    <Sidebar authStatus={authStatus} setAuthStatus={setAuthStatus}/>
                    <main className="content">
                        <Topbar/>     
                            <Routes>
                                <Route path="/" element={<LoginUser setAuthStatus={setAuthStatus}/>}/>
                                <Route path="/register" element={<RegisterUser setAuthStatus={setAuthStatus}/>}/>
                                <Route path="/logout" element={<LogoutUser setAuthStatus={setAuthStatus}/>} />
                                <Route path="/dashboard" element={<Dashboard/>}/>
                                <Route path="/documents/approved" element={<ApprovedDocuments/>}/>
                                <Route path="/documents/rejected" element={<RejectedDocuments/>}/>
                                <Route path="/documents/pending" element={<PendingDocuments/>}/>
                                <Route path="/documents/upload" element={<UploadDocs/>}/>
                                <Route path="/documents/issue" element={<Issue/>}/>
                            </Routes>
                    </main>    
                 </div>
            </ThemeProvider>
        </ColorModeContext.Provider>
   )
}

export default App
