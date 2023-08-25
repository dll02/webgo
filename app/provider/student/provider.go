package student

import "github.com/dll02/goweb/framework"

type SubjectProvider struct {
	framework.ServiceProvider

	c framework.Container
}

func (sp *SubjectProvider) Name() string {
	return StudentKey
}

func (sp *SubjectProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (sp *SubjectProvider) IsDefer() bool {
	return false
}

func (sp *SubjectProvider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}

func (sp *SubjectProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}
