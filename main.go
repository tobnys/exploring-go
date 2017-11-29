package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test") // Response
	// %T = TYPEOF
	//fmt.Printf("%T", w) // *http.response*http.responsee
	// %v = VALUE
	fmt.Printf("%v", w) /* &{0xc04204c820 0xc0420f8000 {} 0x50d490 true false false false 0xc04203e2c0 {0xc042108000 map[] false false} map[] false 4 -1 200 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0
	0 0 0 0] [0 0 0] 0xc042040070 0}&{0xc04204c820 0xc042118000 {} 0x50d490 true false false false 0xc04210e080 {0xc04212c000 map[] false false} map[] false 4 -1 200 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc04212a000 0}
	*/

	fmt.Printf("%v", r)
}
