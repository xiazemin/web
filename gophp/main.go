package main

import (
	"fmt"
	"github.com/taowen/go-php7/engine"
	"os"
)

func main() {

	engine.Initialize()
	ctx := &engine.Context{
		Output: os.Stdout,
	}
	err := engine.RequestStartup(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer engine.RequestShutdown(ctx)
	err = ctx.Exec("/Users/didi/xiazemin/php/main.php")
	if err != nil {
		fmt.Println(err)
	}

}
