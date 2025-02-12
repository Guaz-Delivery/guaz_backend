package models

type Mutation struct {
	LOGIN_ADMIN    *Admin_Output
	LOGIN_COURIER  *Courier_Output
	SIGNUP_COURIER *Courier_Output
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLData struct {
	Couriers            []Couriers           `json:"couriers"`
	Admins              []Admin              `json:"admin"`
	Insert_Couriers_One RegisterResponseBody `json:"insert_couriers_one"`
}

type Response struct {
	Data   GraphQLData    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

type Courier_Output struct {
	Token      string `json:"token"`
	Courier_id string `json:"courier_id"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
}
