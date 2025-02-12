package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Guaz-Delivery/guaz_backend/helpers"
	"github.com/Guaz-Delivery/guaz_backend/models"
	"github.com/Guaz-Delivery/guaz_backend/queries"
	"golang.org/x/crypto/bcrypt"
)

func HandleCourierLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var actionPayload models.LoginCourierActionPayload
	if err := json.NewDecoder(r.Body).Decode(&actionPayload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	result, err := loginCourier(actionPayload.Input, r.Header.Get("x-hasura-admin-secret"))
	if err != nil {
		helpers.CourierResponseWithError(w, err.Error())
		return
	}

	// Respond with JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
	}
}

func loginCourier(args models.LOGIN_COURIERArgs, secret string) (interface{}, error) {
	variables := map[string]interface{}{
		"email":        args.Args.Email,
		"phone_number": args.Args.Phone_Number,
	}

	var checkRes models.Response
	err := helpers.SendGraphQLRequest(queries.LOGIN_COURIER, variables, secret, &checkRes)
	if err != nil {
		return nil, err
	}
	if checkRes.Errors != nil {
		return nil, errors.New(checkRes.Errors[0].Message)
	}

	couriers := checkRes.Data.Couriers
	log.Printf("%s", checkRes)
	if len(couriers) == 0 || bcrypt.CompareHashAndPassword([]byte(couriers[0].Password), []byte(args.Args.Password)) != nil {
		return nil, errors.New("authentication failed")
	}

	token, err := helpers.GenerateJWTToken(couriers[0].Id, []string{"courier"})
	if err != nil {
		return nil, err
	}

	return models.Courier_Output{
		Token:      token,
		Courier_id: couriers[0].Id,
		Error:      false,
		Message:    "successful",
	}, nil
}
