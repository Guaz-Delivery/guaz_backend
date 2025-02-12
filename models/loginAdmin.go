package models

type Login_Admin_ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            LOGIN_ADMINArgs        `json:"input"`
}

type Login_Admin_Input struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Admin_Output struct {
	Token    string `json:"token"`
	Admin_id string `json:"admin_id"`
	Error    bool   `json:"error"`
	Message  string `json:"message"`
}

type LOGIN_ADMINArgs struct {
	Args Login_Admin_Input `json:"args"`
}

type Admin struct {
	Id           string `json:"id"`
	First_Name   string `json:"first_name"`
	Middle_Name  string `json:"middle_name"`
	Last_Name    string `json:"last_name"`
	Email        string `json:"email"`
	Phone_Number string `json:"phone_number"`
	Password     string `json:"password"`
	Created_At   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
