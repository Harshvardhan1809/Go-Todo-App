import React, {Fragment, useState, useEffect} from 'react'
import { Typography, Card, CardHeader, CardContent, CardActions, Box, TextField, Button, CssBaseline } from '@mui/material';
import {useTheme} from '@mui/material';
import axios from "axios";
import { Navigate } from "react-router-dom";
import { useQuery, useMutation } from "react-query";
import { queryClient } from '../index';

export type loginFormDataType = {
    username: string,
    password: string,
}
  
export type LoginProps = {
    setUser: React.Dispatch<React.SetStateAction<null>>, 
}

// Define a mutation so that it can be stored anywhere
// queryClient.setMutationDefaults('login', {
//   mutationFn: login,
//   onMutate: (data: loginFormDataType) => { 
//     const config = {
//       headers: {
//         "Content-Type": "application/json"
//       }
//     } 
//     const {username, password} = data
//     const body = JSON.stringify({username, password});
//     // returned data gets stored in context
//     return axios.post("test/auth", body, config)
//   }, 
//   onSuccess: (result, variables, context) => {
//     console.log("Logged in successfully",result.data[0])
//     return (<Navigate to=""/>) 
//   }, 
//   onError: (result, variables, context) => {
//     return (<Navigate to=""/>) 
//   }

// })

const Login = () => {

    const theme = useTheme();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const onChangeUsername = (e: React.ChangeEvent<HTMLTextAreaElement>) =>{
      console.log(e.target.value);
      setUsername(e.target.value);
    }
    const onChangePassword = (e: React.ChangeEvent<HTMLTextAreaElement>) =>{
      console.log(e.target.value);
      setPassword(e.target.value);
    }  

    // Mutation to do login
    const loginMutation = useMutation({
        mutationFn: (data: loginFormDataType) => { 
          const config = {
            headers: {
              "Content-Type": "application/json"
            }
          }
          const body = JSON.stringify({username, password});
          return axios.post("test/auth", body, config)
        },  
        onSuccess: (result) => {
          console.log("Logged in successfully",result.data[0])
          return (<Navigate to=""/>) 
        }, 
        onError: (result) => {
          console.log("Error logging in", result)
          return (<Navigate to=""/>) 
        }
      })
    
      // if(loginMutation.isSuccess){
      //   // queryClient.setQueryData(["login"], ()) 
    
      //   // redirect to other component
      //   return (<h1>Logged in successfully</h1>)
      // }
      // else if(loginMutation.isLoading){
      //   // show loading spinner
      //   return (<h1>Loading ...</h1>)
      // }
      // else if(loginMutation.isError){
      //   console.log("Error logging in")
      //   return (<h1>Error logging in</h1>)
      // }
    

    // here we make a query to the /auth/validate to check if we are logged in

    return (
        <Box sx={{maxWidth: "100%", width: "100%" }}>
          <Box sx={{paddingTop: "10px", textAlign:"center"}}>
            <Typography variant="h4">To-do App Login</Typography>
          </Box>
          <Card sx={{ width: {xs:"85%", sm:"60%", md:"30%", lg:"30%", xl:"20%"}, margin: "auto", marginTop: "100px" }}>
            <Box sx={{display: "flex", justifyContent: "center", backgroundColor: { xs: "red", sm: "orange", md: "green", lg: "blue", xl: "pink" }, }}>
              <CardHeader title="Login Page"/>
            </Box> 
              <CardContent>
                <form action="">
                  <Box sx={{display:"flex", flexDirection:"column"}}>
                    <TextField label="Username" variant="outlined" sx={{paddingBottom: "10px"}} name="username" onChange={onChangeUsername} value={username}/>
                    <TextField label="Password" variant="outlined" sx={{paddingBottom: "10px"}} name="password" onChange={onChangePassword} value={password}/>
                    <Button variant="outlined" sx={{margin: "auto", marginBottom: "10px", marginTop: "10px"}} 
                    onClick={() => { loginMutation.mutate({ username: username, password: password }) } }>Login</Button> 
                  </Box>  
                </form> 
              </CardContent>
          </Card>
        </Box>
    )
}

export default Login;
