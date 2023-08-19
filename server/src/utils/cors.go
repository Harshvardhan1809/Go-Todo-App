package utils

import (
	"net/http"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json") 
	(*w).Header().Set("Access-Control-Max-Age", "2160000")
	// (*w).Header().Set("Access-Control-Allow-Headers", "*")
    // (*w).Header().Set("Access-Control-Allow-Origin", "*")
    // (*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
