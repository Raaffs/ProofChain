import { Box, Typography, useTheme, Button } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import Header from "../../components/Header";
import { GetPendingDocuments, IsApprovedInstitute, ApproveDocument } from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import PopUp from "../../components/PopUp";
        const PendingDocuments = () => {
            const theme = useTheme();
            const colors = tokens(theme.palette.mode);
            const [docs, setDocs] = useState([]);
            const [error, setError] = useState(null);
            const [message,setMessage]=useState("")
            const [columns, setColumns] = useState([
                { field: "Requester", headerName: "Requester", flex: 1 },
                { field: "Verifier", headerName: "Verifier", flex: 1 }, // Adjusted this field
                { field: "Name", headerName: "Name", flex: 1 },
                { field: "ShaHash", headerName: "Hash", flex: 1 },
                { field: "IpfsAddress", headerName: "IPFS Address", flex: 1 }
            ]);

            const handleApprove = (id) => {
                let result=docs.find((doc)=>doc.ID===id)
                console.log("result : ",docs)
                console.log("result : ",result)
                console.log("document hash:",result.ShaHash)
                ApproveDocument(0,result.ShaHash).then(()=>{
                    setMessage("Document verified successfully")
                    console.log("approved successfully")
                })
                .catch((err)=>{
                    setError(err)
                    console.log("error approving:",err)
                })
            };

    const handleReject = (id) => {
        const hash =docs[id].ShaHash
        ApproveDocument(1,hash).then(()=>{
            setMessage("Document verified successfully")
        })
        .catch((err)=>{
            setError(err)
        })
        console.log('Approved ID:', id);
    };

    useEffect(() => {
        const fetchDocuments = async () => {
            try {
                const result = await GetPendingDocuments();
                if (!result || result.length === 0) {
                    setDocs([{
                        ID: "",
                        Requester: "",
                        Verifier: "",  
                        Name: "",
                        ShaHash: "",
                        IpfsAddress: ""
                    }]);
                    setError("No Pending Documents");
                } else {
                    const updatedDocs = result.map((doc) => ({
                        ...doc,
                        IpfsAddress: doc.IpfsAddress === '' ? 'private' : doc.IpfsAddress,
                        ShaHash: doc.ShaHash  
                    }));
                    setDocs(updatedDocs);
                    console.log("docs ",docs)
                }
            } catch (err) {
                setError(err.message);
            }

            try {
                const isApproved = await IsApprovedInstitute();
                if (isApproved) {
                    setColumns((prevColumns) => [
                        ...prevColumns,
                        {
                            field: "verify",
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
                        },
                    ]);
                }
            } catch (err) {
                console.log(err.message);
            }
        };

        fetchDocuments();
    },[]);

    return (
        <Box m="20px" sx={{ width: 'dynamic', maxWidth: '95%', justifyContent: 'center' }}>
            <Header title="Pending Documents" />
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
                        fontSize: '1.1rem',
                    },
                    "& .name-column--cell": {
                        color: colors.greenAccent[300],
                    },
                    "& .MuiDataGrid-columnHeaders": {
                        backgroundColor: colors.blueAccent[700],
                        borderBottom: "none",
                        fontSize: "1.2rem",
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
                    getRowId={(row) => row.ID} // Use `ID` as a unique identifier
                    sx={{ width: "dynamic", maxWidth: "170vh" }}
                />
            </Box>
        </Box>
    );
};

export default PendingDocuments;
