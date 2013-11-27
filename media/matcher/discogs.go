
package matcher

import (
    "github.com/sinks/mr_organiser/media"
    "net/http"
    "reflect"
    "fmt"
)

var (
    api_url = "http://api.discogs.com"
    search_path = "/database/search" // api_url /database/search?q=<term1>%20<term2>
    releases_path = "/releases" // api_url /releases/<id>
)

// Discogs Matcher
// http://www.discogs.com/developers/index.html
type DiscogsMatcher struct {
}

func (d *DiscogsMatcher) FindTrack(metadata media.Metadata) (media.Metadata, error) {
    // add all values to query

    fmt.Println(http.Get(api_url + "/database/search?q=" + allQ))

    return metadata, nil
}

// find the most likely releases
// returns a list of release ids or nil and an error
// or nil and a nil if no valid match found
func search(metadata media.Metadata) ([]string, error) {
    url, err := url.Parse(api_url + search_path)

    if err != nil {
        return nil, err
    }

    // add values from metadata to query
    s := reflect.ValueOf(&metadata).Elem()
    allQ := ""
    for i := 0; i < s.NumField(); i++ {
        field := s.Field(i)
    }
    url.Query().Add("q", allQ)
}


