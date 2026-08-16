package closekit

import "io"

func CloseAll(closers ...io.Closer) error {
	for _, c := range closers {
		if err := c.Close(); err != nil {
			return err // BUG: stop early
		}
	}
	return nil
}
