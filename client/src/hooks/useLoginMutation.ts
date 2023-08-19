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
            const login = await axios.post("test/auth", body, config)
            return login;
        },  
        onSuccess: () => {
            navigate("");
        }, 
        onError: () => {
          navigate("");
        }
      })
      
}

export default useLoginMutation;