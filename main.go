package main

import (
	"log"
	"net/http"

	"github.com/yousefzinsazk78/test_2_web_app/pkg/handle"
)

func main() {
	http.HandleFunc("/", handle.HandleIndex)
	//name router
	http.HandleFunc("/add/name/", handle.HandleAddName)
	http.HandleFunc("/view/name/", handle.HandleViewRoute)
	//blog post router
	http.HandleFunc("/view/blog/", handle.HandleViewBlog)
	http.HandleFunc("/new/blog/", handle.HandleNewBlog)
	http.HandleFunc("/save/blog/", handle.HandleSaveBlog)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
