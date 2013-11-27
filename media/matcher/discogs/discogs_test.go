
package matcher

import (
    "testing"
    "github.com/sinks/mr_organiser/media"
)

func TestFindTrack(t *testing.T) {
    matcher := DiscogsMatcher{}

    metadata := media.Metadata{Artist: "Animal Collective",
                               Album: "Feels",
                               Title: "Grass",
                              }
    matcher.FindTrack(metadata)
}
