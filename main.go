package main

import (
	"fmt"
	"log"
	"os"
	"sbom-tool/adapters/gradle"
	"sbom-tool/adapters/maven"
	"sbom-tool/adapters/npm"
	"sbom-tool/interfaces"
	"sbom-tool/structs"
)

func main() {

	cwd, err := os.Getwd()

	if err == nil {

		var filesystem = structs.NewFilesystem(cwd)
		var resultInfos = filesystem.Scan()

		for f := 0; f < len(resultInfos); f++ {

			var processBuild interfaces.ProcessBuilder
			var resultInfo = resultInfos[f]

			fmt.Println(resultInfo)

			switch resultInfo.Type {
			case "maven":
				processBuild = &maven.Maven{}
			case "npm":
				processBuild = &npm.NPM{}
			case "gradle":
				processBuild = &gradle.Gradle{}
			// case "conan":

			// case "pypi":

			// case "rust":

			default:
				log.Println("file found does not match the provided metadata. damn")
				continue
			}

			if processBuild != nil {
				processBuild.Generate(resultInfo.Path)
			}

		}

	} else {
		fmt.Println("You gotta be root, stoopid!")
	}

}
