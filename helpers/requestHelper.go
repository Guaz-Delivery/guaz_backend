package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Guaz-Delivery/guaz_backend/models"
)

// Helper function to send GraphQL request
func SendGraphQLRequest(query string, variables map[string]interface{}, secret string, response interface{}) error {
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

// Helper function to parse request body
func ParseRequestBody(body io.ReadCloser, dest interface{}) error {
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}
