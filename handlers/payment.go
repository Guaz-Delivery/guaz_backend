package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"

	"github.com/Guaz-Delivery/guaz_backend/helpers"
	"github.com/Guaz-Delivery/guaz_backend/models"
)

func HandlePayment(w http.ResponseWriter, r *http.Request) {

	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")

	// read request body
	var actionPayload models.PaymentActionPayload
	if err := json.NewDecoder(r.Body).Decode(&actionPayload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Send the request params to the Action's generated handler function
	result, err := creditPayment(actionPayload.Input)

	if err != nil {
		helpers.PaymentResponseWithError(w, err.Error())
		return
	}

	// Respond with JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
	}

	// throw if an error happens

}

func creditPayment(args models.CREDIT_PAYMENTArgs) (response interface{}, err error) {

	payload := models.Payment{
		Amount:       float64(args.Args.Amount),
		Currency:     "ETB",
		Email:        args.Args.Email,
		Phone_Number: args.Args.Phone_number,
		Tx_Ref:       args.Args.Delivery_id,
		Courier_id:   args.Args.Delivery_id,
		Callback_URL: os.Getenv("CALLBACK_URL"),
		Return_URL:   args.Args.Return_URL,
		Customization: map[string]interface{}{
			"title":       "Laptop",
			"description": "deliver this performant laptop to my son",
		}}
	var res models.ChapaResponse
	err = sendHttpRequest(http.MethodPost, os.Getenv("CHAPA_URL"), payload, &res)
	if err != nil {
		log.Println(err.Error())
	}
	response = models.Credit_Output{
		Checkout_url: res.Data.Checkout_URL,
		Status:       res.Status,
	}
	return response, nil
}

func randomString(length int) string {
	var chars []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.")
	var result []byte
	for range length {
		result = append(result, chars[int(math.Floor(rand.Float64()*float64(len(chars))))])
	}
	return string(result)
}

func sendHttpRequest(method string, url string, payload models.Payment, response interface{}) error {

	client := &http.Client{}
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(payloadByte))

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CHAPA_SECRET_KEY")))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
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
