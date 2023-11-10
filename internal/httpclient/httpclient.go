package httpclient

import "fmt"

func Get(name string) string {
	return fmt.Sprintf("GET %v", name)
}
