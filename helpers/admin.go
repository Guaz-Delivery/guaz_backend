package helpers

import (
	"encoding/json"
	"github.com/Guaz-Delivery/guaz_backend/models"
	"net/http"
)

// Helper function to respond with error message
func AdminResponseWithError(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(models.Admin_Output{
		Token:    "",
		Admin_id: "",
		Error:    true,
		Message:  message,
	})
	return
}
