package webcrawler

type Fetcher interface {
	Fetch(chan interface{})
}
