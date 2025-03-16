package interfaces

import (
	"io"
)

type Http interface {
	AddHeader(key string, value string)
	Get(url string) ([]byte, error)
	Post(url string, body io.Reader) ([]byte, error)
	Put(url string, body io.Reader) ([]byte, error)
	Delete(url string) ([]byte, error)
}
