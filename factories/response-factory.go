package factories

import (
	"encoding/json"
	"net/http"

	myTypes "github.com/nathangds/order-api/types"
)

func ResponseFactory(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(myTypes.CustomResponse{
		StatusCode: statusCode,
		Data:       data,
	})
}

func ErrorResponse(errorMessages []string) myTypes.ErrorResponse {
	return myTypes.ErrorResponse{
		Messages: errorMessages,
	}
}

func SuccessResponse(statusCode int, data interface{}) myTypes.CustomResponse {
	return myTypes.CustomResponse{
		StatusCode: statusCode,
		Data:       data,
	}
}
