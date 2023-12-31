package config

import (
	"testing"

	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
	"github.com/dll02/webgo/framework/provider/app"
	"github.com/dll02/webgo/framework/provider/env"
	tests "github.com/dll02/webgo/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebgoConfig_Normal(t *testing.T) {
	Convey("test Webgo config normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewWebgoContainer()
		c.Bind(&app.WebgoAppProvider{BaseFolder: basePath})
		c.Bind(&env.WebgoEnvProvider{})

		err := c.Bind(&WebgoConfigProvider{})
		So(err, ShouldBeNil)

		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		So(conf.GetString("database.mysql.hostname"), ShouldEqual, "127.0.0.1")
		So(conf.GetInt("database.mysql.timeout"), ShouldEqual, 1)
		So(conf.GetFloat64("database.mysql.readtime"), ShouldEqual, 2.3)
		// So(conf.GetString("database.mysql.password"), ShouldEqual, "mypassword")

		maps := conf.GetStringMap("database.mysql")
		So(maps, ShouldContainKey, "hostname")
		So(maps["timeout"], ShouldEqual, 1)

		maps2 := conf.GetStringMapString("databse.mysql")
		So(maps2["timeout"], ShouldEqual, "")

		type Mysql struct {
			Hostname string
			Username string
		}
		ms := &Mysql{}
		err = conf.Load("database.mysql", ms)
		Println(ms)
		So(err, ShouldBeNil)
		So(ms.Hostname, ShouldEqual, "127.0.0.1")
	})
}
