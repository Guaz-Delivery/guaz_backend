package models

type Payment struct {
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency,"`
	Email         string                 `json:"email"`
	Phone_Number  string                 `json:"phone_number"`
	Tx_Ref        string                 `json:"tx_ref"`
	Courier_id    string                 `json:"courier_id"`
	Callback_URL  string                 `json:"callback_url"`
	Return_URL    string                 `json:"return_url"`
	Customization map[string]interface{} `json:"customization"`
}

type ChapaResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		Checkout_URL string `json:"checkout_url"`
	} `json:"data"`
}
type PaymentActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CREDIT_PAYMENTArgs     `json:"input"`
}

type CREDIT_PAYMENTArgs struct {
	Args Credit_Input `json:"args"`
}

type Credit_Input struct {
	Email        string  `json:"email"`
	Phone_number string  `json:"phone_number"`
	Amount       float32 `json:"amount"`
	Courier_id   string  `json:"courier_id"`
	Delivery_id  string  `json:"delivery_id"`
	Return_URL   string  `json:"return_url"`
}

type Credit_Output struct {
	Checkout_url string `json:"checkout_url"`
	Status       string `json:"status"`
}
