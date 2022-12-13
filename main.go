package main

import (
	"fmt"
	"log"
	"os"
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
		var adapterMap = make(map[string]bool)

		for _, v := range resultInfos {
			adapterMap[v.Type] = false
		}

		// Check which build tool exist in filetype
		for k := range adapterMap {
			var processBuild interfaces.ProcessBuilder

			switch k {
			case "maven":
				processBuild = &maven.Maven{}
			case "npm":
				processBuild = &npm.NPM{}
			default:
				log.Println("file found does not match the provided metadata. damn")
				continue
			}
			adapterMap[k] = processBuild.BuildToolsExist()
		}

		for f := 0; f < len(resultInfos); f++ {

			var processBuild interfaces.ProcessBuilder
			var resultInfo = resultInfos[f]

			fmt.Println(resultInfo)

			// Check which build tool is installed
			if !adapterMap[resultInfo.Type] {
				log.Printf("%s is not installed on your machine !", resultInfo.Type)
				continue
			}

			switch resultInfo.Type {
			case "maven":
				processBuild = &maven.Maven{}
			case "npm":
				processBuild = &npm.NPM{}
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
