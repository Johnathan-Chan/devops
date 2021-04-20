package artifactory

import "github.com/yametech/devops/pkg/service"

type Artifactory struct {
	service.IService
}

func NewArtifactory(i service.IService) *Artifactory {
	return &Artifactory{i}
}