package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	logrus.SetLevel(logrus.PanicLevel)

	var userParam string
	var passParam string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Destination: &userParam,
			Name:        "user",
			Usage:       "user to login",
		},
		cli.StringFlag{
			Destination: &passParam,
			Name:        "pass",
			Usage:       "password to login",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var err error

		ctxt, cancel := context.WithCancel(context.Background())
		defer cancel()

		c, err := chromedp.New(ctxt, chromedp.WithRunnerOptions(
			runner.Flag("disable-gpu", true),
			// runner.Flag("headless", true),
			runner.Flag("incognito", true),
			runner.Flag("no-default-browser-check", true),
			runner.Flag("no-first-run", true),
		),
			chromedp.WithLog(logrus.Printf),
		)
		if err != nil {
			panic(err)
		}

		var res string
		err = c.Run(ctxt, chromedp.Tasks{
			chromedp.Navigate(`http://192.168.0.1`),
			chromedp.WaitVisible(`input[class="submitBtn"]`, chromedp.NodeVisible),
			chromedp.SendKeys(`input[id="UserName"]`, userParam, chromedp.NodeVisible),
			chromedp.SendKeys(`input[id="Password"]`, passParam, chromedp.NodeVisible),
			chromedp.Click("input.submitBtn", chromedp.NodeVisible),
			chromedp.WaitVisible(`ul[class="tabNavigation"]`, chromedp.NodeVisible),
			chromedp.Navigate(`http://192.168.0.1/?firewall_dmz`),
			chromedp.WaitVisible(`input[id="WANIP"]`, chromedp.NodeVisible),
			chromedp.Value(`input[id="WANIP"]`, &res, chromedp.NodeVisible),
		})
		if err != nil {
			panic(err)
		}

		err = c.Shutdown(ctxt)
		if err != nil {
			panic(err)
		}

		err = c.Wait()
		if err != nil {
			panic(err)
		}

		fmt.Print(res)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
