package interfaces

type ProcessBuilder interface {
	BuildToolsExist() bool
	Generate(file string) []byte
}
