package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"os"

	"github.com/Guaz-Delivery/guaz_backend/models"
	"github.com/Guaz-Delivery/guaz_backend/queries"
)

func HandleCourierSignup(w http.ResponseWriter, r *http.Request) {

	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")

	// read request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "invalid body")
	}

	// parse the body as action payload
	var actionPayload models.SignupCourierActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "invalid payload")
	}

	// Send the request params to the Action's generated handler function
	result, err := SIGNUP_COURIER(actionPayload.Input, r.Header.Get("x-hasura-admin-secret"))

	// throw if an error happens
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	}

	// Write the response as JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "unable to send processed data", http.StatusInternalServerError)
		return
	}

}

func SIGNUP_COURIER(args models.SIGNUP_COURIERArgs, secret string) (response models.Signup_Output, err error) {
	fmt.Print("args", args.Args)
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Args.Password), 10)
	if err != nil {
		return models.Signup_Output{}, err
	}

	//  prepare the request
	variables := map[string]interface{}{
		"email":        args.Args.Email,
		"first_name":   args.Args.First_name,
		"middle_name":  args.Args.Middle_name,
		"last_name":    args.Args.Last_name,
		"location":     args.Args.Location,
		"phone_number": args.Args.Phone_number,
		"rate":         args.Args.Rate,
		"password":     hashedPassword,
	}
	reqBody := models.GraphQLRequest{
		Query:     queries.SIGNUP_COURIERS,
		Variables: variables,
	}
	reqBytes, err := json.Marshal(reqBody)

	// add the new user to the table by using hashed password
	bodyReader := bytes.NewReader(reqBytes)
	req, err := http.NewRequest(http.MethodPost, os.Getenv("GRAPHQL_URL"), bodyReader)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("x-hasura-admin-secret", secret)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return models.Signup_Output{}, err
	}
	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Signup_Output{}, err
	}
	regRes := models.RegisterResponse{}

	err = json.Unmarshal(resByte, &regRes)

	if err != nil {
		return models.Signup_Output{}, err
	}

	fmt.Println(regRes)

	response = models.Signup_Output{
		Token:      "<sample value>",
		Courier_id: "<sample value>",
		Error:      false,
	}
	return response, nil
}

func responseWithError(w http.ResponseWriter, status int, message string) {
	errorResponse := models.Signup_Output{
		Token:      "",
		Courier_id: "",
		Error:      true,
		Message:    &message,
	}
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		http.Error(w, "invalid payload", http.StatusInternalServerError)
	}
	return
}
