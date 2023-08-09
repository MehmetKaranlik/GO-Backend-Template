package Middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type LoggerMiddleWare struct {
	Next http.Handler
}

func (m *LoggerMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Next == nil {
		m.Next = http.DefaultServeMux
	}

	logRequest(r)

	m.Next.ServeHTTP(w, r)

}

func NewLoggerMiddleWare(child http.Handler) *LoggerMiddleWare {
	return &LoggerMiddleWare{
		Next: child,
	}
}

func logRequest(r *http.Request) {

	fmt.Printf("\n\nUrl: %s\n", r.URL)
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Proto: %s\n", r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Printf("Header: { \n%s }\n", formattedHeaders(r.Header))
	if r.Header.Get("Content-Type") == "application/json" {
		logBody(r)
	}

}

func formattedHeaders(headers http.Header) string {
	var formattedHeaders string
	for key, value := range headers {
		formattedHeaders += fmt.Sprintf("%s: %s\n", key, value)
	}
	return formattedHeaders
}

func logBody(r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("Body : Cant be readed because of " + err.Error())
	}
	fmt.Printf("Body: %s\n", b)
	r.Body = io.NopCloser(bytes.NewBuffer(b))
}
