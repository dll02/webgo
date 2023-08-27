package env

import (
	tests "github.com/dll02/webgo/test"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
	"github.com/dll02/webgo/framework/provider/app"
)

func TestWebgoEnvProvider(t *testing.T) {
	Convey("test Webgo env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewWebgoContainer()
		sp := &app.WebgoAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &WebgoEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}
