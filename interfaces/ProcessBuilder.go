package interfaces

import "sbom-tool/structs"

type ProcessBuilder interface {
	BuildToolsExist() bool
	Generate(resultInfo structs.ResultInfo) []byte
}
