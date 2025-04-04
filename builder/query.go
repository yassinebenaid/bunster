package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type query struct {
	module string
	commit string
}

var queryRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})(?:/[a-zA-Z0-9._-]+)+(?:@[0-9a-f]{40})$`)

func parseQuery(v string) (query, error) {
	var q query

	if !queryRegex.MatchString(v) {
		return q, fmt.Errorf("module path %q is not in an expected format", v)
	}

	vslice := strings.SplitN(v, "@", 2)
	q.module, q.commit = vslice[0], vslice[1]

	return q, nil
}
