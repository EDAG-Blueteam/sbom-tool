package gradle

import "sbom-tool/structs"

type Gradle struct {
}

func (m *Gradle) BuildToolsExist() bool {
	return true
}

func (m *Gradle) Generate(resultInfo structs.ResultInfo) []byte {
	return nil
}
