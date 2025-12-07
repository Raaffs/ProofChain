import { Box,Typography } from "@mui/material";
import { useTheme } from "@emotion/react";
import { tokens } from "../../themes";
import StatBox from "../../components/StatBox";
import CancelOutlinedIcon from '@mui/icons-material/CancelOutlined';
import AddTaskOutlinedIcon from '@mui/icons-material/AddTaskOutlined'
import TimerOutlinedIcon from '@mui/icons-material/TimerOutlined';
import { useEffect, useState } from "react";
import { GetAllDocs } from "../../../wailsjs/go/main/App";
import { IsApprovedInstitute } from "../../../wailsjs/go/main/App";
import StatusBreakdownChart from "../../components/PieChart";
import InstituteStatusChart from "../../components/BarChart";
const Dashboard=()=>{
    const theme=useTheme()
    // const colors=tokens(theme.palette.mode)
    const colors=tokens(theme.palette.mode)
    const [docs, setdDocs]=useState([])
    const [approved,setApproved]=useState(0)
    const [pending,setPending]=useState(0)
    const [rejected,setRejected]=useState(0)
    const [loading, setLoading] = useState(true)
    const [isInstitute,setIsInstitute]=useState(false)
    const getDocuments=()=>{
        // const navigate=useNavigate()
        GetAllDocs()
        .then((result)=>{
            console.log(result)
            let approvedCount =result.filter(doc=>{return doc.Stats==0})
            let rejectedCount =result.filter(doc=>{return doc.Stats==1})
            let pendingCount =result.filter(doc=>{return  doc.Stats==2})
            setApproved(approvedCount.length)
            setPending(pendingCount.length)
            setRejected(rejectedCount.length)
            setdDocs(result)
            setLoading(false)
        })
        .catch((err)=>{
            console.log("error retrieving documents: ",err)
        })
    }

    useEffect(()=>{
        IsApprovedInstitute()
        .then((result)=>{
          setIsInstitute(result)
        })
        .catch(err=>{
          console.log(err)
        })
        getDocuments()
    },[])

    if (loading) {
        return (
            <Box display="flex" justifyContent="center" alignItems="center" height="100vh">
                loading....
            </Box>
        );
    }

    return(
        <Box m="20px">
           
            <Box
             display="grid"
             gridTemplateColumns="repeat(12, 1fr)"
             gridAutoRows="140px"
             gap="20px"
            >
                <Box
                  gridColumn="span 3"
                  backgroundColor={colors.primary[600]}
                  display="flex"
                  alignItems="center"
                  justifyContent="center"
                >
                  <StatBox
                    title={approved}
                    subtitle="Approved Documents"
                    icon={
                        <AddTaskOutlinedIcon
                            sx={{ color: colors.greenAccent[500], fontSize: "26px",margin:"5px" }}
                        />
                    }
                    path='/documents/approved'
                    fontColor={colors.greenAccent[500]}
                  />
                </Box>
                <Box

                  gridColumn="span 3"
                  backgroundColor={colors.primary[600]}
                  display="flex"
                  alignItems="center"
                  justifyContent="center"
                >
                  <StatBox
                    title={pending}
                    subtitle="Pending Documents"
                    fontColor={colors.blueAccent[500]}
                    icon={
                        <TimerOutlinedIcon
                            sx={{ color: colors.blueAccent[600], fontSize: "26px",margin:"5px" }}
                        />
                    }
                    path='/documents/pending'
                  />
                </Box>
                <Box
                  gridColumn="span 3"
                  backgroundColor={colors.primary[600]}
                  display="flex"
                  alignItems="center"
                  justifyContent="center"
                >
                  <StatBox
                    title={rejected}
                    subtitle="Rejected Documents"
                    fontColor={colors.redAccent[500]}
                    icon={
                        <CancelOutlinedIcon
                            sx={{ color: colors.redAccent[600], fontSize: "26px",margin:"5px" }}
                        />
                    }
                    path='/documents/rejected'
                  />
                </Box>
                <Box
                  gridColumn="span 3"
                  backgroundColor={colors.primary[600]}
                  display="flex"
                  alignItems="center"
                  justifyContent="center"
                >
                  <StatBox
                    title={approved+pending+rejected}
                    subtitle="Total"
                  fontColor={colors.greenAccent[500]}
                  />
                </Box>
                <Box
                  gridColumn="span 8"
                  gridRow="span 3"
                  backgroundColor={colors.primary[500]}
                  overflow="auto"
                >
                    <Box
                        display="flex"
                        justifyContent="space-between"
                        alignItems="center"
                        borderBottom={`4px solid ${colors.primary[500]}`}
                        colors={colors.grey[500]}
                        p="15px"
                    >
                      
                     <Typography color={colors.grey[100]} variant="h5" fontWeight={600} >Transactions History</Typography>
                    </Box>
                    {docs.map((doc, i) => (
                      <Box 
                          key={i}  
                          display="flex"
                          justifyContent="space-between"
                          alignItems="center"
                          borderBottom={`4px solid ${colors.primary[900]}`}
                          p="15px"
                      >
                        <Box>
                          <Typography
                            color={colors.greenAccent[500]}
                            variant="h5"
                            fontWeight="600"
                          >
                            {doc.ID}
                          </Typography>
                          <Typography color={colors.grey[100]}>
                            {doc.Name}
                          </Typography>
                        </Box>
                        <Box color={colors.grey[100]}>{isInstitute?doc.Requester:doc.Verifier}</Box>
                       <Box
                        p="5px 10px"
                        backgroundColor={
                          doc.Stats == '0'
                              ? colors.greenAccent[500]  // Approved
                              : doc.Stats == '1'
                              ? colors.redAccent[500]    // Rejected
                              : colors.blueAccent[600]   // Pending
                      }
                        borderRadius="4px"
                      >
                        {doc.Stats == '0' ? "Approved" : doc.Stats == '1' ? "Rejected" : "Pending"}
                      </Box>
                      </Box>
                      ))}
                      
                </Box>
                <Box
                  gridColumn="span 4"
                  gridRow="span 3"
                  backgroundColor={colors.primary[500]}
                  overflow="auto"
                >
                  <StatusBreakdownChart data={docs}/>

                </Box>
                <Box
                  gridColumn="span 12 "
                  gridRow="span 2"
                  backgroundColor={colors.primary[500]}
                  overflow="auto"
                >
                  <InstituteStatusChart data={docs}/>
                </Box>
            </Box>
        </Box>
    )
}

export default Dashboard