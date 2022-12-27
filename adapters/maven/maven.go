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

func (m *Maven) Generate(resultInfo structs.ResultInfo) []byte {

	var result []byte

	buffer, err := os.ReadFile(resultInfo.Path)

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

			workingDir := "SBOMWorkingDir/" + resultInfo.Uuid + "/"
			err := utils.CreateFolder(workingDir)
			if err != nil {
				fmt.Println("Unable to create SBOMWorkingDir !! ")
			}

			// Generate SBOM
			err = os.WriteFile(workingDir+"pom.xml", pomXml, 0644)
			if err != nil {
				log.Fatal("Error writing modified injected pom.xml")
			} else {
				var shell = structs.NewShell(workingDir)
				// TODO add checker if the SBOM successfully created
				_, err := shell.Execute("mvn", []string{"cyclonedx:makeBom"})
				if err != nil {
					log.Println("Error executing maven cyclonedx command, err:", err)
				}
				result, err = os.ReadFile(workingDir + "/target/bom.json")
				if err != nil {
					log.Println("Error reading created bom.json, err:", err)
				}
			}
		} else {
			log.Fatal("Error marshalling XML")
		}

	}

	return result

}

func (m *Maven) BuildToolsExist() bool {

	cwd, _ := os.Getwd()
	var shell = structs.NewShell(cwd)
	output, err := shell.Execute("mvn", []string{"-v"})

	if err != nil {
		log.Println("Cannot execute maven")
		return false
	}

	log.Println("The MAVEN version is : ", output)

	return true
}
