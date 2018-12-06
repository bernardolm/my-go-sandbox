package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://192.168.0.1/?firewall_dmz", nil)
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.
			EncodeToString([]byte("user:pass")))

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			panic(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			panic(err)
		}

		doc.Find("input#WANIP").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("found\n%+v\n", s.Text())
		})

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
