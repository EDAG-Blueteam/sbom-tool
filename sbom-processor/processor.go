package sbomprocessor

import (
	"encoding/json"
	"log"
	"os"
	"path"
	restclient "sbom-tool/rest-client"
	"sbom-tool/sbom-processor/parser"
)

type Sbom struct {
	workingDir string
}

func (s *Sbom) ProcessSbom() error {
	files, err := os.ReadDir(s.workingDir)
	if err != nil {
		log.Println(err)
		return err
	}

	var sbomList []parser.CycloneDxSbom
	for _, file := range files {
		fullPath := path.Join(s.workingDir, file.Name())
		parsedSbom, err := parser.ParseCycloneDxSbom(fullPath)
		if err != nil {
			log.Println(err)
			return err
		}

		sbomList = append(sbomList, *parsedSbom)
	}

	payload, err := json.Marshal(sbomList)
	if err != nil {
		log.Println(err)
	}

	restclient.SendSbom("/api/dummy/sbom", payload)
	return nil
}
