package id

import (
	"testing"

	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
	"github.com/dll02/webgo/framework/provider/app"
	"github.com/dll02/webgo/framework/provider/config"
	"github.com/dll02/webgo/framework/provider/env"
	"github.com/dll02/webgo/framework/util"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test Webgo console log normal case", t, func() {
		basePath := util.GetExecDirectory()
		c := framework.NewWebgoContainer()
		c.Bind(&app.WebgoAppProvider{BaseFolder: basePath})
		c.Bind(&env.WebgoEnvProvider{})
		c.Bind(&config.WebgoConfigProvider{})

		err := c.Bind(&WebgoIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		t.Log(xid)
		So(xid, ShouldNotBeEmpty)
	})
}
