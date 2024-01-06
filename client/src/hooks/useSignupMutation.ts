import axios from "axios";
import { useMutation } from "@tanstack/react-query";
import { NavigateFunction } from "react-router-dom";
import { config } from "../utils/axios";

export interface SignupFormDataType {
    username: string,
    password: string,
}

interface Options {
    navigate: NavigateFunction;
}

const useSignupMutation = ({navigate}: Options) => {
    return useMutation({
        mutationKey: ["signup"],
        mutationFn: async (data: SignupFormDataType) => {
            const { username, password } = data;
            const body = JSON.stringify({username, password});
            const signup = await axios.post("http://localhost:9010/auth/signup", body, config)
            return signup;
        },  
        onSuccess: () => {
            navigate("/login");
        }, 
        onError: () => {
            navigate("/signup");
        }
    })
}

export default useSignupMutation;