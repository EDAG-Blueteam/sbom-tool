package main

import (
	"fmt"
	"os"
	"sbom-tool/structs"
)


func main() {

	cwd, err := os.Getwd();

	if err == nil {

		var filesystem = structs.NewFilesystem(cwd);
		var files      = filesystem.Scan();
	
	
		for f := 0; f < len(files); f++ {
	
			fmt.Println(files[f]);
	
		}
		
	} else {
		fmt.Println("You gotta be root, stoopid!");
	}
	
}
