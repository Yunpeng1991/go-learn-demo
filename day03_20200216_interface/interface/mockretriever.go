package _interface

type Retriever interface {
	Get(url string) string
}
