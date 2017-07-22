package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/captncraig/stack"
)

var ctx = context.Background()

func printJson(i interface{}) {
	dat, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(dat))
}

func main() {
	c := stack.NewClient()
	sites, resp, err := c.Network().Sites().GetAllSites(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printJson(resp)
	printJson(sites)

}
