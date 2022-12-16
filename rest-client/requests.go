package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

var (
	ErrSendingSbom       = errors.New("failed to send SBOM to EDAG service")
	ErrCreateSbomRequest = errors.New("failed to create SBOM request")
)

func sendSbom(restPath string, content []byte) error {
	request, err := http.NewRequest("POST", "http://some-dummy-url", bytes.NewBuffer(content))
	if err != nil {
		return ErrCreateSbomRequest
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(ErrSendingSbom)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		log.Printf("failed to send SBOM to EDAG service: Status code %s", response.Status)
		return ErrSendingSbom
	}

	if content, err = io.ReadAll(response.Body); err != nil {
		log.Println("Error parsing response")
		return err
	}

	var unstructuredResponseStruct map[string]interface{}
	if err := json.Unmarshal(content, &unstructuredResponseStruct); err != nil {
		log.Println("Error unmarshalling json")
		return err
	}

	log.Println("Response is:", unstructuredResponseStruct)

	return nil
}
