package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Guaz-Delivery/guaz_backend/models"
)

func HandleVerifyPayment(w http.ResponseWriter, r *http.Request) {
	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")

	// read request body
	var verifyPayload models.VerifyPaymentPayload
	if err := json.NewDecoder(r.Body).Decode(&verifyPayload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	log.Printf("%v", verifyPayload)

}

func verifyPayment(method string, url string, response interface{}) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CHAPA_SECRET_KEY")))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return json.Unmarshal(body, response)
}
