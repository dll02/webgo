package redis

import (
	"context"
	"github.com/dll02/webgo/framework/provider/config"
	"github.com/dll02/webgo/framework/provider/log"
	tests "github.com/dll02/webgo/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestWebgoService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.WebgoConfigProvider{})
	container.Bind(&log.WebgoLogServiceProvider{})

	Convey("test get client", t, func() {
		webgoRedis, err := NewWebgoRedis(container)
		So(err, ShouldBeNil)
		service, ok := webgoRedis.(*WebgoRedis)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		ctx := context.Background()
		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
