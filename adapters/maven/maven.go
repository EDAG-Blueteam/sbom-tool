package maven

import (
	"sbom-tool/structs"
	"encoding/xml"
	"fmt"
	"os"
)

type Maven struct {

}

func (m *Maven) Generate(file string) []byte {

	var result []byte;

	buffer, err := os.ReadFile(file)

	if err == nil {

		var schema Schema

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

				schema.Build.Plugins = append(schema.Build.Plugins, Plugin{
					GroupId:    "org.cyclonedx",
					ArtifactId: "cyclonedx-maven-plugin",
					Version:    "2.7.3",
				})

			}

		}


		fmt.Println(schema)

	}
	

	// TODO: Read maven config file into schema
	// TODO: Find out target folder (after build success)
	// TODO: Inject cyclonedx maven plugin
	// TODO: Run build process (and wait)
	// TODO: Read target/bom.json

	return result

}

func IsPackage(resultInfo structs.ResultInfo) bool {

	var result bool = false

	// TODO: Validate file for being a pom.xml

	return result

}
