package restclient

import (
	SbomModels "sbom-tool/sbom-processor/parser"
)

type SbomRequest struct {
	SbomModels.CycloneDxSbom `json:"sbom"`
}
