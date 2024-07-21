import { Box, Button, TextField } from "@mui/material";
import { Formik } from "formik";
import * as yup from "yup";
import Header from "../../components/Header";

const UploadDocument = () => {

  const handleFormSubmit = (values) => {
    console.log(values);
  };

  return (
        <Box m="20px">
        <Header title="CREATE USER" subtitle="Create a New User Profile" />

        <Formik
            onSubmit={handleFormSubmit}
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
            <Box
              display="grid"
              gap="30px"
              gridTemplateColumns="repeat(4, minmax(0, 1fr))"
            >
              <TextField
                fullWidth
                variant="filled"
                type="text"
                label="Document Name"
                onBlur={handleBlur}
                onChange={handleChange}
                value={values.documentName}
                name="documentName"
                error={!!touched.documentName && !!errors.documentName}
                helperText={touched.documentName && errors.documentName}
                sx={{ gridColumn: "span 2" }}
              />
              <TextField
                fullWidth
                variant="filled"
                type="text"
                label="Institute Name"
                onBlur={handleBlur}
                onChange={handleChange}
                value={values.institute}
                name="institute"
                error={!!touched.institute && !!errors.institute}
                helperText={touched.institute && errors.institute}
                sx={{ gridColumn: "span 2" }}
              />  
              <TextField
              fullWidth
              variant="filled"
              type="text"
              label="Description"
              onBlur={handleBlur}
              onChange={handleChange}
              value={values.description}
              name="description"
              error={!!touched.description && !!errors.description}
              helperText={touched.description && errors.description}
              sx={{ gridColumn: "span 2" }}
            />
            </Box>
            <Box display="flex" justifyContent="end" mt="20px">
              <Button type="submit" color="secondary" variant="contained">
                Create New User
              </Button>
            </Box>
          </form>
        )}
      </Formik>
    </Box>
  );
};


const checkoutSchema = yup.object().shape({
  documentName  : yup.string()
                  .min(3,"document name needs to be at least 3 characters long")
                  .max(15,"document name cannot be more than 15 characters long")
                  .required("required"),
  institute: yup.string()
            .max(15,"institute name cannot be more than 15 characters long")
            .min(2,"institute name must be atleast two characters long")
            .required("required"),
  description: yup.string()
              .max(15,"institute name cannot be more than 15 characters long")
              .min(2,"institute name must be atleast two characters long")
              .required("required")
});
const initialValues = {
  documentName: "",
  institute: "",
  description: "",
};

export default UploadDocument;