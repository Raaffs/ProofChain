import { Box, Typography, useTheme, Button } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import Header from "../../components/Header";
import { GetPendingDocuments, IsApprovedInstitute, ApproveDocument, ViewDocument } from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import Modal from '@mui/material/Modal';
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
        {field: "Institute", headerName: "Institute", flex: 1},
        {
            field: "view",
            headerName: "View",
            flex: 1,
            renderCell: (params) => {
                const doc = docs.find((doc) => doc.ID === params.id);
                return (
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={() => {
                            handleView(params.row.ShaHash,params.row.Institute,params.row.Requester)
                            console.log("document in handle view: ",params.row.ShaHash)
                        }}
                    >
                        View
                    </Button>
                );
            },
        },
    ]);

    const [image, setImage] = useState(null); // To store the base64 image
    const [isModalOpen, setModalOpen] = useState(false); // Modal state

    const handleView = (shaHash, institute, requester) => {
        ViewDocument(shaHash, institute, requester)
            .then((base64Image) => {
                if (base64Image) {
                    setImage(base64Image); // Set the image received from the backend
                    setModalOpen(true); // Open the modal to display the image
                } else {
                    setError("No image received from the server");
                }
            })
            .catch((err) => setError(err.message));
    };

    const handleApprove = (doc) => {
        console.log("document id: ",doc.ID)
        if (doc===undefined||doc===null) {
            setError("Document not found");
            console.log("Document id not found, id: ", id);
            return;
        }

        ApproveDocument(0, doc.ShaHash)
            .then(() => {
                setMessage("Document approved successfully");
                fetchDocuments(); // Refresh documents after approval
            })
            .catch((err) => setError(err.message));
    };

    const handleReject = (doc) => {
        if (doc===undefined||doc===null) {
            setError("Document not found");
            console.log("Document id not found, id: ", id);
            return;
        }

        ApproveDocument(1, doc.ShaHash)
            .then(() => {
                setMessage("Document rejected successfully");
                fetchDocuments(); // Refresh documents after rejection
            })
            .catch((err) => setError(err.message));
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
                                        onClick={() => handleApprove(params.row)}
                                        style={{ marginRight: "10px" }}
                                    >
                                        Approve
                                    </Button>
                                    <Button
                                        variant="contained"
                                        color="error"
                                        onClick={() => handleReject(params.row)}
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
    }, []); // Runs only once on mount

    return (
        <Box m="20px" sx={{ width: "dynamic", maxWidth: "95%", justifyContent: "center" }}>
            <Header title="Pending Documents" />
            {error && (
                <Typography color="error" align="center" style={{ marginBottom: "16px" }}>
                    {error}
                </Typography>
            )}
            {message && (
                <Typography color="success" align="center" style={{ marginBottom: "16px" }}>
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
                            fontSize: "1.1rem",
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

            {/* Modal for image display */}
            <Modal
                open={isModalOpen}
                onClose={() => setModalOpen(false)}
                sx={{ display: "flex", alignItems: "center", justifyContent: "center" }}
            >
                <Box
                    sx={{
                        backgroundColor: "white",
                        padding: "20px",
                        borderRadius: "10px",
                        boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.25)",
                        textAlign: "center",
                    }}
                >
                    {image ? (
                        <img
                            src={`data:image/png;base64,${image}`}
                            alt="Document"
                            style={{ maxWidth: "100%", maxHeight: "70vh" }}
                        />
                    ) : (
                        <Typography color="error">Failed to load image</Typography>
                    )}
                    <Button
                        variant="contained"
                        color="secondary"
                        onClick={() => setModalOpen(false)}
                        style={{ marginTop: "10px" }}
                    >
                        Close
                    </Button>
                </Box>
            </Modal>
        </Box>
    );
};

export default PendingDocuments;
