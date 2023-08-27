package config

import (
	"github.com/dll02/webgo/framework"
	"path/filepath"
	"testing"

	"github.com/dll02/webgo/framework/contract"
	tests "github.com/dll02/webgo/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebgoConfig_GetInt(t *testing.T) {
	Convey("test Webgo env normal case", t, func() {
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		c := framework.NewWebgoContainer()
		serv, err := NewWebgoConfig(folder, map[string]string{}, contract.EnvDevelopment, c)
		So(err, ShouldBeNil)
		conf := serv.(*WebgoConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
