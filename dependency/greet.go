package dependency

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	return err
}
