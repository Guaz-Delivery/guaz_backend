package models

type SignupCourierActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            SIGNUP_COURIERArgs     `json:"input"`
}

type Signup_Input struct {
	First_name   string `json:"first_name"`
	Middle_name  string `json:"middle_name"`
	Last_name    string `json:"last_name"`
	Password     string `json:"password"`
	Location     string `json:"location"`
	Rate         int    `json:"rate"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type Signup_Output struct {
	Token      string  `json:"token"`
	Courier_id string  `json:"courier_id"`
	Error      bool    `json:"error"`
	Message    *string `json:"message"`
}

type Mutation struct {
	SIGNUP_COURIER *Signup_Output `json:"signup_output"`
}

type SIGNUP_COURIERArgs struct {
	Args Signup_Input `json:"args"`
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type RegisterResponseBody struct {
	Created_At   string `json:"created_at"`
	Email        string `json:"email"`
	First_Name   string `json:"first_name"`
	Id           string `json:"id"`
	Is_Verified  bool   `json:"is_verified"`
	Last_Name    string `json:"last_name"`
	Location     string `json:"location"`
	Middle_Name  string `json:"middle_name"`
	Phone_Number string `json:"phone_number"`
	Rate         int    `json:"rate"`
	Updated_at   string `json:"updated_at"`
	/*
			"created_at": "2025-02-10T13:56:30.546974+00:00",
		      "email": "test3",
		      "first_name": "test",
		      "id": "9028843b-1e15-431f-a734-3d8e0e03d53f",
		      "is_verified": false,
		      "last_name": "test",
		      "location": "(80,80)",
		      "middle_name": "test",
		      "phone_number": "3test",
		      "rate": 2,
		      "updated_at": "2025-02-10T13:56:30.546974+00:00"
	*/
}

type RegisterGraphQLData struct {
	Insert_Couriers_One RegisterResponseBody `json:"insert_couriers_one"`
}

type RegisterResponse struct {
	Data   RegisterGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError      `json:"errors,omitempty"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type Point struct {
	Y float64 `json:"y"` // latitude
	X float64 `json:"x"` // longitude
}
