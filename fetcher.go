package crawler

type Fetcher interface {
	Fetch(chan interface{})
}
