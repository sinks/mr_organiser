package discogs

import (
    "net/http"
    "net/url"
    "reflect"
    "fmt"
    "io/ioutil"
    "encoding/json"
)


var (
    searchPath = "/database/search" // api_url /database/search?q=<term1>%20<term2>
)

// Search parameters
// http://www.discogs.com/developers/resources/database/search-endpoint.html
type SearchParameters struct {
    Q               string      `param:"q"`
    Type            string      `param:"type"`
    Title           string      `param:"title"`
    ReleaseTitle    string      `param:"release_title"`
    Credit          string      `param:"credit"`
    Artist          string      `param:"artist"`
    Anv             string      `param:"anv"`
    Label           string      `param:"label"`
    Genre           string      `param:"genre"`
    Style           string      `param:"style"`
    Country         string      `param:"country"`
    Year            string      `param:"year"`
    Format          string      `param:"format"`
    CatalogNumber   string      `param:"catno"`
    Barcode         string      `param:"barcode"`
    Track           string      `param:"track"`
    Submitter       string      `param:"submitter"`
    Contributer     string      `param:"contributer"`
}

type DiscogsSearch struct {
    Parameters      SearchParameters
}

func (d *DiscogsSearch) Search()  {
    // set the host and path
    searchUrl, err := url.Parse(apiUrl + searchPath)
    if err != nil {
    }

    // get the query parameters
    params := url.Values{}

    t := reflect.TypeOf(&d.Parameters).Elem()
    v := reflect.ValueOf(&d.Parameters).Elem()
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        tag := field.Tag.Get("param")
        value := url.QueryEscape(v.Field(i).String())
        if value != "" {
            params.Add(tag, value)
        }
    }
    searchUrl.RawQuery = params.Encode()

    fmt.Println(searchUrl.String())
    // issue the get request
    resp, err := http.Get(searchUrl.String())
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var f interface{}
    err = json.Unmarshal(body, &f)
}

