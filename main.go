package main

import (
	"MyTinyGin/myGin"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//(1) render array
//http://localhost:8000/date
//hello, gin
//Date: 2021-08-06
//
//(2) custom render function
//http://localhost:9999/students
//   hello, gee
//   0: Gin is 24 years old
//   1: Tom is 25 years old
//
//
//(3) serve static files
//http://localhost:9999/assets/css/mygin.css
//   p {
//       color: orangered;
//       font-weight: 900;
//       font-size: 40px;
//   }


type student struct {
	Name string
	Age  int8
}

func FormatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := myGin.New()
	r.Use(myGin.Logger())
	r.SetFuncMap(template.FuncMap{ //engine.funcMap = funcMap 自定义模版渲染函数
		"FormatDate": FormatDate,  //指定模版中函数
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Gin", Age: 24}
	stu2 := &student{Name: "Tom", Age: 25}
	r.GET("/", func(c *myGin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *myGin.Context) {
		c.HTML(http.StatusOK, "myarray.tmpl", myGin.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *myGin.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", myGin.H{
			"title": "gin",
			"now":   time.Date(2021, 8, 06, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":8000")
}






//
//
//package main
//
//import (
//	"MyTinyGin/myGin"
//	"net/http"
//)
//
//func main() {
//	engin := myGin.New()
//	engin.GET("/", func(c *myGin.Context) {
//		c.HTML(http.StatusOK, "Hello,GIN")
//	})
//
//	engin.GET("/hello", func(c *myGin.Context) {
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//	})
//
//	engin.GET("/hello/:name", func(c *myGin.Context) {
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//	})
//
//	engin.GET("/assets/*filepath", func(c *myGin.Context) {
//		c.JSON(http.StatusOK, myGin.H{"filepath": c.Param("filepath")})
//	})
//
//	engin.Run(":8000")
//}



// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintf(w, "You are at %s", req.URL.Path)
// 	})
// 	http.HandleFunc("/header", func(w http.ResponseWriter, req *http.Request) {
// 		for k, v := range req.Header {
// 			fmt.Fprintf(w, "%v\t%v\n", k, v)
// 		}
// 	})
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

//--------------
// type Engine struct{}

// func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("request...")
// 	switch req.URL.Path {
// 	case "/":
// 		fmt.Fprintf(w, "You are at %s", req.URL.Path)
// 	case "/header":
// 		for k, v := range req.Header {
// 			fmt.Fprintf(w, "%v\t%v\n", k, v)
// 		}
// 	default:
// 		fmt.Fprint(w, "404 NOT FOUND")
// 	}
// }

// func main() {
// 	engine := new(Engine)
// 	log.Fatal(http.ListenAndServe(":8000", engine))
// }

//------------

// func main() {
// 	engine := myGin.New()
// 	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintf(w, "You are at %s", req.URL.Path)
// 	})
// 	engine.GET("/header", func(w http.ResponseWriter, req *http.Request) {
// 		for k, v := range req.Header {
// 			fmt.Fprintf(w, "%v\t%v\n", k, v)
// 		}
// 	})

// 	engine.Run(":8000")
// }

//------------

// func main() {
// 	engin := myGin.New()
// 	engin.GET("/", func(c *myGin.Context) {
// 		c.HTML(http.StatusOK, "Hello,Gin")
// 	})
// 	engin.GET("/hello", func(c *myGin.Context) {
// 		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
// 	})

// 	engin.POST("/login", func(c *myGin.Context) {
// 		c.JSON(http.StatusOK, myGin.H{
// 			"username": c.PostForm("username"),
// 			"password": c.PostForm("password"),
// 		})
// 	})

// 	engin.Run(":8000")
// }

//------------


