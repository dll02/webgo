package ssh

import (
	"github.com/dll02/webgo/framework/provider/config"
	"github.com/dll02/webgo/framework/provider/log"
	tests "github.com/dll02/webgo/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWebgoSSHService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.WebgoConfigProvider{})
	container.Bind(&log.WebgoLogServiceProvider{})

	Convey("test get client", t, func() {
		webgoRedis, err := NewWebgoSSH(container)
		So(err, ShouldBeNil)
		service, ok := webgoRedis.(*WebgoSSH)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-01"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
