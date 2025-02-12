package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/Guaz-Delivery/guaz_backend/helpers"
	"github.com/Guaz-Delivery/guaz_backend/models"
	"github.com/Guaz-Delivery/guaz_backend/queries"
)

func HandleCourierSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var actionPayload models.Signup_Courier_ActionPayload
	if err := parseRequestBody(r.Body, &actionPayload); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Process signup
	result, err := SIGNUP_COURIER(actionPayload.Input, r.Header.Get("x-hasura-admin-secret"))
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Respond with JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
	}
}

func SIGNUP_COURIER(args models.SIGNUP_COURIERArgs, secret string) (models.Signup_Courier_Output, error) {
	// Hash password
	hashedPassword, err := hashPassword(args.Args.Password)
	if err != nil {
		return models.Signup_Courier_Output{}, err
	}

	// Prepare GraphQL variables
	variables := map[string]interface{}{
		"email":           args.Args.Email,
		"first_name":      args.Args.First_name,
		"middle_name":     args.Args.Middle_name,
		"last_name":       args.Args.Last_name,
		"location":        args.Args.Location,
		"phone_number":    args.Args.Phone_number,
		"rate":            args.Args.Rate,
		"shipment_range":  args.Args.Shipment_range,
		"shipment_size":   args.Args.Shipment_size,
		"profile_picture": args.Args.Profile_picture,
		"password":        hashedPassword,
	}

	// Send GraphQL request
	var regRes models.Response
	if err := sendGraphQLRequest(queries.SIGNUP_COURIER, variables, secret, &regRes); err != nil {
		return models.Signup_Courier_Output{}, err
	}

	// Generate JWT token
	token, err := helpers.GenerateJWTToken(regRes.Data.Insert_Couriers_One.Id, []string{"courier"})
	if err != nil {
		return models.Signup_Courier_Output{}, err
	}

	// Return success response
	message := "Successful"
	return models.Signup_Courier_Output{
		Token:      token,
		Courier_id: regRes.Data.Insert_Couriers_One.Id,
		Error:      false,
		Message:    &message,
	}, nil
}

// Helper function to parse request body
func parseRequestBody(body io.ReadCloser, dest interface{}) error {
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// Helper function to hash password
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Helper function to send GraphQL request
func sendGraphQLRequest(query string, variables map[string]interface{}, secret string, response interface{}) error {
	reqBody, err := json.Marshal(models.GraphQLRequest{Query: query, Variables: variables})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, os.Getenv("GRAPHQL_URL"), bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", secret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Printf("Signup response: %s", resByte)
	return json.Unmarshal(resByte, response)
}

// Helper function to generate JWT token

// Helper function to respond with error message
func responseWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.Signup_Courier_Output{
		Token:      "",
		Courier_id: "",
		Error:      true,
		Message:    &message,
	})
}
