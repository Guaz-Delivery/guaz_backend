package models

type Mutation struct {
	LOGIN_COURIER  *Login_Output
	SIGNUP_COURIER *Signup_Output
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLData struct {
	Couriers            []Couriers           `json:"couriers,omitempty"`
	Insert_Couriers_One RegisterResponseBody `json:"insert_couriers_one,omitempty"`
}

type Response struct {
	Data   GraphQLData    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}
