import { useMutation } from "react-query";
import axios from "axios";
import { config } from "../utils/axios";
import { NavigateFunction } from "react-router-dom";

export interface LoginFormDataType {
    username: string,
    password: string,
}

interface Options {
    navigate: NavigateFunction;
}

const useLoginMutation = ({navigate}: Options) => {

    return useMutation({
        mutationKey: "login",
        mutationFn: async (data: LoginFormDataType)  => { 
            const { username, password } = data;
            const body = JSON.stringify({username, password});
            console.log("body of login mutation ", username, password)
            const login = await axios.post("http://localhost:9010/auth/login", body, config)
            console.log("login print ", login)
            return login;
        },  
        onSuccess: () => {
            console.log("In onSuccess")
            navigate("/");
        }, 
        onError: () => {
            console.log("Getting error for login mutation")
            navigate("");
        }
      })
      
}

export default useLoginMutation;