package id

import (
	"github.com/rs/xid"
)

type WebgoIDService struct {
}

func NewWebgoIDService(params ...interface{}) (interface{}, error) {
	return &WebgoIDService{}, nil
}

func (s *WebgoIDService) NewID() string {
	return xid.New().String()
}
