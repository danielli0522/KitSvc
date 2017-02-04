package middleware

import (
	"github.com/TeaMeow/KitSvc/module/mq"
	"github.com/TeaMeow/KitSvc/module/mq/mqstore"
	"github.com/TeaMeow/KitSvc/shared/mqutil"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

func MQ(c *cli.Context, m *mqutil.Engine) gin.HandlerFunc {
	v := setupMQ(c, m)
	return func(c *gin.Context) {
		mq.ToContext(c, v)
		c.Next()
	}
}

func setupMQ(c *cli.Context, m *mqutil.Engine) mq.MQ {
	return mqstore.NewProducer(
		c.String("url"),
		c.String("nsq-producer"),
		c.String("nsq-producer-http"),
		c.StringSlice("nsq-lookupds"),
		m,
	)
}
