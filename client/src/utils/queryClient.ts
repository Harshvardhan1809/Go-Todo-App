import { QueryClient } from "react-query";

export const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            staleTime: 1000 * 60 * 60, 
            cacheTime: 1000 * 60 * 60 * 2,
            refetchInterval: 1000 * 60 * 60,
            refetchOnReconnect: true,
            retry: 0,
        },
    },
});
