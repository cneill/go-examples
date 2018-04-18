package main

// Slightly modified from Rob Pike here (2010): https://web.stanford.edu/class/ee380/Abstracts/100428-pike-stanford.pdf
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s.", r.URL.Path[1:])
}
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}
