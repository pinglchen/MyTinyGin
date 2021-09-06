# MyTinyGin
动手实现Gin框架的核心功能


engine实现了ServeHTTP,可以传入ListenAndServe: http.ListenAndServe("8000",engine) 并由engine.Run()封装
这样可以调用engine.ServeHTTP  在engine.ServeHTTP里又调用了添加路由自定义的handlerfunc函数

ServeHTTP(w http.ResponseWriter, req *http.Request)的两个参数封装进了Context中 HandlerFunc调用时
参数是Context类型，再由HandlerFunc中的c.JSON/c.HTML处理，主要是传导http.ResponseWriter，*http.Request

engine封装了路由router，由router.handler方法调用HandlerFunc
router封装了roots（map[string]*node）和handlers（map[string]HandlerFunc）
router.roots[method] /GET 每种方法对应一个trie树
router.handlers[key] key=method+'-'+pattern 对应路由的处理方法
