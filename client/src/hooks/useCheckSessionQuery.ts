import { useQuery } from "react-query";
import axios from "axios"
import { config } from "../utils/axios";

const useCheckSessionQuery = () => {
    return useQuery({
        queryKey: ["auth", "session"],
        queryFn: async () => {
            const session = await axios.get("http://localhost:9010/auth/session", config);
            console.log("In check session query")
            console.log(session);
            return session;
        },
        onSuccess: (data) => {
            console.log("Successfully checked sessions", data)
        },
        onError: (error) => {
            console.log("Error checking sessions", error)
        } 
    }
    )
} 

export default useCheckSessionQuery;