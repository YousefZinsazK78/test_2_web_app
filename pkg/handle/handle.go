package handle

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/yousefzinsazk78/test_2_web_app/model/post"
	"github.com/yousefzinsazk78/test_2_web_app/model/user"
	filereaderr "github.com/yousefzinsazk78/test_2_web_app/pkg/file_readerr"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatalf("i think we have already some bugs %s", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("i think i can't execute this template in index handle func\n %s", err)
		return
	}
	// fmt.Fprintf(w, "hello world! %s", r.URL.Path[1:])
}

func HandleViewRoute(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/view/name/"):]
	res, err := filereaderr.ReadFile(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, res)
}

func HandleAddName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/add/name/"):]
	err := filereaderr.WriteData(name, user.User{Username: name, Password: "1234"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "added new name into files")
}

func HandleViewBlog(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/blog/"):]
	tmpl, err := template.ParseFiles("./static/view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Fprintln(w, title)
	tmpl.Execute(w, post.Post{Title: title, Description: "empty"})

}

func HandleSaveBlog(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	err := filereaderr.WriteDataBlog("blogPost", post.Post{Title: title, Description: description})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/blog/"+title, http.StatusAccepted)
	fmt.Println("file created")
}

func HandleNewBlog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/add.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
