package structs

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"strings"
)

type ResultInfo struct {
	Path string `json:"path"`
	Type string `json:"type"`
	Uuid string `json:"uuid"`
}

func CreateProjectUuid(path string) string {
	uuid, _ := uuid2.NewUUID()
	pathUuid := strings.Replace(path, ":", "", -1)
	pathUuid = strings.Replace(pathUuid, "\\", "-", -1)
	pathUuid = strings.Replace(pathUuid, "/", "-", -1)
	pathUuid = strings.TrimPrefix(pathUuid, "-")

	projectUuid := fmt.Sprintf("%s-%s", pathUuid, uuid.String())
	return projectUuid
}
