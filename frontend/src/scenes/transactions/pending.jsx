import { Box, Typography, useTheme } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import Header from "../../components/Header";
import { GetPendingDocuments } from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import { IsApprovedInstitute } from "../../../wailsjs/go/main/App";
const PendingDocuments=()=>{
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);
    const [docs,setDocs]=useState([])
    const [error,setError]=useState(null);
    const handleApprove = (id) => {
      console.log('Approved ID:', id);
      return 0;
    };
  
    const handleReject = (id) => {
      console.log('Rejected ID:', id);
      return 1;
    };
    useEffect(() => {
        GetPendingDocuments()
        .then((result) => {
            console.log(" result : ",result)
            if (result===null){
              setDocs([{
                "ID":"",
                "Requester":"",
                "Verifier":"",
                "Name":"",
                "Desc":"",
                "IpfsAddress":""
              }])
              setError("No Pending Documents")
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
            setError(err.message );
          });
          IsApprovedInstitute()
          .then((result)=>{
            if (result==true){
              columns.push({
                field:"verify",
                headerName: "Verify",
                flex: 1,
                renderCell: (params) => (
                  <Box>
                    <Button
                      variant="contained"
                      color="success"
                      onClick={() => handleApprove(params.row.ID)}
                      style={{ marginRight: '10px' }}
                    >
                      Approve
                    </Button>
                    <Button
                      variant="contained"
                      color="error"
                      onClick={() => handleReject(params.row.ID)}
                    >
                      Reject
                    </Button>
                  </Box>
                ),
              })
              .catch((error)=>{console.log(error)})
            }
          })

    }, []); // Empty dependency array ensures this runs once on mount
    const columns=[
        {"field":"Requester",headerName:"Requester",flex:1},
        {"field":"Institute",headerName:"Institute",flex:1},
        {"field":"Name",headerName:"Name",flex:1},
        {"field":"Desc",headerName:"Description",flex:1},
        {"field":"IpfsAddress",headerName:"Ipfs Address",flex:1},
    ]
    return (
        <Box m="20px"
          sx={{width:'dynamic',maxWidth:'95%',justifyContent:'center'}}
        >
            <Header title="Rejected Documents"></Header>
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
                    color: `${colors.greenAccent[200]} !impor       tant`,
                  },
                }}
            >
                <DataGrid 
                columns={columns}
                rows={docs}
                getRowId={(row)=>{return row.ID}} 
                sx={{width:"dynamic", maxWidth:"170vh"}}
                >
                </DataGrid>
            </Box>
        </Box>
    )
}

export default PendingDocuments