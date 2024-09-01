import React, { useContext, useEffect, useState } from "react";
import { IsLoggedIn } from "../../../wailsjs/go/main/App";
import { Logout } from "../../../wailsjs/go/main/App";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
// Component to render content only if the user is signed out
const Redirect=({to})=>{
  return(
      <Link to={to}/>
  )
}
const LogoutUser = ({setAuthStatus}) => {
  const navigate=useNavigate()
  Logout()
  .then(()=>{
    setAuthStatus(false)
    navigate('/')
  })
  .catch((error)=>{
    console.log(error)
  })
  return(
    <Redirect
      to='/'
    />
  )
}
export default LogoutUser