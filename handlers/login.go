package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Guaz-Delivery/guaz_backend/models"
	"github.com/Guaz-Delivery/guaz_backend/queries"
)

func HandleCourierLogin(w http.ResponseWriter, r *http.Request) {

	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")

	// read request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// parse the body as action payload
	var actionPayload models.LoginCourierActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Send the request params to the Action's generated handler function
	result, err := LOGIN_COURIER(actionPayload.Input, r.Header.Get("x-hasura-admin-secret"))

	// throw if an error happens
	if err != nil {
		errorBody, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	// Write the response as JSON
	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Auto-generated function that takes the Action parameters and must return it's response type
func LOGIN_COURIER(args models.LOGIN_COURIERArgs, secret string) (response models.Login_Output, err error) {
	// fectch the user data
	variables := map[string]interface{}{
		"email":        args.Args.Email,
		"phone_number": args.Args.Phone_Number,
	}
	reqBody := models.GraphQLRequest{
		Query:     queries.LOGIN_COURIER,
		Variables: variables,
	}
	reqBytes, err := json.Marshal(reqBody)

	bodyReader := bytes.NewReader(reqBytes)

	req, err := http.NewRequest(http.MethodPost, os.Getenv("GRAPHQL_URL"), bodyReader)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("x-hasura-admin-secret", secret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Login_Output{}, err
	}
	resByte, err := io.ReadAll(res.Body)

	if err != nil {
		return models.Login_Output{}, errors.New("unable to parse the response")
	}

	checkRes := models.Response{}
	err = json.Unmarshal(resByte, &checkRes)

	if err != nil {
		return models.Login_Output{}, errors.New("unable to parse the response bytes")
	}

	log.Printf("resbytes %s ", checkRes.Data.Couriers[0])
	// compare the password from stored hash

	// generate token

	// return with the token
	message := "succssful"
	response = models.Login_Output{
		Token:      "<sample value>",
		Courier_id: "<sample value>",
		Error:      false,
		Message:    &message,
	}
	return response, nil
}
