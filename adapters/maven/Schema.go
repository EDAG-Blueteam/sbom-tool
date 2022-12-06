package maven

import (
	"encoding/xml"
	"io"
)

type Build struct {
	// TODO: final name ?
	Plugins []Plugin `xml:"plugins>plugin"`
}

type Configuration struct {
	XMLName xml.Name `xml:"configuration"`
	Url     string   `xml:"url`
	Timeout int      `xml:"timeout"`
	Options []string `xml:"options>option"`
}

type DependencyManagement struct {
	Dependencies []Dependency `xml:"dependencies>dependency"`
}

type Dependency struct {
	XMLName    xml.Name    `xml:"dependency"`
	GroupId    string      `xml:"groupId"`
	ArtifactId string      `xml:"artifactId"`
	Version    string      `xml:"version"`
	Classifier string      `xml:"classifier"`
	Type       string      `xml:"type"`
	Scope      string      `xml:"scope"`
	Exclusions []Exclusion `xml:"exclusions>exclusion"`
}

type Exclusion struct {
	XMLName    xml.Name `xml:"exclusion"`
	GroupId    string   `xml:"groupId"`
	ArtifactId string   `xml:"artifactId"`
}

type Parent struct {
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

type Plugin struct {
	XMLName    xml.Name `xml:"plugin"`
	GroupId    string   `xml:"groupId"`
	ArtifactId string   `xml:"artifactId"`
	Version    string   `xml:"version"`

	// TODO: This might not work
	Configuration Configuration `xml:"configuration"`

	// https://maven.apache.org/guides/mini/guide-configuring-plugins.html
	// TODO something like: Configuration map[string]string `xml:"configuration"`
	// TODO executions

}

type PluginRepository struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
	Url  string `xml:"url"`
}

type Profile struct {
	Id    string `xml:"id"`
	Build Build  `xml:"build"`
}

type Properties map[string]string

type Repository struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
	Url  string `xml:"url"`
}

type Schema struct {
	XMLName              xml.Name             `xml:"project"`
	ModelVersion         string               `xml:"modelVersion"`
	Parent               Parent               `xml:"parent"`
	GroupId              string               `xml:"groupId"`
	ArtifactId           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	Packaging            string               `xml:"packaging"`
	Name                 string               `xml:"name"`
	Description          string               `xml:"description"`
	Repositories         []Repository         `xml:"repositories>repository"`
	Properties           Properties           `xml:"properties"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
	Dependencies         []Dependency         `xml:"dependencies>dependency"`
	Profiles             []Profile            `xml:"profiles"`
	Build                Build                `xml:"build"`
	PluginRepositories   []PluginRepository   `xml:"pluginRepositories>pluginRepository"`
	Modules              []string             `xml:"modules>module"`
}

func (properties *Properties) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {

	*properties = map[string]string{}

	for {

		var key string = ""
		var val string = ""

		token, err := decoder.Token()

		if err == io.EOF {
			break
		}

		switch tokentype := token.(type) {
		case xml.StartElement:

			key = tokentype.Name.Local
			err := decoder.DecodeElement(&val, &start)

			if err != nil {
				return err
			}

			(*properties)[key] = val

		}

	}

	return nil

}
