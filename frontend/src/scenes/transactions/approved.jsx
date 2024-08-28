import { Box, Typography, useTheme } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import AdminPanelSettingsOutlinedIcon from "@mui/icons-material/AdminPanelSettingsOutlined";
import LockOpenOutlinedIcon from "@mui/icons-material/LockOpenOutlined";
import SecurityOutlinedIcon from "@mui/icons-material/SecurityOutlined";
import Header from "../../components/Header";
import { WidthFull } from "@mui/icons-material";
import { GetAcceptedDocs, GetAllDocs } from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
const ApprovedDocuments=()=>{
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);
    const [docs,setDocs]=useState([])
    const [error,setError]=useState(null);
    useEffect(() => {
      const getDocuments = () => {
        GetAcceptedDocs()
        .then((result) => {
            if (result===null){
              setDocs([{
                "ID":"",
                "Requester":"",
                "Verifier":"",
                "Name":"",
                "Desc":"",
                "IpfsAddress":""
              }])
              setError("No Verified Documents")
              }else{
              const updatedDocs = result.map((doc) => {
                 if (doc.IpfsAddress === '') {
                   doc.IpfsAddress = 'private';
                 }
                 return doc;
            });
              setDocs(updatedDocs);
            }
          })
          .catch((err) => {
            setError(err.message);
          });
      };
      getDocuments();
    }, []); // Empty dependency array ensures this runs once on mount
    const columns=[
        {"field":"Requester",headerName:"Requester",flex:1},
        {"field":"Verifier",headerName:"Verifier",flex:1},
        {"field":"Name",headerName:"Name",flex:1},
        {"field":"Desc",headerName:"Description",flex:1},
        {"field":"IpfsAddress",headerName:"Ipfs Address",flex:1},
    ]
    return (
        <Box m="20px"
          sx={{width:'dynamic',maxWidth:'95%',justifyContent:'center'}}
        >
            <Header title="Approved Documents"></Header>
                            {error && (
                  <Typography color="error" align="center" style={{ marginBottom: '16px' }}>
                    {error}
                  </Typography>
                )}

            <Box
                m="40px 0 0 0"
                height="70vh"
                justifyContent="center" 
                sx={{
                  "& .MuiDataGrid-root": {
                    border: "none",
                  },
                  "& .MuiDataGrid-cell": {
                    borderBottom: "none",
                    fontSize:'1.1rem'
                  },
                  "& .name-column--cell": {
                    color: colors.greenAccent[300],
                  },
                  "& .MuiDataGrid-columnHeaders": {
                    backgroundColor: colors.blueAccent[700],
                    borderBottom: "none",
                    fontSize:"1.2rem"
                  },
                  "& .MuiDataGrid-virtualScroller": {
                    // backgroundColor: colors.blueAccent[900],
                  },
                  "& .MuiDataGrid-footerContainer": {
                    borderTop: "none",
                    backgroundColor: colors.blueAccent[900],
                  },
                  "& .MuiCheckbox-root": {
                    color: `${colors.greenAccent[200]} !important`,
                  },
                }}
            >
                <DataGrid 
                columns={columns}
                rows={docs}
                getRowId={(row)=>{return row.Requester}} 
                sx={{width:"dynamic", maxWidth:"170vh"}}
                >
                </DataGrid>
            </Box>
        </Box>
    )
}

export default ApprovedDocuments