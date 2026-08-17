package closekit

import (
	"errors"
	"io"
)

// CloseAll closes every closer and returns the combined error. It keeps
// closing even when one returns an error, so a failing resource does not
// leak the ones that follow. All non-nil errors are joined so callers can
// still inspect them with errors.Is / errors.As.
func CloseAll(closers ...io.Closer) error {
	var errs []error
	for _, c := range closers {
		if err := c.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
