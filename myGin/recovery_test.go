package myGin


//http://localhost:8000"
//Hello,GIN
//
//http://localhost:8000/panic
//=== RUN   TestRecovery
//2021/08/06 01:52:02 Route  GET - /
//2021/08/06 01:52:02 Route  GET - /panic
//2021/08/06 01:52:09 [200] / in 29.917µs
//2021/08/06 01:52:17 runtime error: index out of range [10] with length 1
//Traceback:
///usr/local/go/src/runtime/panic.go:965
///usr/local/go/src/runtime/panic.go:88
///Users/chen/GolandProjects/MyTinyGin/myGin/recovery_test.go:16
///Users/chen/GolandProjects/MyTinyGin/myGin/context.go:41
///Users/chen/GolandProjects/MyTinyGin/myGin/recovery.go:37
///Users/chen/GolandProjects/MyTinyGin/myGin/context.go:41
///Users/chen/GolandProjects/MyTinyGin/myGin/logger.go:15
///Users/chen/GolandProjects/MyTinyGin/myGin/context.go:41
///Users/chen/GolandProjects/MyTinyGin/myGin/router.go:99
///Users/chen/GolandProjects/MyTinyGin/myGin/mygin.go:130
///usr/local/go/src/net/http/server.go:2868
///usr/local/go/src/net/http/server.go:1933
///usr/local/go/src/runtime/asm_arm64.s:1131
//
//2021/08/06 01:52:17 [500] /panic in 251.292µs




import (
	"net/http"
	"testing"
)

func TestRecovery(t *testing.T) {
	r := Default()
	r.GET("/", func(c *Context) {
		c.String(http.StatusOK, "Hello,GIN\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Context) {
		names := []string{"MyTinyGin"}
		c.String(http.StatusOK, names[10])
	})

	r.Run(":8000")
}
