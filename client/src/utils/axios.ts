export const config = {
    headers : {
        "Content-Type" : "application/json",
        "Access-Control-Allow-Origin": "*",
    },
    withCredentials : true, 
}

// To debug the CORS error, not much is needed from the client side
// The server and the browser should allow the request
