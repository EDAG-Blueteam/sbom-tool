package parser

import "time"

type CycloneDxSbom struct {
	BomFormat    string           `json:"bomFormat"`
	SpecVersion  string           `json:"specVersion"`
	SerialNumber string           `json:"serialNumber"`
	Version      int              `json:"version"`
	Metadata     SbomMetadata     `json:"metadata"`
	Components   []SbomComponent  `json:"components"`
	Dependencies []SbomDependency `json:"dependencies"`
}

type SbomMetadata struct {
	Timestamp time.Time `json:"timestamp"`
	Tools     []struct {
		Vendor  string `json:"vendor"`
		Name    string `json:"name"`
		Version string `json:"version"`
		Hashes  []struct {
			Alg     string `json:"alg"`
			Content string `json:"content"`
		} `json:"hashes"`
	} `json:"tools"`
	Component struct {
		Group       string `json:"group"`
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description"`
		Licenses    []struct {
			License struct {
				ID string `json:"id"`
			} `json:"license"`
		} `json:"licenses"`
		Purl               string `json:"purl"`
		ExternalReferences []struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"externalReferences"`
		Type   string `json:"type"`
		BomRef string `json:"bom-ref"`
	} `json:"component"`
}

type SbomComponent struct {
	Publisher   string `json:"publisher,omitempty"`
	Group       string `json:"group"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Scope       string `json:"scope,omitempty"`
	Hashes      []struct {
		Alg     string `json:"alg"`
		Content string `json:"content"`
	} `json:"hashes"`
	Licenses []struct {
		License struct {
			ID string `json:"id"`
		} `json:"license"`
	} `json:"licenses"`
	Purl               string `json:"purl"`
	ExternalReferences []struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"externalReferences,omitempty"`
	Type   string `json:"type"`
	BomRef string `json:"bom-ref"`
}

type SbomDependency struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn"`
}
