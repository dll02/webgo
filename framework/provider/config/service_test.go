package config

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
	"github.com/dll02/webgo/framework/provider/app"
	"github.com/dll02/webgo/framework/provider/env"
	"path/filepath"
	"testing"

	tests "github.com/dll02/webgo/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebgoConfig_GetInt(t *testing.T) {
	Convey("test Webgo env normal case", t, func() {
		basePath := tests.BasePath
		configFolder := filepath.Join(basePath, "config")

		c := framework.NewWebgoContainer()
		c.Bind(&app.WebgoAppProvider{BaseFolder: basePath})
		c.Bind(&env.WebgoEnvProvider{})
		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		env := envServ.AppEnv()
		envFolder := filepath.Join(configFolder, env)
		serv, err := NewWebgoConfig(c, envFolder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*WebgoConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
