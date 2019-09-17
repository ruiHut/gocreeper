package parser

import (
	"crepper/zhengai/engine"
	"fmt"
	"github.com/ericchiang/css"
	"golang.org/x/net/html"
	"strings"
)

const (
	URL_AND_ADDRESS = "article.g-container>dl>dd>a"
)

func ParseCityList(contexts []byte) engine.ParseResult {
	sel, err := css.Compile(URL_AND_ADDRESS)
	if err != nil {
		fmt.Printf("css compile error:  %s\n", err)
	}

	node, err := html.Parse(strings.NewReader(string(contexts)))
	if err != nil {
		fmt.Printf("html parse error:  %s\n", err)
	}

	result := engine.ParseResult{}
	for _, ele := range sel.Select(node) {
		for _, node := range ele.Attr {
			if node.Key == "href" {
				request := engine.Request{
					Url:        node.Val,
					ParserFunc: engine.NilParser,
				}

				result.Requests = append(result.Requests, request)
				result.Items = append(result.Items, node.Val)
			}
		}
	}

	return result
}
