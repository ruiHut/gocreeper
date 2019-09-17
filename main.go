package main

import (
	"crepper/zhengai/engine"
	"crepper/zhengai/parser"
)

const (
	REQUEST_URL = "https://www.zhenai.com/zhenghun"
)

func main() {
	engine.SimpleEngine.Run(engine.Request{
		Url:        REQUEST_URL,
		ParserFunc: parser.ParseCityList,
	})
}
