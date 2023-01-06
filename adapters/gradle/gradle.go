package gradle

import (
	"fmt"
	"io"
	"log"
	"os"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type Gradle struct {
}

func (m *Gradle) Generate(resultInfo structs.ResultInfo) []byte {
	workingDir := "SBOMWorkingDir/" + resultInfo.Uuid + "/"

	injectErr := InjectGradlePlugin(workingDir, resultInfo.Path)

	if injectErr != nil {
		log.Println("Error injecting cycloneDx plugin to build.gradle, ", injectErr)
	}

	var shell = structs.NewShell(workingDir)

	var gradleArgument = []string{
		"cyclonedxBom",
	}
	_, err := shell.Execute("gradle", gradleArgument)

	if err != nil {
		log.Println("Error executing sbom generation cmd for Gradle, ", err)
		return nil
	}

	result, err := os.ReadFile(workingDir + "build/reports/bom.json")

	if err != nil {
		log.Println("Error reading created bom.json, err:", err)
	}

	return result
}

func (m *Gradle) BuildToolsExist() bool {

	cwd, _ := os.Getwd()
	var shell = structs.NewShell(cwd)
	output, err := shell.Execute("gradle", []string{"-v"})

	if err != nil {
		log.Println("Cannot execute gradle")
		return false
	}

	log.Println("The gradle version is : ", output)

	return true
}

func InjectGradlePlugin(workingDir string, file string) error {
	err := utils.CreateFolder(workingDir)
	if err != nil {
		fmt.Println("Unable to create SBOMWorkingDir !! ")
		return err
	}

	newGradleFile, err := os.Create(workingDir + "build.gradle")
	if err != nil {
		return err
	}
	defer newGradleFile.Close()

	gradleFile, err2 := os.Open(file)
	if err2 != nil {
		return err
	}

	defer gradleFile.Close()

	newGradleFile.WriteString("plugins {\n")
	newGradleFile.WriteString("\tid 'org.cyclonedx.bom' version '1.7.2'\n")
	newGradleFile.WriteString("}\n")

	_, err3 := io.Copy(newGradleFile, gradleFile)
	if err3 != nil {
		return err
	}
	return nil
}
