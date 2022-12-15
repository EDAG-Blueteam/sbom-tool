package maven

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sbom-tool/structs"
	"sbom-tool/utils"
)

type Maven struct {
}

func (m *Maven) Generate(file string) []byte {

	var result []byte

	buffer, err := os.ReadFile(file)

	if err == nil {

		var schema Schema

		// Unmarshall from xml to structs
		err2 := xml.Unmarshal(buffer, &schema)

		if err2 == nil {

			if len(schema.Build.Plugins) > 0 {

				var found Plugin

				for p := 0; p < len(schema.Build.Plugins); p++ {

					var plugin = schema.Build.Plugins[p]

					if plugin.GroupId == "org.cyclonedx" && plugin.ArtifactId == "cyclonedx-maven-plugin" {
						found = plugin
						break
					}

				}

				if found.GroupId != "org.cyclonedx" {

					schema.Build.Plugins = append(schema.Build.Plugins, Plugin{
						GroupId:    "org.cyclonedx",
						ArtifactId: "cyclonedx-maven-plugin",
						Version:    "2.7.3",
					})

				}

			} else {

				// Inject into Strutcs
				schema.Build.Plugins = append(schema.Build.Plugins, Plugin{
					GroupId:    "org.cyclonedx",
					ArtifactId: "cyclonedx-maven-plugin",
					Version:    "2.7.3",
				})

			}

		}

		// Marshall Structs to XML
		fmt.Println("New appended plugins structs: \n", schema)

		if pomXml, err := xml.MarshalIndent(schema, "", "    "); err == nil {

			//path := "SBOMWorkingDir"
			//if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			//	err := os.Mkdir(path, os.ModePerm)
			//	if err != nil {
			//		log.Println(err)
			//	}
			//}

			err := utils.CreateFolder("SBOMWorkingDir")
			if err != nil {
				fmt.Println("Unable to create SBOMWorkingDir !! ")
			}

			// Generate SBOM
			err = os.WriteFile("SBOMWorkingDir/pom.xml", pomXml, 0644)
			if err != nil {
				log.Fatal("Error writing modified injected pom.xml")
			} else {
				// TODO add checker if the SBOM successfully created
				utils.ExecCmd("mvn", "cyclonedx:makeBom", "-f", "SBOMWorkingDir/pom.xml")
			}
		} else {
			log.Fatal("Error marshalling XML")
		}

	}

	// TODO: Read maven config file into schema 			-> Done
	// TODO: Find out target folder (after build success)	-> Done
	// TODO: Inject cyclonedx maven plugin					-> Done
	// TODO: Run build process (and wait)					-> Done
	// TODO: Create the Sbom								-> Done
	// TODO: Verify that the sbom is created				-> Pending
	// TODO: Find out why the Sbom wasnt created
	// 		in Hazim										-> Pending
	// TODO: Read target/sbom.json							-> Pending

	//!!
	// TODO: merge multiple sbom and send to backend		-> Pending

	return result

}

func IsPackage(resultInfo structs.ResultInfo) bool {

	var result bool = false

	// TODO: Validate file for being a pom.xml -> Pending

	return result

}

func (m *Maven) BuildToolsExist() bool {

	output, err := utils.ExecCmd("mvn", "-v")

	if err != nil {
		log.Println("Cannot execute maven")
		return false
	}

	log.Println("The MAVEN version is : ", output)

	return true
}
