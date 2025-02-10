package models

type Login_Input struct {
	Password     string  `json:"password"`
	Phone_Number *string `json:"phone_number"`
	Email        *string `json:"email"`
}

type Login_Output struct {
	Token      string `json:"token"`
	Courier_id string `json:"courier_id"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
}

type LOGIN_COURIERArgs struct {
	Args Login_Input `json:"args"`
}

type LoginCourierActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            LOGIN_COURIERArgs      `json:"input"`
}

type Couriers struct {
	Id           string `json:"id"`
	First_Name   string `json:"first_name"`
	Middle_Name  string `json:"middle_name"`
	Last_Name    string `json:"last_name"`
	Email        string `json:"email"`
	Phone_Number string `json:"phone_number"`
	Password     string `json:"password"`
	Rate         int    `json:"rate"`
	Location     string `json:"location"`
	Is_Verified  bool   `json:"is_verified"`
	Created_At   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
