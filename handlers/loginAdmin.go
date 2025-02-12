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

func HandleAdminLogin(w http.ResponseWriter, r *http.Request) {

	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")
	// parse the body as action payload
	var actionPayload models.Login_Admin_ActionPayload
	if err := json.NewDecoder(r.Body).Decode(&actionPayload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Send the request params to the Action's generated handler function
	result, err := loginAdmin(actionPayload.Input, r.Header.Get("x-hasura-admin-secret"))

	if err != nil {
		helpers.AdminResponseWithError(w, err.Error())
		return
	}

	// Respond with JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
	}
}

// Auto-generated function that takes the Action parameters and must return it's response type
func loginAdmin(args models.LOGIN_ADMINArgs, secret string) (response interface{}, err error) {
	// fetch admins
	variables := map[string]interface{}{
		"email": args.Args.Email,
	}
	var checkRes models.Response
	err = helpers.SendGraphQLRequest(queries.LOGIN_ADMIN, variables, secret, &checkRes)
	if err != nil {
		return nil, err
	}
	if checkRes.Errors != nil {
		return nil, errors.New(checkRes.Errors[0].Message)
	}
	// compare the password
	log.Printf("%s", checkRes)
	admins := checkRes.Data.Admins
	if len(admins) == 0 || bcrypt.CompareHashAndPassword([]byte(admins[0].Password), []byte(args.Args.Password)) != nil {
		return nil, errors.New("authentication failed")
	}
	// generate token

	token, err := helpers.GenerateJWTToken(admins[0].Id, []string{"admin"})
	if err != nil {
		return nil, err
	}

	// send token
	response = models.Admin_Output{
		Token:    token,
		Admin_id: admins[0].Id,
		Error:    false,
		Message:  "sucess",
	}
	return response, nil
}
