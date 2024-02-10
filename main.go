package WeBlog

import "net/http"

func main() {
	http.ListenAndServe(":8080", handler: nil)
}
