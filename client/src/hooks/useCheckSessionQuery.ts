import { useQuery } from "react-query";
import axios from "axios"
import { config } from "../utils/axios";

const useCheckSessionQuery = () => {
    return useQuery({
        queryKey: ["auth", "session"],
        queryFn: async () => {
            const session = await axios.get("http://localhost:9010/auth/session", config);
            console.log(session);
            return session;
        },
    }
    )
} 

export default useCheckSessionQuery;