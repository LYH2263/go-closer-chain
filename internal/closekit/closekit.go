package closekit

import "io"

func CloseAll(closers ...io.Closer) error {
	var first error
	for _, c := range closers {
		if err := c.Close(); err != nil && first == nil {
			first = err
		}
	}
	return first
}
