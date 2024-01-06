import { useMutation } from "@tanstack/react-query";
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
        mutationKey: ["login"],
        mutationFn: async (data: LoginFormDataType)  => { 
            const { username, password } = data;
            const body = JSON.stringify({username, password});
            const login = await axios.post("http://localhost:9010/auth/login", body, config);
            return login;
        },  
        onSuccess: () => {
            navigate("/");
        }, 
        onError: (e: any) => {
            console.log("Print error from login mutation", e);
        }
      })
      
}

export default useLoginMutation;