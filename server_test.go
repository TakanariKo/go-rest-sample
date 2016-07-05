package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	main()
	os.Exit(m.Run())
}

func TestNotfoundHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(NotfoundHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode, ts.URL)
}

func TestTopHandler(t *testing.T) {
	// TODO assertion
	ts := httptest.NewServer(http.HandlerFunc(TopHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode, ts.URL)
	body := []byte("foo=bar")

	resp, err = http.Post(ts.URL, "text/html", bytes.NewReader(body))
	if err != nil {
		t.Error("Cannot send POST request")
	}
	t.Log(resp.StatusCode, ts.URL)

	request, err := http.NewRequest("PUT", ts.URL, bytes.NewReader(body))
	if err != nil {
		t.Error("Cannot create PUT request")
	}
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error("Cannot send PUT request")
	}
	t.Log(resp.StatusCode, ts.URL)

	request, err = http.NewRequest("DELETE", ts.URL, bytes.NewReader(body))
	if err != nil {
		t.Error("Cannot create DELETE request")
	}
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error("Cannot send DELETE request")
	}
	t.Log(resp.StatusCode, ts.URL)

}

func TestHandleGet(t *testing.T) {
	URL := url.URL{}
	URL.RawQuery = "param1=p1"
	HandleGet(&URL)
}

func TestHandlePost(t *testing.T) {
	r := strings.NewReader("foo=bar")
	readCloser := ioutil.NopCloser(r)

	request := http.Request{
		Method: "POST",
		Host:   "localhost",
		Body:   readCloser,
	}
	HandlePost(&request)
}

func TestHandlePut(t *testing.T) {
	HandlePut()
}

func TestHandleDelete(t *testing.T) {
	HandleDelete()
}

func TestInfoHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(InfoHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode, ts.URL)
}

func TestFileHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(FileHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode, ts.URL)
}

func TestRedirectHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(RedirectHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode, ts.URL)
}

func TestServeHTTP(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode)
	resp, err = http.Get("http://localhost:8000/redirect")
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode)
	resp, err = http.Get("http://localhost:8000/info.html")
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode)
	resp, err = http.Get("http://localhost:8000/favicon.ico")
	if err != nil {
		t.Error("Cannot send GET request")
	}
	t.Log(resp.StatusCode)
}
