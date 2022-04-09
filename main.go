package main

import (
	"./lib"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/urfave/cli"
	"os"
	"time"
)

func scenarioRun(config string, dist string)  {
	c :=  lib.LoadConfig(config)

	var successStream, failStream string

	for _, param := range c.Scenario.Params {
		// TODO; handle error
		var expect, _ = lib.Request(&lib.RequestOption{
			Method: c.Scenario.Method,
			Url: c.Expect.Url+"?"+param,
			Header: c.Expect.Header,
			Cookie: c.Expect.Cookie,
		})

		// TODO; handle error
		var actual, _ = lib.Request(&lib.RequestOption{
			Method: c.Scenario.Method,
			Url: c.Actual.Url+"?"+param,
			Header: c.Actual.Header,
			Cookie: c.Actual.Cookie,
		})

		// key sort
		expect = lib.Sort(expect)
		actual = lib.Sort(actual)

		differ := diffmatchpatch.New()
		diff := differ.DiffMain(expect, actual, false)
		diff = differ.DiffCleanupEfficiency(diff)

		var diffStream string
		var equal = true

		for _, line := range diff {
			switch line.Type {
			case diffmatchpatch.DiffDelete:
				equal = false
				diffStream += "<span style=\"color: red; \">~~" + line.Text +"~~</span>"
			case diffmatchpatch.DiffInsert:
				equal = false
				diffStream += "<span style=\"color: green; \">" + line.Text +"</span>"
			default:
				diffStream += line.Text
			}
		}

		if equal {
			fmt.Printf("\x1b[32m%s\x1b[0m", "✓ ▶ "+ param +"\n")
			successStream += fmt.Sprintf("- %s \n", param)
		} else {
			fmt.Printf("\x1b[31m%s\x1b[0m", "☓ ▶ "+ param + "\n")
			failStream += fmt.Sprintf("```%s```   \n\n", param)
			failStream += fmt.Sprintf("%s  \n\n", diffStream)

		}
		time.Sleep(time.Second * 1)
	}

	// report format
	out := "# diff result  \n"
	out += fmt.Sprintf("## success \n%s \n", successStream)
	out += fmt.Sprintf("## fail \n%s \n", failStream)

	file, err := os.Create(dist)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte(out)
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "CLI for validating delta comparisons of API responses"
	app.Description = "Compare API response differences"
	app.Flags = []cli.Flag {
		&cli.StringFlag {
			Name:        "c",
			Usage:       "config file path",
			Required:    true,
		},

		&cli.StringFlag {
			Name:        "o",
			Usage:       "dist report path",
			Required:    true,
		},
	}

	app.Action = func(c *cli.Context) error {
		config := c.String("c")
		out := c.String("o")

		scenarioRun(config, out)
		return nil
	}


	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
