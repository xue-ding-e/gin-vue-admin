package middleware

import (
	"fmt"

	"github.com/unrolled/secure"
)

// 用https把这个中间件在router里面use一下就好

func LoadTls() fiber.Handler {
	return func(c *fiber.Ctx) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
