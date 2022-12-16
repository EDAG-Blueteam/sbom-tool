package npm

import (
	"log"
	"os"
	"path/filepath"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type NPM struct {
}

func (npm *NPM) Generate(resultInfo structs.ResultInfo) []byte {

	workingDir := "SBOMWorkingDir/" + resultInfo.Uuid + "/"
	err := utils.CreateFolder(workingDir)
	if err != nil {
		log.Printf("Unable to create %v !! ", workingDir)
	}

	input, err := os.ReadFile(resultInfo.Path)
	if err != nil {
		log.Println(err)

	}

	err = os.WriteFile(workingDir+"package.json", input, 0644)
	if err != nil {
		log.Println("Error creating file in SBOMWorkingDir/npmWorkingDir")
		log.Println(err)
	}

	var shell = structs.NewShell(filepath.Dir(workingDir + "package.json"))

	// Execute npm sbom installation command
	_, err = shell.Execute("npm", []string{"install", "--save-dev", "@cyclonedx/cyclonedx-npm"})

	// Execute npm sbom generation command
	_, err = shell.Execute("npx", []string{"@cyclonedx/cyclonedx-npm", "--output-file", "sbom.json"})
	if err != nil {
		log.Println("Error executing npm/npx shell command, err:", err)
	}
	result, err := os.ReadFile(workingDir + "sbom.json")
	if err != nil {
		log.Println("Error reading created bom.json, err:", err)
	}
	return result
}

func (npm *NPM) BuildToolsExist() bool {

	output, err := utils.ExecCmd("npm", "-v")

	if err != nil {
		log.Println("Cannot execute npm")
		return false
	}

	log.Println("The NPM version is : ", output)

	return true
}
