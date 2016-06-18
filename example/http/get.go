//ref https://golang.org/pkg/net/http/
package main
import "net/http"
func main() {
    resp, err := http.Get("http://example.com/")
}
