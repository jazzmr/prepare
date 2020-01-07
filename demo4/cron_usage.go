package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		parse *cronexpr.Expression
		err   error
	)

	if parse, err = cronexpr.Parse("* * * * * * *"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(parse.Next(time.Now()))
}
