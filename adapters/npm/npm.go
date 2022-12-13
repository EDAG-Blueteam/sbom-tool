package npm

type NPM struct {
}

func (npm *NPM) Generate(file string) []byte {

	return nil
}

func IsPackage(file string) bool {

	var result bool = false

	return result

}

func (npm *NPM) BuildToolsExist() bool {
	//TODO implement me
	return true
}
