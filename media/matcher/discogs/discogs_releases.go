package discogs

type Rating struct {
    Count                   int
    Average                 float32
}

type Community struct {
    Status                  string
    Have                    string
    Want                    string
    Rating                  Rating
    Contributors            []User
    Submitter               User
}

type User struct {
    Username                string
    ResourceUrl             string
}

type Image struct {
    Uri                     string
    Height                  int
    Width                   int
    ResourceUrl             string
    Type                    string
    Uri150                  string
}

type Track struct {
    Duration                string
    Position                string
    Title                   string
}

type Artist struct {
    Id                      int
    Join                    string
    Name                    string
    Anv                     string
    Tracks                  string
    Role                    string
    ResourceUrl             string
}

type Video struct {
    Duration                int
    Description             string
    Embed                   bool
    Uri                     string
    Title                   string
}

type Release struct {
    Id                      int
    Title                   string
    ResourceUrl             string
    Uri                     string
    Status                  string
    DataQuality             string
    MasterId                string
    MasterUrl               string
    Country                 string
    Year                    string
    Released                string
    ReleasedFormatted       string
    Notes                   []string
    Styles                  []string
    Genres                  []string
    EstimatedWeight         int
    FormatQuantity          int
    Community               Community
    Sumbitter               User
    Contributors            []User
    Labels                  []Label
    Companies               []Company
    ExtraArtists            []????
    Videos                  []Video
    Artists                 []Artist
    Formats             []???
    Images              []
    Identifiers         []
    Tracklist               []Track
}


