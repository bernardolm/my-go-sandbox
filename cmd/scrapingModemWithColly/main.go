package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

		c.OnHTML("input", func(e *colly.HTMLElement) {
			fmt.Printf("found\n%#v", e)
		})

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Authorization", "Basic "+base64.StdEncoding.
				EncodeToString([]byte("user:pass")))
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Printf("response\n%#v\n", string(r.Body))
		})

		c.Visit("http://192.168.0.1/?firewall_dmz")

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
