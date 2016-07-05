package main

import (
	"flag"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// NotfoundHandler is an example of notfound.html with HTTP status code 404
func NotfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Not Found")
	glog.Info("Not Found:" + r.URL.String())
}

// TopHandler is an example how to respond to a request of various HTTP Method.
func TopHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infoln(
		r.RequestURI,
		r.Method,
		r.Header,
		r.Header.Get("User-Agent"),
		r.Body,
	)

	// Status Code
	w.WriteHeader(http.StatusOK)

	// Write response body
	io.WriteString(w, "Hello world! ")

	// Another example of write response body
	res := []byte("OK")
	w.Write(res)

	switch r.Method {
	case "GET":
		HandleGet(r.URL)
	case "POST":
		HandlePost(r)
	case "PUT":
		HandlePut()
	case "DELETE":
		HandleDelete()
	}
	return

}

// HandleGet is an example of HTTP GET request handling
func HandleGet(url *url.URL) {
	param1 := url.Query().Get("param1")
	glog.Info(param1)
}

// HandlePost is an example of HTTP POST request handling
func HandlePost(r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		glog.Error("Could not read body")
	}
	glog.Infoln(
		"POST",
		string(body),
	)
	r.Body.Close()
}

// HandlePut is an example of HTTP PUT request handling
func HandlePut() {
	glog.Info("PUT")
}

// HandleDelete is an example of HTTP DELETE request handling
func HandleDelete() {
	glog.Info("DELETE")
}

// InfoHandler is an example of handling info page
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Info page!")
}

// FileHandler is an example of opening a file in host's file system.
func FileHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infoln(
		r.URL.Path)
	w.WriteHeader(http.StatusOK)

	b, err := ioutil.ReadFile("." + r.URL.Path)
	if err != nil {
		glog.Error("Cannot read file " + r.URL.Path)
		return
	}

	w.Header().Set("Content-Type", "image/x-icon")
	if _, err := w.Write(b); err != nil {
		glog.Error("Cannot write response")
	}
}

// RedirectHandler is a redirect sample
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, RedirectURL, http.StatusSeeOther)
}

// RedirectURL is test redirect URL
const RedirectURL = "https://github.com"

var mux map[string]func(http.ResponseWriter, *http.Request)
var handler = &myHandler{}

func main() {
	flag.Parse()
	// handler := &myHandler{}
	server := http.Server{
		Addr:    ":8000",
		Handler: handler, // if we want to use a custom handler, set a custom handler
	}

	// add entrypoints via Default Handler
	// http.HandleFunc("/index.html", TopHandler)
	// another approach using custom handler
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = TopHandler
	mux["/redirect"] = RedirectHandler
	mux["/info.html"] = InfoHandler
	mux["/favicon.ico"] = FileHandler
	http.Handle("/", handler)
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.Path]; ok {
		h(w, r)
		return
	}
	// handle not implemented pages
	NotfoundHandler(w, r)
}
