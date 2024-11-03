import { Box, Typography, useTheme, Button } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import Header from "../../components/Header";
import { GetPendingDocuments, IsApprovedInstitute, ApproveDocument } from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";

const PendingDocuments = () => {
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);
    const [docs, setDocs] = useState([]);
    const [error, setError] = useState(null);
    const [message, setMessage] = useState("");
    const [columns, setColumns] = useState([
        { field: "Requester", headerName: "Requester", flex: 1 },
        { field: "Verifier", headerName: "Verifier", flex: 1 },
        { field: "Name", headerName: "Name", flex: 1 },
        { field: "ShaHash", headerName: "Hash", flex: 1 },
        { field: "IpfsAddress", headerName: "IPFS Address", flex: 1 }
    ]);

      const handleApprove = (id) => {
        setDocs((prevDocs) => {
            const result = prevDocs.find((doc) => doc.ID === id);
            if (!result) {
                setError("Document not found");
                console.log("Document id not found, id: ", id);
                return prevDocs; // Return previous state if not found
            }

            // Approve the document
            ApproveDocument(0, result.ShaHash)
                .then(() => {
                    setMessage("Document approved successfully");
                    // Optionally refresh documents after approval
                    fetchDocuments();
                })
                .catch((err) => setError(err.message));
            
            return prevDocs; // Return previous state
        });
    };

    const handleReject = (id) => {
        setDocs((prevDocs) => {
            const result = prevDocs.find((doc) => doc.ID === id);
            if (!result) {
                console.log("result and docs",prevDocs,result)
                setError("Document not found");
                return prevDocs; // Return previous state if not found
            }

            // Reject the document
            ApproveDocument(1, result.ShaHash)
                .then(() => {
                    setMessage("Document rejected successfully");
                    // Optionally refresh documents after rejection
                    fetchDocuments();
                })
                .catch((err) => setError(err.message));
            
            return prevDocs; // Return previous state
        });
    };

    const setupColumns = () => {
        IsApprovedInstitute()
            .then((isApproved) => {
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
                                        onClick={() => handleApprove(params.id)}
                                        style={{ marginRight: '10px' }}
                                    >
                                        Approve
                                    </Button>
                                    <Button
                                        variant="contained"
                                        color="error"
                                        onClick={() => handleReject(params.id)}
                                    >
                                        Reject
                                    </Button>
                                </Box>
                            ),
                        },
                    ]);
                }
            })
            .catch((err) => console.log(err.message));
    };

    const fetchDocuments = () => {
        GetPendingDocuments()
            .then((result) => {
                if (!result || result.length === 0) {
                    setDocs([]);
                    setError("No Pending Documents");
                    return;
                } 
                setDocs(result);
                setError(null); // Clear error if documents are fetched successfully
            })
            .catch((err) => setError(err.message));
    };

    useEffect(() => {
        fetchDocuments();
        setupColumns();
    }, []);  // Runs only once on mount

    return (
        <Box m="20px" sx={{ width: 'dynamic', maxWidth: '95%', justifyContent: 'center' }}>
            <Header title="Pending Documents" />
            {error && (
                <Typography color="error" align="center" style={{ marginBottom: '16px' }}>
                    {error}
                </Typography>
            )}
            {message && (
                <Typography color="success" align="center" style={{ marginBottom: '16px' }}>
                    {message}
                </Typography>
            )}
            {docs.length > 0 && (
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
                        getRowId={(row) => row.ID}
                        sx={{ width: "dynamic", maxWidth: "170vh" }}
                    />
                </Box>
            )}
        </Box>
    );
};

export default PendingDocuments;
