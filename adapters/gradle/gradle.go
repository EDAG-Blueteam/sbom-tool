package gradle

import (
	"fmt"
	"io"
	"log"
	"os"
	"sbom-tool/console"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type Gradle struct {
}

func (m *Gradle) Generate(resultInfo structs.ResultInfo) []byte {

	InjectGradlePlugin(resultInfo.Path)

	workingDir := "SBOMWorkingDir/" + resultInfo.Uuid + "/"

	var shell = structs.NewShell(workingDir)

	var gradleArgument = []string{
		"cyclonedxBom",
		"--build-file",
		"SBOMWorkingDir/build.gradle",
	}
	_, err := shell.Execute("gradle", gradleArgument)
	// _, err := utils.ExecCmd("gradle", "cyclonedxBom", "--build-file", "SBOMWorkingDir/build.gradle")

	if err != nil {
		console.Error(err.Error())
	}

	return nil
}

func IsPackage(resultInfo structs.ResultInfo) bool {

	var result bool = false

	// TODO: validate gradle file

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

	// output, err := utils.ExecCmd("gradle", "-v")

	// if err != nil {
	// 	log.Println("Cannot execute gradle")
	// 	return false
	// }

	// log.Println("The gradle version is : ", output)

	return true
}

func InjectGradlePlugin(file string) {
	err := utils.CreateFolder("SBOMWorkingDir")
	if err != nil {
		fmt.Println("Unable to create SBOMWorkingDir !! ")
	}

	newGradleFile, err := os.Create("SBOMWorkingDir/build.gradle")
	if err != nil {
		log.Fatal(err)
	}
	defer newGradleFile.Close()

	gradleFile, err2 := os.Open(file)
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
