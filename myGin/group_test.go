package myGin

import (
	"net/http"
	"testing"
)

//Postman
//http://localhost:8000/v2/login?username=gin&password=123456
//{
//"password": "123456",
//"username": "gin"
//}

func TestGroup(t *testing.T) {
	r := New()
	r.GET("/index", func(c *Context) {
		c.HTMLMiddleware(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1") // Group is defined to create a new RouterGroup
	{
		v1.GET("/", func(c *Context) {
			c.HTMLMiddleware(http.StatusOK, "<h1>Hello GIN</h1>")
		})

		v1.GET("/hello", func(c *Context) {
			// expect /hello?name=gin
			c.StringMiddleware(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *Context) {
			// expect /hello/gin
			c.StringMiddleware(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *Context) {
			c.JSON(http.StatusOK, H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":8000")
}
