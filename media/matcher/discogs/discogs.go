package discogs

var (
    api_url = "http://api.discogs.com"
    search_path = "/database/search" // api_url /database/search?q=<term1>%20<term2>
    releases_path = "/releases" // api_url /releases/<id>
)

// Search parameters
// http://www.discogs.com/developers/resources/database/search-endpoint.html
type DiscogsSearchParameters struct {
    Q string
    Type string
    Title string
    ReleaseTitle string
    Credit string
    Artist string
    Anv string
    Label string
    Genre string
    Style string
    Country string
    Year string
    Format string
    CatalogNumber string
    Barcode string
    Track string
    Submitter string
    Contributer string
}

type DiscogsSearch struct {
    Parameters DiscogsSearchParameters
}

func (d *DiscogsSearch) Search()  {
}



// Search discogs
// returns a list of release ids or nil and an error
// or nil and a nil if no valid match found
// options see http://www.discogs.com/developers/resources/database/search-endpoint.html
//         use any but 'q', which you pass in as keywords
func Search(keywords []string, options []string) ([]string, error) {
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


search({"animal", "collective", "feels",}, {"type=release",})


