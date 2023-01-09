package main

import (
	"flag"
	"fmt"
	"os"
	"sbom-tool/adapters/gradle"
	"sbom-tool/adapters/maven"
	"sbom-tool/adapters/npm"
	"sbom-tool/console"
	"sbom-tool/interfaces"
	"sbom-tool/structs"
)

var TOOLCHAINS = map[string]interfaces.ProcessBuilder{
	"maven":  &maven.Maven{},
	"npm":    &npm.NPM{},
	"gradle": &gradle.Gradle{},
}

var TOOLCHAINS_AVAILABLE = map[string]bool{
	"maven":  TOOLCHAINS["maven"].BuildToolsExist(),
	"npm":    TOOLCHAINS["npm"].BuildToolsExist(),
	"gradle": TOOLCHAINS["gradle"].BuildToolsExist(),
}

/*
Retrieves the adapter and check if the program support it

@param adapter
@return process_builder
*/
func toProcessBuilder(adapter string) interfaces.ProcessBuilder {

	var process_builder interfaces.ProcessBuilder

	if TOOLCHAINS_AVAILABLE[adapter] == true {

		if TOOLCHAINS[adapter] != nil {
			process_builder = TOOLCHAINS[adapter]
		} else {
			console.Error("main: Sorry, we didn't implement support for \"" + adapter + "\" yet!")
		}

	} else {
		console.Error("main: Please install the toolchain for \"" + adapter + "\"!")
	}

	return process_builder

}

func main() {

	cwd, err := os.Getwd()

	var projectDirectoryPath string
	var adapter string

	flag.StringVar(&adapter, "adapter", "", "Adapter selection. If not specified it will scan all existing adapters")
	flag.StringVar(&projectDirectoryPath, "projectDirectory", cwd, "Path to project directory")
	flag.Parse()

	if err == nil {

		var filesystem = structs.NewFilesystem(projectDirectoryPath)
		var resultInfos = filesystem.Scan()

		for f := 0; f < len(resultInfos); f++ {

			var process_builder interfaces.ProcessBuilder
			var resultInfo = resultInfos[f]

			if adapter == "" {
				process_builder = toProcessBuilder(resultInfo.Type)
			} else {
				process_builder = toProcessBuilder(adapter)
			}

			if process_builder != nil {
				// TODO make it as a gorountine to run all the adapters concurrently
				process_builder.Generate(resultInfo)
			}

		}

	} else {
		fmt.Println("You gotta be root, stoopid!")
	}
}
