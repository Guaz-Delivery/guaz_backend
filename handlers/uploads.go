package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Guaz-Delivery/guaz_backend/helpers"
	"github.com/Guaz-Delivery/guaz_backend/models"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	// set the response header as JSON
	w.Header().Set("Content-Type", "application/json")

	// parse the body as action payload
	var actionPayload models.UploadActionPayload
	if err := helpers.ParseRequestBody(r.Body, &actionPayload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Send the request params to the Action's generated handler function
	result := saveImageToFile(actionPayload.Input)

	// Write the response as JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
	}

}

func saveImageToFile(input models.UPLOADArgs) []models.Upload_output {

	var image []models.Upload_output = make([]models.Upload_output, len(input.Arg))
	for i, img := range input.Arg {
		// create a decoder with the base64 string from request
		dec, err := base64.StdEncoding.DecodeString(string(img.Base64))
		if err != nil {
			log.Println("unable to create base64 decoder")
			image[i].Error = true
			continue
		}

		dir, err := filepath.Abs("./upload")
		if err != nil {
			log.Println("unable to find upload's directory absolute path")
			image[i].Error = true
			continue
		}
		// create file and wait to close it after the function is about to return

		file, err := os.Create(filepath.Join(dir, img.File_name))
		if err != nil {

			log.Println("unable to create a file in the upload directory")
			image[i].Error = true
			continue
		}
		defer file.Close()
		// write the byte to the file
		if _, err = file.Write(dec); err != nil {

			log.Println("unable to write the byte to the file")
			image[i].Error = true
			continue
		}
		//  save the file
		if err := file.Sync(); err != nil {
			log.Println("unable to save the file")
			image[i].Error = true
			continue
		}
		if os.Getenv("DEBUG") != "" {
			image[i].Image_url = fmt.Sprintf(os.Getenv("TESTING_HOST_URL"), img.File_name)

		} else {
			image[i].Image_url = fmt.Sprintf(os.Getenv("HOST_URL"), img.File_name)
		}

	}

	return image
}
