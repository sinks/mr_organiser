
package matcher

import (
    "github.com/sinks/mr_organiser/media"
)

type Matcher interface {
    FindTrack(media.Metadata) (media.Metadata, error)
}
