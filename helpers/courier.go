package helpers

import (
	"encoding/json"
	"github.com/Guaz-Delivery/guaz_backend/models"
	"net/http"
)

// Helper function to respond with error message
func CourierResponseWithError(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(models.Courier_Output{
		Token:      "",
		Courier_id: "",
		Error:      true,
		Message:    message,
	})
	return
}
