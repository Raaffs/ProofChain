import { Box, Typography, useTheme, Button, Fade } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import Header from "../../components/Header";
import { btnstyle } from "../../styles/styles";
import {
  GetPendingDocuments,
  IsApprovedInstitute,
  CreateDigitalCopy,
  ViewDocument,
} from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import Modal from "@mui/material/Modal";
import IssueCard from "../../components/cards/certificate";
import RemoveRedEyeSharpIcon from "@mui/icons-material/RemoveRedEyeSharp";
import { DataGridSx, DataGridDarkSx } from "../../styles/styles";
let isInstitute = false;
const PendingDocuments = () => {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const [docs, setDocs] = useState([]);
  const [error, setError] = useState(null);
  const [selectedDoc, setSelectedDoc] = useState(null);
  const [message, setMessage] = useState("");
  const [columns, setColumns] = useState([
    { field: "Requester", headerName: "Requester", flex: 1 },
    { field: "Verifier", headerName: "Verifier", flex: 1 },
    { field: "ShaHash", headerName: "Hash", flex: 1 },
    { field: "Institute", headerName: "Institute", flex: 1 },
    {
      field: "view",
      headerName: "View",
      flex: 0.5,
      justifyContent: "center",
      renderCell: (params) => {
        const doc = docs.find((doc) => doc.ID === params.id);
        return (
          <Box>
            <Button
              // Use 'outlined' variant for a cleaner, professional look
              // variant="outlined"
              // Use 'inherit' color to let the icon define the main color visually
              color="info"
              onClick={() => {
                setSelectedDoc(params.row); // save the row
                handleView(
                  params.row.ShaHash,
                  params.row.Institute,
                  params.row.Requester
                );
              }}
              sx={{
                // Ensures the button is compact for an icon-only use case
                minWidth: "auto",
                padding: "6px 8px",
                margin: 4,
                // Subtle background for better visibility against white
                // backgroundColor: 'rgba(0, 0, 0, 0.03)',
                // Use hardcoded color for hover border
                "&:hover": {
                  backgroundColor: "rgba(0, 0, 0, 0.1)",
                  // borderColor: '#64B5F6', // A light blue border on hover
                },
              }}
            >
              <RemoveRedEyeSharpIcon
                sx={{
                  // Professional Blue Color (e.g., MUI's standard info.main: #2196F3)
                  color: `linear-gradient(45deg, #00C6FF 30%, #0072FF 90%)`,
                  // fontSize: 20, // Adjusted size for visual balance
                }}
              />
            </Button>
          </Box>
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

  const handleApprove = (doc, data) => {
    console.log("document id: ", doc.ID);
    if (doc === undefined || doc === null) {
      setError("Document not found");
      console.log("Document id not found, id: ", id);
      return;
    }

    CreateDigitalCopy(0, doc.ShaHash, data)
      .then(() => {
        setMessage("Document approved successfully");
        fetchDocuments(); // Refresh documents after approval
      })
      .catch((err) => setError(err.message));
  };

  const handleReject = (doc) => {
    if (doc === undefined || doc === null) {
      setError("Document not found");
      console.log("Document id not found, id: ", id);
      return;
    }

    CreateDigitalCopy(1, doc.ShaHash)
      .then(() => {
        setMessage("Document rejected successfully");
        fetchDocuments(); // Refresh documents after rejection
      })
      .catch((err) => setError(err.message));
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
  }, []); // Runs only once on mount

  return (
    <Box
      m="20px"
      sx={{
        width: "dynamic",
        maxWidth: "95%",
        justifyContent: "center",
      }}
    >
      <Header title="Pending Documents" />
      {error && (
        <Typography
          color="error"
          align="center"
          style={{ marginBottom: "16px" }}
        >
          {error}
        </Typography>
      )}
      {message && (
        <Typography
          color="success"
          align="center"
          style={{ marginBottom: "16px" }}
        >
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
            sx={theme.palette.mode === "dark" ? DataGridDarkSx : DataGridSx}
          />
        </Box>
      )}
      <Modal
        open={isModalOpen}
        onClose={() => setModalOpen(false)}
        slotProps={{
          backdrop: {
            sx: {
              backgroundColor: "rgba(0,0,0,0.65)",
              backdropFilter: "blur(4px)",
            },
          },
        }}
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          p: 2,
        }}
      >
        <Box
          sx={{
            backgroundColor: "white",
            borderRadius: "18px",
            boxShadow:
              "0 12px 40px rgba(0,0,0,0.22), 0 0 0 1px rgba(0,0,0,0.06)",
            display: "flex",
            gap: 4,
            maxWidth: "1400px",
            width: "95vw",
            maxHeight: "92vh",
            overflow: "hidden",
            p: 4,
          }}
        >
          {/* Left Side: Image */}
          <Box
            sx={{
              flex: 1,
              borderRight: "1px solid #e5e5e5",
              pr: 4,
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              overflowY: "auto",
            }}
          >
            {image ? (
              <img
                src={`data:image/png;base64,${image}`}
                alt="Document Preview"
                style={{
                  width: "100%",
                  height: "auto",
                  maxHeight: "70vh",
                  objectFit: "contain",
                  borderRadius: "10px",
                }}
              />
            ) : (
              <Typography color="error" variant="h6" sx={{ py: 4 }}>
                Failed to load image
              </Typography>
            )}

            <Button
              variant="contained"
              onClick={() => setModalOpen(false)}
              sx={{
                mt: 3,
                borderRadius: "10px",
                px: 4,
                py: 1.5,
                textTransform: "none",
                fontSize: "16px",
              }}
            >
              Close Preview
            </Button>
          </Box>

          {/* Right Side: Issue Card */}
          <Box
            sx={{
              flex: 1,
              pl: 2,
              overflowY: "auto",
            }}
          >
            <IssueCard
              data={null}
              viewTitle="viewTitleForCard"
              onIssue={(certData) => {
                handleApprove(selectedDoc, certData);
              }}
            />
            <Box sx={{ mt: 2 }}>
              {" "}
              {/* Added margin-top for separation */}
              <Button
                variant="contained"
                color="error"
                onClick={() => handleReject(selectedDoc)}
                fullWidth // This MUI prop makes the button take the full width of its parent
                sx={{
                  borderRadius: "10px", // Optional: to match the style of the other button
                  px: 4,
                  py: 1.5,
                  textTransform: "none",
                  fontSize: "16px",
                }}
              >
                Reject
              </Button>
            </Box>
          </Box>
        </Box>
      </Modal>
    </Box>
  );
};

export default PendingDocuments;
