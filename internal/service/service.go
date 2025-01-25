package service

import "github.com/gggallahad/diana/internal/project"

type (
	service struct {
	}
)

func NewService() project.Service {
	return &service{}
}
