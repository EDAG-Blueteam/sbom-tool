package gradle

import (
	"io"
	"log"
	"os"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type Gradle struct {
}

func (m *Gradle) Generate(file string) []byte {

	InjectGradlePlugin()

	utils.ExecCmd("gradle", "cyclonedxBom", "--build-file", "cyclonedxBuild.gradle")

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

func InjectGradlePlugin() {
	newGradleFile, err := os.Create("cyclonedxBuild.gradle")
	if err != nil {
		log.Fatal(err)
	}
	defer newGradleFile.Close()

	gradleFile, err2 := os.Open("./build.gradle")
	if err2 != nil {
		log.Fatal(err2)
	}

	defer gradleFile.Close()

	newGradleFile.WriteString("plugins {\n")
	newGradleFile.WriteString("\tid 'org.cyclonedx.bom' version '1.7.2'\n")
	newGradleFile.WriteString("}\n")

	_, err3 := io.Copy(newGradleFile, gradleFile)
	if err3 != nil {
		log.Fatal(err3)
	}
}
