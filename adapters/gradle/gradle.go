package gradle

import (
	"log"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type Gradle struct {
}

func (m *Gradle) Generate(file string) []byte {
	return nil
}

func IsPackage(resultInfo structs.ResultInfo) bool {

	var result bool = false

	// TODO: validate gradle file

	return result

}

func (m *Gradle) BuildToolsExist() bool {

	output, err := utils.ExecCmd("gradle", "-v")

	if err != nil {
		log.Println("Cannot execute gradle")
		return false
	}

	log.Println("The gradle version is : ", output)

	return true
}
