package models

type Upload_input struct {
	File_name string `json:"file_name"`
	Base64    string `json:"base64"`
}

type Upload_output struct {
	Image_url string `json:"image_url"`
	Error     bool   `json:"error"`
}

type UploadActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            UPLOADArgs             `json:"input"`
}
type UPLOADArgs struct {
	Arg []Upload_input `json:"arg"`
}
