package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/nakaaaa/nicovideo-api-go/pkg/search"
)

func main() {
	var encoder = schema.NewEncoder()
	q := url.Values{}
	req := &search.Query{
		Q:       "初音ミク",
		Targets: "title",
		Fields:  "contentId,title,viewCounter,userId",
		Sort:    "-viewCounter",
		Offset:  0,
		Limit:   3,
		Context: "apiguide_application",
	}
	if err := encoder.Encode(req, q); err != nil {
		panic(err)
	}

	r, err := http.NewRequest(http.MethodGet, search.Endpoint, nil)
	if err != nil {
		panic(err)
	}
	r.URL.RawQuery = q.Encode()

	r.Header.Add("User-Agent", "apiguide application")

	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	sd := new(search.Response)
	err = json.Unmarshal(body, sd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", sd)
}
