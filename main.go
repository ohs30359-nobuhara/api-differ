package main

import (
	"bytes"
	"diff-api/lib"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"time"
)

func scenarioRun(config string, dist string)  {
	c :=  lib.LoadConfig(config)

	// create result dir
	if err := os.Mkdir(dist, 0777); err != nil {
		fmt.Println(err)
	}

	for index, param := range c.Scenario.Params {
		time.Sleep(time.Second * 1)

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

		// response sort
		expect = lib.Sort(expect)
		actual = lib.Sort(actual)

		differ := diffmatchpatch.New()
		diff := differ.DiffMain(expect, actual, false)
		diff = differ.DiffCleanupEfficiency(diff)

		var b bytes.Buffer
		var equal = true

		for _, line := range diff {
			switch line.Type {
			case diffmatchpatch.DiffDelete:
				equal = false
				b.WriteString(fmt.Sprintf("<span style=\"color: red; font-size: 2em\">~~%s~~</span>", line.Text))
			case diffmatchpatch.DiffInsert:
				equal = false
				b.WriteString(fmt.Sprintf("<span style=\"color: green; font-size: 2em\">~~%s~~</span>", line.Text))
			default:
				b.WriteString(line.Text)
			}
		}

		if equal {
			fmt.Printf("\x1b[32m%s\x1b[0m", fmt.Sprintf("[GREEN] %s \n", param))
			continue
		}
		fmt.Printf("\x1b[31m%s\x1b[0m", fmt.Sprintf("[ RED ] %s \n", param))

		// create diff file to "xxx/diff_x.md"
		file, err := os.Create(fmt.Sprintf("%s/diff_%s.md", dist, strconv.Itoa(index))); if err != nil {
			panic(err.Error())
		}
		_, err = file.WriteString(b.String()); if err != nil {
			panic(err.Error())
		}
		file.Close()
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
			Value:       "report",
			Required:    false,
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
