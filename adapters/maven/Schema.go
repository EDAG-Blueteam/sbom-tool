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
	XMLName xml.Name `xml:"configuration,omitempty"`
	Url     string   `xml:"url,omitempty"`
	Timeout int      `xml:"timeout,omitempty"`
	Options []string `xml:"options>option,omitempty"`
}

type DependencyManagement struct {
	Dependencies []Dependency `xml:"dependencies>dependency,omitempty"`
}

type Dependency struct {
	XMLName    xml.Name    `xml:"dependency,omitempty"`
	GroupId    string      `xml:"groupId,omitempty"`
	ArtifactId string      `xml:"artifactId,omitempty"`
	Version    string      `xml:"version,omitempty"`
	Classifier string      `xml:"classifier,omitempty"`
	Type       string      `xml:"type,omitempty"`
	Scope      string      `xml:"scope,omitempty"`
	Exclusions []Exclusion `xml:"exclusions>exclusion,omitempty"`
}

type Exclusion struct {
	XMLName    xml.Name `xml:"exclusion,omitempty"`
	GroupId    string   `xml:"groupId,omitempty"`
	ArtifactId string   `xml:"artifactId,omitempty"`
}

type Parent struct {
	GroupId      string `xml:"groupId,omitempty"`
	ArtifactId   string `xml:"artifactId,omitempty"`
	Version      string `xml:"version,omitempty"`
	RelativePath string `xml:"relativePath,omitempty"`
}

type Plugin struct {
	XMLName    xml.Name `xml:"plugin,omitempty"`
	GroupId    string   `xml:"groupId,omitempty"`
	ArtifactId string   `xml:"artifactId,omitempty"`
	Version    string   `xml:"version,omitempty,omitempty"`

	// TODO: This might not work
	Configuration Configuration `xml:"configuration,omitempty"`

	// https://maven.apache.org/guides/mini/guide-configuring-plugins.html
	// TODO something like: Configuration map[string]string `xml:"configuration"`
	// TODO executions

}

type PluginRepository struct {
	Id   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	Url  string `xml:"url,omitempty"`
}

type Profile struct {
	Id    string `xml:"id,omitempty"`
	Build Build  `xml:"build,omitempty"`
}

type Properties map[string]string

type Repository struct {
	Id   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	Url  string `xml:"url,omitempty"`
}

type Schema struct {
	XMLName              xml.Name             `xml:"project,omitempty"`
	ModelVersion         string               `xml:"modelVersion,omitempty"`
	Parent               Parent               `xml:"parent,omitempty"`
	GroupId              string               `xml:"groupId,omitempty"`
	ArtifactId           string               `xml:"artifactId,omitempty"`
	Version              string               `xml:"version,omitempty"`
	Packaging            string               `xml:"packaging,omitempty"`
	Name                 string               `xml:"name,omitempty"`
	Description          string               `xml:"description,omitempty"`
	Repositories         []Repository         `xml:"repositories>repository,omitempty"`
	Properties           Properties           `xml:"properties,omitempty"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement,omitempty"`
	Dependencies         []Dependency         `xml:"dependencies>dependency,omitempty"`
	Profiles             []Profile            `xml:"profiles,omitempty"`
	Build                Build                `xml:"build,omitempty"`
	PluginRepositories   []PluginRepository   `xml:"pluginRepositories>pluginRepository,omitempty"`
	Modules              []string             `xml:"modules>module,omitempty"`
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

func (properties Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range properties {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	return e.Flush()
}
