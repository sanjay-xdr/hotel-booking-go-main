package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
    key   string
    value string
}

var theTests = []struct {
    name               string
    url                string
    method             string
    params             []postData
    expectedStatusCode int
}{
    {"home", "/", "GET", []postData{}, http.StatusOK},
    {"about", "/about", "GET", []postData{}, http.StatusOK},
    {"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
    {"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
    {"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
    {"contact", "/contact", "GET", []postData{}, http.StatusOK},
    {"make-res", "/make-reservation", "GET", []postData{}, http.StatusOK},
    {"post-search-avail", "/search-availability", "POST",[]postData{
        {key:"start", value: "2023-02-02"},
        {key:"end", value: "2023-02-02"},
    },http.StatusOK},
    {"post-search-avail-Json", "/search-availability-json", "POST",[]postData{
        {key:"start", value: "2023-02-02"},
        {key:"end", value: "2023-02-02"},
    },http.StatusOK},
    {"reservation post", "/make-reservation", "POST",[]postData{
        {key:"first_name", value: "test"},
        {key:"last_name", value: "test"},
        {key:"email", value: "test@gmail.com"},
        {key:"phone", value: "test@gmail.com"},

    },http.StatusOK},
    
}

func TestHandlers(t *testing.T) {
    routes := getRoutes()

    ts := httptest.NewTLSServer(routes)
    defer ts.Close()

    for _, e := range theTests {
        if e.method == "GET" {
            resp, err := ts.Client().Get(ts.URL + e.url)
            if err != nil {
                t.Log(err)
                t.Fatal(err)
            }

            if resp.StatusCode != e.expectedStatusCode {
                t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
            }
        } else {
                values:= url.Values{}
                for _,x:=range e.params {

                    values.Add(x.key,x.value)
                }

                resp,err:= ts.Client().PostForm(ts.URL + e.url, values)
                if err != nil {
                    t.Log(err)
                    t.Fatal(err)
                }
    
                if resp.StatusCode != e.expectedStatusCode {
                    t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
                }
        }
    }
}
