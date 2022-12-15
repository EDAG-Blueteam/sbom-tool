package npm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type NPM struct {
}

func (npm *NPM) Generate(file string) []byte {

	err := utils.CreateFolder("SBOMWorkingDir/npmWorkingDir")
	if err != nil {
		fmt.Println("Unable to create SBOMWorkingDir/npmWorkingDir !! ")
	}

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)

	}

	err = os.WriteFile("SBOMWorkingDir/npmWorkingDir/package.json", input, 0644)
	if err != nil {
		fmt.Println("Error creating file in SBOMWorkingDir/npmWorkingDir")
		fmt.Println(err)
	}

	var shell = structs.NewShell(filepath.Dir("SBOMWorkingDir/npmWorkingDir/package.json"))

	// Execute npm sbom installation command
	_, err = shell.Execute("npm", []string{"install", "--save-dev", "@cyclonedx/cyclonedx-npm"})

	// Execute npm sbom generation command
	_, err = shell.Execute("npx", []string{"@cyclonedx/cyclonedx-npm", "--output-file", "sbom.json"})

	fmt.Println(err)

	return nil
}

func IsPackage(file string) bool {

	var result bool = false

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
