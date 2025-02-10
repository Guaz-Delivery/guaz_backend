package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/Guaz-Delivery/guaz_backend/models"
	"github.com/Guaz-Delivery/guaz_backend/queries"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
		return models.Login_Output{
			Token:      "",
			Error:      true,
			Message:    "unable to parse the response",
			Courier_id: "",
		}, errors.New("unable to parse the response")
	}

	checkRes := models.Response{}
	err = json.Unmarshal(resByte, &checkRes)

	if err != nil {
		return models.Login_Output{
			Token:      "",
			Error:      true,
			Message:    err.Error(),
			Courier_id: "",
		}, errors.New("unable to parse the response bytes")
	}

	// compare the password from stored hash
	err = bcrypt.CompareHashAndPassword([]byte(checkRes.Data.Couriers[0].Password), []byte(args.Args.Password))
	if err != nil {
		return models.Login_Output{
			Token:      "",
			Error:      true,
			Message:    "account doesn't exist or Password is not correct",
			Courier_id: "",
		}, errors.New("account doesn't exist or Password is not correct")
	}

	// generate token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "guaz-webhooks",
			"sub": checkRes.Data.Couriers[0].Id,
			"https://hasura.io/jwt/claims": map[string]interface{}{
				"x-hasura-default-role":  "courier",
				"x-hasura-allowed-roles": []string{"courier"},
				"x-hasura-user-id":       checkRes.Data.Couriers[0].Id,
			},
		})
	s, err := t.SignedString([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {

		return models.Login_Output{
			Token:      "",
			Error:      true,
			Message:    err.Error(),
			Courier_id: "",
		}, err
	}

	// return with the token
	message := "succssful"
	response = models.Login_Output{
		Token:      s,
		Courier_id: checkRes.Data.Couriers[0].Id,
		Error:      false,
		Message:    message,
	}

	return response, nil
}
