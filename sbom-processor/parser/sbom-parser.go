package parser

import (
	"encoding/json"
	"log"
	"os"
)

// parseCycloneDXSbom takes the sbom json file absolute path.
// It opens the file and parse the json
// It returns the sbom as a struct
func parseCycloneDxSbom(filePath string) (*CycloneDxSbom, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("%s cannot be opened", filePath)
		return nil, err
	}

	var sbom CycloneDxSbom
	if err := json.Unmarshal(content, &sbom); err != nil {
		log.Printf("cannot parse provided json file")
		return nil, err
	}

	return &sbom, nil
}
