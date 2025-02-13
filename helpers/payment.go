package helpers

import (
	"encoding/json"
	"github.com/Guaz-Delivery/guaz_backend/models"
	"net/http"
)

// Helper function to respond with error message
func PaymentResponseWithError(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(models.Credit_Output{
		Checkout_url: "",
		Status:       "failed",
	})
	return
}
