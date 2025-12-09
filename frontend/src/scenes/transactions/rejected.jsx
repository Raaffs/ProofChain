import { Box, Typography, useTheme } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import { tokens } from "../../themes";
import Header from "../../components/Header";
import { GetRejectedDocuments } from "../../../wailsjs/go/main/App";
import { useEffect, useState } from "react";

const RejectedDocuments = () => {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const [docs, setDocs] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getDocuments = () => {
      GetRejectedDocuments()
        .then((result) => {
          if (!result || result.length === 0) {
            setDocs([
              {
                ID: "",
                Requester: "",
                Verifier: "",
                Name: "",
                ShaHash: "",
              },
            ]);
            setError("No Rejected Documents");
          } else {
            const updatedDocs = result.map((doc) => {
              return {
                ...doc,
                ShaHash: doc.ShaHash,
              };
            });
            setDocs(updatedDocs);
          }
        })
        .catch((err) => {
          setError(err.message);
        });
    };
    getDocuments();
  }, []);

  const columns = [
    { field: "Requester", headerName: "Requester", flex: 1 },
    { field: "Verifier", headerName: "Verifier", flex: 1 },
    { field: "Name", headerName: "Name", flex: 1 },
    { field: "ShaHash", headerName: "Hash", flex: 1 },
  ];

  return (
    <Box
      m="20px"
      sx={{ width: "dynamic", maxWidth: "95%", justifyContent: "center" }}
    >
      <Header title="Rejected Documents" />
      {error && (
        <Typography
          color="error"
          align="center"
          style={{ marginBottom: "16px" }}
        >
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
            fontSize: "1.1rem",
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
          sx={theme.palette.mode === "dark" ? DataGridDarkSx : DataGridSx}
        />
      </Box>
    </Box>
  );
};

export default RejectedDocuments;
