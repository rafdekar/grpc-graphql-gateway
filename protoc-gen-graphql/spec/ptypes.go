package spec

import (
	"fmt"
	"strings"

	"path/filepath"
)

var supportedPtypesLater = []string{
	"timestamppb",
	"wrapperspb",
	"emptypb",
}

func getImplementedPtypes(m *Message) (string, error) {
	ptype := strings.ToLower(filepath.Base(m.GoPackage()))

	var found bool
	for _, v := range supportedPtypesLater {
		if ptype == v {
			found = true
		}
	}
	if !found {
		return "", fmt.Errorf("google's ptype \"%s\" does not implement for now.", ptype)
	}

	return ptype, nil
}
