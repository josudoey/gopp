package main

import "net/url"

func main() {
	url, _ := url.Parse("http://example.com/search?q=foo")
	println(url.Query().Get("q"))
}
