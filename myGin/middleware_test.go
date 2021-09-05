package myGin

import (
	"log"
	"net/http"
	"testing"
	"time"
)


//=== RUN   TestMiddleware
//2021/08/06 02:09:40 Route  GET - /
//2021/08/06 02:09:40 Route  GET - /v2/hello/:name
//2021/08/06 02:09:42 [200] / in 23.042µs
//2021/08/06 02:09:44 [200] / in 7.625µs
//2021/08/06 02:09:46 [200] / in 7µs
//2021/08/06 02:10:04 [500] /v2/hello/gin in 107.583µs for group v2
//2021/08/06 02:10:04 [500] /v2/hello/gin in 134.75µs



func middlewareForV2() MyHandleFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}



func TestMiddleware(t *testing.T){
	r := New()
	r.Use(Logger()) // global middleware
	r.GET("/", func(c *Context) {
		c.HTMLMiddleware(http.StatusOK, "<h1>Hello,GIN</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(middlewareForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *Context) {
			// expect /hello/gin
			c.StringMiddleware(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":8000")
}