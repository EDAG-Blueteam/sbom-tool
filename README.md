```bash
# Run the program within sbom tools project folder
$ go run main.go

# Run specifying project root path
$ go run main.go -projectDirectory=/path/to/project

# Generate a binary executable. Run this on linux machine to produce binary for linux/mac. Run this on window to produce .exe binary
$ go build

# Available flags
Usage of ./sbom-tool:
  -adapter string
        Adapter selection: npm, gradle, and maven
        If not specified it will scan all existing adapters
  -projectDirectory string
        Path to project directory
        If not specified it will use the current working directory of the sbom-tool
# Example for Linux
$ ./sbom-tool -projectDirectory=/path/to/project -adapter=gradle        
# Example for Windows
$ sbom-tool.exe -projectDirectory=/path/to/project -adapter=gradle     
```

Note for development :
1. Gradle version must be at least Gradle 4.4.1
2. Java version must be at least Java 8
3. Maven version must be at least Maven 3.8.6
4. Node version must be at least Node v16.15.0