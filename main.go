package main

import (
	"fmt"
	"os"
	"sbom-tool/adapters/maven"
	"sbom-tool/adapters/npm"
	"sbom-tool/console"
	"sbom-tool/interfaces"
	"sbom-tool/structs"
	"strings"
)

var TOOLCHAINS = map[string]interfaces.ProcessBuilder{
	"maven": &maven.Maven{},
	"npm":   &npm.NPM{},
}

var TOOLCHAINS_AVAILABLE = map[string]bool{
	"maven": TOOLCHAINS["maven"].BuildToolsExist(),
	"npm":   TOOLCHAINS["npm"].BuildToolsExist(),
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

	var adapter string

	if len(os.Args) == 2 {
		// sbom-tool <adapter>
		adapter = strings.TrimSpace(os.Args[1])
	}

	cwd, err := os.Getwd()

	if err == nil {

		var filesystem = structs.NewFilesystem(cwd)
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
				process_builder.Generate(resultInfo.Path)
			}

		}

	} else {
		fmt.Println("You gotta be root, stoopid!")
	}
}
