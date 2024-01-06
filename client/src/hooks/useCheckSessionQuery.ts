import { useQuery } from "@tanstack/react-query";
import axios from "axios"
import { config } from "../utils/axios";

const useCheckSessionQuery = () => {
    return useQuery({
        queryKey: ["auth", "session"],
        queryFn: async () => {
            const response = await axios.get("http://localhost:9010/auth/check_auth", config);
            return response.data;
        },
        onSuccess: (data) => {
            return data
        },
    }
    )
} 

export default useCheckSessionQuery;