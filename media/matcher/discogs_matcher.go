
package matcher

import (
    "github.com/sinks/mr_organiser/media"
    "github.com/sinks/go_discogs"
    "net/http"
    "reflect"
    "fmt"
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


