package main

type Fetcher interface {
	Fetch(chan interface{})
}
