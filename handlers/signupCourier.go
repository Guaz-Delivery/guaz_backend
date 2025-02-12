package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Guaz-Delivery/guaz_backend/helpers"
	"github.com/Guaz-Delivery/guaz_backend/models"
	"github.com/Guaz-Delivery/guaz_backend/queries"
)

func HandleCourierSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var actionPayload models.Signup_Courier_ActionPayload
	if err := helpers.ParseRequestBody(r.Body, &actionPayload); err != nil {
		helpers.CourierResponseWithError(w, "Invalid request body")
		return
	}

	// Process signup
	result, err := signupCourier(actionPayload.Input, r.Header.Get("x-hasura-admin-secret"))
	if err != nil {
		helpers.CourierResponseWithError(w, err.Error())
		return
	}

	// Respond with JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
	}
}

func signupCourier(args models.SIGNUP_COURIERArgs, secret string) (interface{}, error) {
	// Hash password
	hashedPassword, err := helpers.HashPassword(args.Args.Password)
	if err != nil {
		return nil, err
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
	if err := helpers.SendGraphQLRequest(queries.SIGNUP_COURIER, variables, secret, &regRes); err != nil {
		return nil, err
	}
	if regRes.Errors != nil {
		log.Printf("signup: %s", regRes.Errors[0].Message)
		return nil, errors.New("already registered!")
	}

	// Generate JWT token
	token, err := helpers.GenerateJWTToken(regRes.Data.Insert_Couriers_One.Id, []string{"courier"})
	if err != nil {
		return nil, err
	}

	// Return success response
	return models.Courier_Output{
		Token:      token,
		Courier_id: regRes.Data.Insert_Couriers_One.Id,
		Error:      false,
		Message:    "successful",
	}, nil
}
