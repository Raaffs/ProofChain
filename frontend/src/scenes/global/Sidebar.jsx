import 'react-pro-sidebar/dist/css/styles.css'
import { useState } from "react"
import {ProSidebar, Menu, MenuItem} from "react-pro-sidebar"
import {Box,Button,IconButton,Typography,useTheme} from "@mui/material"
import { Link } from 'react-router-dom'
import { tokens } from '../../themes'
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import MenuOutlinedIcon from "@mui/icons-material/MenuOutlined";
import UploadFileOutlinedIcon from '@mui/icons-material/UploadFileOutlined';
import CancelOutlinedIcon from '@mui/icons-material/CancelOutlined';
import AddTaskOutlinedIcon from '@mui/icons-material/AddTaskOutlined'
import TimerOutlinedIcon from '@mui/icons-material/TimerOutlined';
import ReceiptOutlinedIcon from '@mui/icons-material/ReceiptOutlined';
import LogoutOutlinedIcon from '@mui/icons-material/LogoutOutlined';
import LoginOutlinedIcon from '@mui/icons-material/LoginOutlined';
import PersonAddOutlinedIcon from '@mui/icons-material/PersonAddOutlined';
import sidebarlogo from '../../assets/images/sidebar-3.jpg'
import zIndex from '@mui/material/styles/zIndex'
const Item = ({ title, to, icon, selected, setSelected }) => {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  let color = colors.grey[100]; // Default color

  // Check title to set color accordingly
  if (title === "Approved") {
    color = "green";
  } else if (title === "Rejected") {
    color = "red"
  } else if (title === "Pending") {
    color = "blue"
  }
  return (
    <MenuItem
      active={selected === title}
      style={{
        color:color,
      }}
      onClick={() => setSelected(title)}
      icon={icon}
    >
      <Typography>{title}</Typography>
      <Link to={to} />
    </MenuItem>
  );
};

const Sidebar = () => {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [selected, setSelected] = useState("Dashboard");

  return (
    <Box
      sx={{
        position: "sticky",
        display: "flex",
        height: "100vh",
        top: 0,
        bottom: 0,
        zIndex: 10000,
        "& .ps-menu-root":{
          position:'fixed'
        },
        "& .pro-sidebar-inner": {
          background: `${theme.palette.mode==="dark" ? 'transparent' : 'transparent'} !important`,
        },
        "& .pro-icon-wrapper": {
          backgroundColor: "transparent !important",
        },
        "& .pro-inner-item": {
          padding: "5px 35px 5px 10px !important",
        },
        "& .pro-inner-item:hover": {
          background: `${colors.blueAccent[700]} !important`,
        },
        "& .pro-menu-item.active": {
          color: `${colors.blueAccent[500]} !important`,
        },
      }}
    >
      <ProSidebar collapsed={isCollapsed}
        // image={theme.palette.mode=="dark"?'https://user-images.githubusercontent.com/25878302/144499035-2911184c-76d3-4611-86e7-bc4e8ff84ff5.jpg':sidebarlogo}
      >
        <Menu iconShape="square">
          {/* LOGO AND MENU ICON */}
          <MenuItem
            onClick={() => setIsCollapsed(!isCollapsed)}
            icon={isCollapsed ? <MenuOutlinedIcon /> : undefined}
            style={{
              margin: "10px 0px 20px 0px",
              color: colors.grey[100],
            }}
          >
            {!isCollapsed && (
              <Box
                display="flex"
                justifyContent="space-between"
                alignItems="center"
                ml="15px"
              >
                <Typography variant="h3" color={colors.greenAccent[300]}>
                  ProofChain
                </Typography>
                <IconButton onClick={() => setIsCollapsed(!isCollapsed)}>
                  <MenuOutlinedIcon />
                </IconButton>
              </Box>
            )}
          </MenuItem>

          {!isCollapsed && (
            <Box mb="20px">
            </Box>
          )}

          <Box paddingLeft={isCollapsed ? undefined : "10%"} paddingRight={isCollapsed?undefined:"10%"}>
           
          <Typography
              variant="h4"
              color={colors.grey[300]}
              sx={{ m: "15px 0 10px 2px" }}
            >Dashboard</Typography>
            <Item
              title="Dashboard"
              to="/dashboard"
              icon={<HomeOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />

            <Typography
              variant="h4"
              color={colors.grey[300]}
              sx={{ m: "15px 0 10px 2px" }}
            >
              Documents
            </Typography>
    
            <Item
              title="Approved"
              to="/documents/approved"
              icon={<AddTaskOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
            <Item
              title="Rejected"
              to="/documents/rejected"
              icon={<CancelOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
              style={{color: colors.redAccent[200]}}
            />
             <Item
              title="Pending"
              to="/documents/pending"
              icon={<TimerOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
            <Typography
              variant="h4"
              color={colors.grey[300]}
              sx={{ m: "15px 0 5px 5px" }}
            >
              Transacts
            </Typography>
            <Item
              title="All"
              to="/transactions"
              icon={<ReceiptOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
             <Item
              title="Approved Transactions"
              to="/transactions/approved"
              icon={<AddTaskOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
             <Item
              title="Rejected Transactions"
              to="/transactions/rejected"
              icon={<CancelOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
             <Item
              title="Pending Trasactions"
              to="/transactions/pending"
              icon={<TimerOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
            <Typography
              variant="h4"
              color={colors.grey[300]}
              sx={{ m: "15px 0 5px 15px" }}
            >
                
            Verify
            </Typography>
            <Item
              title="Upload"
              to="/upload"
              icon={<UploadFileOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
          

            <Typography
              variant="h4"
              color={colors.grey[300]}
              sx={{ m: "15px 0 5px 10px" }}
            >
              Accounts
            </Typography>
            <Item
              title="Logout"
              to="/logout"
              icon={<LogoutOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
              <Item
              title="Login"
              to="/"
              icon={<LoginOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
            <Item
              title="New Account"
              to="/register"
              icon={<PersonAddOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
          </Box>
        </Menu>
      </ProSidebar>
    </Box>
  );
};

export default Sidebar;