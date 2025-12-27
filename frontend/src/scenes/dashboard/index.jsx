import { Box, Typography } from "@mui/material";
import { useTheme } from "@emotion/react";
import { tokens } from "../../themes";
import StatBox from "../../components/StatBox";
import CancelOutlinedIcon from "@mui/icons-material/CancelOutlined";
import AddTaskOutlinedIcon from "@mui/icons-material/AddTaskOutlined";
import TimerOutlinedIcon from "@mui/icons-material/TimerOutlined";
import { useEffect, useState } from "react";
import { GetAllDocs } from "../../../wailsjs/go/main/App";
import { IsApprovedInstitute } from "../../../wailsjs/go/main/App";
import StatusBreakdownChart from "../../components/PieChart";
import InstituteStatusChart from "../../components/BarChart";
import { statBoxStyles, flexHeaderBoxStyles } from "../../styles/styles";

const Dashboard = () => {
  const theme = useTheme();
  // const colors=tokens(theme.palette.mode)
  const colors = tokens(theme.palette.mode);
  const [docs, setdDocs] = useState([]);
  const [approved, setApproved] = useState(0);
  const [pending, setPending] = useState(0);
  const [rejected, setRejected] = useState(0);
  const [loading, setLoading] = useState(true);
  const [isInstitute, setIsInstitute] = useState(false);
  const getDocuments = () => {
    // const navigate=useNavigate()
    GetAllDocs()
      .then((result) => {
        console.log(result);
        let approvedCount = result.filter((doc) => {
          return doc.Stats == 0;
        });
        let rejectedCount = result.filter((doc) => {
          return doc.Stats == 1;
        });
        let pendingCount = result.filter((doc) => {
          return doc.Stats == 2;
        });
        setApproved(approvedCount.length);
        setPending(pendingCount.length);
        setRejected(rejectedCount.length);
        setdDocs(result);
        setLoading(false);
      })
      .catch((err) => {
        console.log("error retrieving documents: ", err);
      });
  };

  useEffect(() => {
    IsApprovedInstitute()
      .then((result) => {
        setIsInstitute(result);
      })
      .catch((err) => {
        console.log(err);
      });
    getDocuments();
  }, []);

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        loading....
      </Box>
    );
  }

  return (
    <Box m="20px"
      backgroundColor='transparent'
    >
      <Box
        display="grid"
        gridTemplateColumns="repeat(12, 1fr)"
        gridAutoRows="140px"
        gap="20px"
      >
        <Box sx={statBoxStyles(theme, colors)}>
          <StatBox
            title={approved}
            subtitle="Approved Documents"
            icon={
              <AddTaskOutlinedIcon
                sx={{
                  color: colors.greenAccent[500],
                  fontSize: "26px",
                  margin: "5px",
                }}
              />
            }
            path="/documents/approved"
            fontColor={colors.greenAccent[500]}
          />
        </Box>
        <Box sx={statBoxStyles(theme, colors)}>
          <StatBox
            title={pending}
            subtitle="Pending Documents"
            fontColor={colors.blueAccent[500]}
            icon={
              <TimerOutlinedIcon
                sx={{
                  color: colors.blueAccent[600],
                  fontSize: "26px",
                  margin: "5px",
                }}
              />
            }
            path="/documents/pending"
          />
        </Box>
        <Box sx={statBoxStyles(theme, colors)}>
          <StatBox
            title={rejected}
            subtitle="Rejected Documents"
            fontColor={colors.redAccent[500]}
            icon={
              <CancelOutlinedIcon
                sx={{
                  color: colors.redAccent[600],
                  fontSize: "26px",
                  margin: "5px",
                }}
              />
            }
            path="/documents/rejected"
          />
        </Box>
        <Box sx={statBoxStyles(theme, colors)}>
          <StatBox
            title={approved + pending + rejected}
            subtitle="Total"
            fontColor={colors.greenAccent[500]}
          />
        </Box>
        <Box
          gridColumn="span 8"
          gridRow="span 3"
          backgroundColor={`${
            theme.palette.mode === "dark" ? colors.primary[800] :colors.blueAccent[900]
          }`}
          overflow="auto"
          borderRadius="12px"
          sx={{
            boxShadow:
              theme.palette.mode === "dark"
                ? "0 4px 15px rgba(0,0,0,0.6)"
                : "0 4px 15px rgba(0,0,0,0.08)",
          }}
        >
          <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            borderBottom={`4px solid ${
              theme.palette.mode === "dark"
                ? colors.primary[800]
                : "transparent"
            }`}
            p="15px 20px"
          >
            <Typography
              
              variant="h5"
              fontWeight={600}
            >
              Transactions History
            </Typography>
          </Box>

          {docs.map((doc, i) => (
            <Box
              key={i}
              display="flex"
              justifyContent="space-between"
              alignItems="center"
              borderBottom={`4px solid ${
                theme.palette.mode === "dark"
                  ? colors.primary[900]
                  : "transparent"
              }`}
              p="15px 20px"
              backgroundColor={`${
                theme.palette.mode === "dark" ? colors.primary[500] : colors.blueAccent[900]
              }`}
              sx={{
                transition: "all 0.25s",
                borderRadius: "8px",
                mb: "10px",
                boxShadow:
                  theme.palette.mode === "dark"
                    ? "none"
                    : "0 2px 6px rgba(0,0,0,0.03)",
                // "&:hover": {
                //   backgroundColor:
                //     theme.palette.mode === "dark"
                //       ? colors.primary[500]
                //       : "#f7f7f7",
                //   boxShadow:
                //     theme.palette.mode === "dark"
                //       ? "none"
                //       : "0 4px 12px rgba(0,0,0,0.08)",
                // },
              }}
            >
              <Box>
                <Typography
                  color={colors.greenAccent[500]}
                  variant="h5"
                  fontWeight="600"
                >
                  {doc.ID}
                </Typography>
              </Box>

              <Box
                color={
                  theme.palette.mode === "dark"
                    ? colors.grey[100]
                    : colors.grey[700]
                }
              >
                {isInstitute ? doc.Requester : doc.Verifier}
              </Box>

              <Box
                p="5px 10px"
                backgroundColor={
                  doc.Stats == "0"
                    ? colors.greenAccent[500] // Approved
                    : doc.Stats == "1"
                    ? colors.redAccent[500] // Rejected
                    : colors.blueAccent[600] // Pending
                }
                borderRadius="6px"
                sx={{
                  color: "#fff",
                  fontWeight: 600,
                  textAlign: "center",
                }}
              >
                {doc.Stats == "0"
                  ? "Approved"
                  : doc.Stats == "1"
                  ? "Rejected"
                  : "Pending"}
              </Box>
            </Box>
          ))}
        </Box>

        <Box
          gridColumn="span 4"
          gridRow="span 3"
          borderRadius= "12px"
          sx={{
            boxShadow:"0px 4px 12px rgba(0,0,0,0.1)"
          }}
          backgroundColor={`${
            theme.palette.mode === "dark" ? colors.primary[800] :colors.blueAccent[900]
          }`}
          overflow="auto"
        >
          <StatusBreakdownChart data={docs} />
        </Box>
        <Box
          gridColumn="span 12 "
          gridRow="span 2"
          borderRadius= "12px"
          backgroundColor={`${
            theme.palette.mode === "dark" ? colors.primary[800] :colors.blueAccent[900]
          }`}
          sx={{
            boxShadow:"0px 4px 12px rgba(0,0,0,0.2)"
          }}
          overflow="auto"
        >
          <InstituteStatusChart data={docs} />
        </Box>
      </Box>
    </Box>
  );
};

export default Dashboard;
