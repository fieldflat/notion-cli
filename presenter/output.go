package presenter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func StdOutput(byteArray []byte, pretty bool) {
	var buf bytes.Buffer
	if pretty {
		err := json.Indent(&buf, byteArray, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(os.Stdout, buf.String())
	} else {
		fmt.Fprintln(os.Stdout, string(byteArray))
	}
}
