package models

type Signup_Courier_ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            SIGNUP_COURIERArgs     `json:"input"`
}

type Signup_Courier_Input struct {
	First_name      string `json:"first_name"`
	Middle_name     string `json:"middle_name"`
	Last_name       string `json:"last_name"`
	Password        string `json:"password"`
	Location        string `json:"location"`
	Rate            int    `json:"rate"`
	Phone_number    string `json:"phone_number"`
	Email           string `json:"email"`
	Profile_picture string `json:"profile_picture"`
	Shipment_range  string `json:"shipment_range"`
	Shipment_size   string `json:"shipment_size"`
}

type SIGNUP_COURIERArgs struct {
	Args Signup_Courier_Input `json:"args"`
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
}
