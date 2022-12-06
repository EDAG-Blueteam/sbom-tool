package interfaces

type ProcessBuilder interface {

	Generate(file string) []byte

}