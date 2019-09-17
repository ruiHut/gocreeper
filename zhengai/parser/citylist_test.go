package parser

import (
	"crepper/fetcher"
	"testing"
)

const URL = "https://www.zhenai.com/zhenghun" // TODO 待持久化

func TestParseCityList(t *testing.T) {
	contexts, err := fetcher.Fetch(URL)
	if err != nil {
		panic(err)
	}
	res := ParseCityList(contexts)

	const resultSize = 470
	if len(res.Requests) != resultSize {
		t.Errorf("result should hava %d, but had  %d", resultSize, len(res.Requests))
	}
}
