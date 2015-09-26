package image

import (
	"fmt"
	"strings"

	derr "github.com/docker/docker/errors"
	"github.com/docker/docker/reference"
)

// ErrImageDoesNotExist is error returned when no image can be found for a reference.
type ErrImageDoesNotExist struct {
	RefOrID string
}

func (e ErrImageDoesNotExist) Error() string {
	return fmt.Sprintf("Error: No such image: %s", e.RefOrID)
}

// IsErrImageDoesNotExist returns true if the error is caused
// when an image is not found in the docker host.
func IsErrImageDoesNotExist(err error) bool {
	_, ok := err.(ErrImageDoesNotExist)
	return ok
}

func ImageDoesNotExistToErrcode(err error) error {
	if dne, isDNE := err.(ErrImageDoesNotExist); isDNE {
		if strings.Contains(dne.RefOrID, "@") {
			return derr.ErrorCodeNoSuchImageHash.WithArgs(dne.RefOrID)
		}
		tag := reference.DefaultTag
		ref, err := reference.ParseNamed(dne.RefOrID)
		if err != nil {
			return derr.ErrorCodeNoSuchImageTag.WithArgs(dne.RefOrID, tag)
		}
		if tagged, isTagged := ref.(reference.NamedTagged); isTagged {
			tag = tagged.Tag()
		}
		return derr.ErrorCodeNoSuchImageTag.WithArgs(ref.Name(), tag)
	}
	return err
}
