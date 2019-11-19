package main

import (
	"encoding/json"
	"go-web/src/db"
	"go-web/src/godis"
	"go-web/src/utils"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("error")
	utils.SetLog(app.Logger())
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
		// godis.GetValue("")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/setJson", func(ctx iris.Context) {
		u := &User{1, "hello", 18}
		uStr, err := json.Marshal(u)
		if err != nil {
			panic(err)
		}
		godis.SetValue("user", string(uStr))
		ctx.WriteString("setJson")
	})

	app.Get("/getJson", func(ctx iris.Context) {
		uStr := godis.GetValue("user")
		var u User
		err := json.Unmarshal([]byte(uStr), &u)
		if err != nil {
			panic(err)
		}
		ctx.JSON(u)
	})

	app.Get("/getDbData", func(ctx iris.Context) {
		ctx.JSON(db.QueryMulti())
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
