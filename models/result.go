package models

type Result struct {
	Url, Body string
	Urls []string
	Err error
	Depth int
}

type FakeResult struct {
	Body string
	Urls []string
}