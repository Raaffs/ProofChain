import { Box, Button, Snackbar, IconButton, TextField } from "@mui/material";
import { Formik } from "formik";
import * as yup from "yup";
import Header from "../../components/Header";
import { ClassNames } from "@emotion/react";
import { useState } from "react";
import { UploadDocument } from "../../../wailsjs/go/main/App";
import FileUploadOutlinedIcon from "@mui/icons-material/FileUploadOutlined";
import PopUp from "../../components/PopUp";
const UploadDocs = () => {
  const handleFormSubmit = (values) => {
    console.log(values);
  };
  const [error, setError] = useState(null);
  const [message, setMessage] = useState("");

  const uploadFile = (values) => {
    UploadDocument(values.institute, values.documentName, values.description)
      .then(() => {
        setMessage("Document uploaded successfully");
      })
      .catch((err) => {
        setError(err);
      });
  };
  return (
    <Box m="20px">
      <Header
        title="Verify Document"
        subtitle="upload document you want to verify"
      />
      {error && (
        <PopUp
          Error={error}
          Message=""
          onClose={() => {
            setError(null);
          }}
        />
      )}
      {message && (
        <PopUp
          Message={message}
          Error={null}
          onClose={() => {
            setError(null);
          }}
        />
      )}
      <Formik
        onSubmit={uploadFile}
        initialValues={initialValues}
        validationSchema={checkoutSchema}
      >
        {({
          values,
          errors,
          touched,
          handleBlur,
          handleChange,
          handleSubmit,
        }) => (
          <form onSubmit={handleSubmit}>
            <Box display="flex" gap="30px" justifyContent="center">
              <TextField
                type="text"
                label="Document Name"
                onBlur={handleBlur}
                onChange={handleChange}
                value={values.documentName}
                name="documentName"
                error={!!touched.documentName && !!errors.documentName}
                helperText={touched.documentName && errors.documentName}
                sx={{ margin: "16px", width: "95%" }}
              />
              <TextField
                type="text"
                label="Institute Name"
                onBlur={handleBlur}
                onChange={handleChange}
                value={values.institute}
                name="institute"
                error={!!touched.institute && !!errors.institute}
                helperText={touched.institute && errors.institute}
                sx={{ margin: "16px", width: "95%" }}
              />
              <TextField
                type="text"
                label="Description"
                onBlur={handleBlur}
                onChange={handleChange}
                value={values.description}
                name="description"
                error={!!touched.description && !!errors.description}
                helperText={touched.description && errors.description}
                sx={{ margin: "16px", width: "95%" }}
              />
            </Box>
            <Box display="flex" justifyContent="center" mt="20px">
              <IconButton type="submit" color="secondary" variant="contained">
                <FileUploadOutlinedIcon sx={{ fontSize: 30 }} />
                Upload
              </IconButton>
            </Box>
          </form>
        )}
      </Formik>
    </Box>
  );
};

const checkoutSchema = yup.object().shape({
  documentName: yup
    .string()
    .min(3, "document name needs to be at least 3 characters long")
    .max(15, "document name cannot be more than 15 characters long")
    .required("required"),
  institute: yup
    .string()
    .max(15, "institute name cannot be more than 15 characters long")
    .min(2, "institute name must be atleast two characters long")
    .required("required"),
  description: yup
    .string()
    .max(15, "institute name cannot be more than 15 characters long")
    .min(2, "institute name must be atleast two characters long")
    .required("required"),
});
const initialValues = {
  documentName: "",
  institute: "",
  description: "",
};

export default UploadDocs;
