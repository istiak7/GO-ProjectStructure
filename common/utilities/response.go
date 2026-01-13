package utilities
import(
	"encoding/json"
	"net/http"
)
type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{}      `json:"error,omitempty"`
}


func SendResponse(w http.ResponseWriter, status int, success bool, message string, data interface{}, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiResponse{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	})
}