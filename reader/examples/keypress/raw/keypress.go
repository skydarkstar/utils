package main

import (
	"log"
	"sync"
	"time"

	"github.com/projectdiscovery/utils/reader"
	stringsutil "github.com/projectdiscovery/utils/strings"
)

func main() {
	stdr := reader.KeyPressReader{
		Timeout: time.Duration(5 * time.Second),
		Once:    &sync.Once{},
		Raw:     true,
	}

	stdr.Start()
	defer stdr.Stop()

	for {
		data := make([]byte, 1)
		n, err := stdr.Read(data)
		if stringsutil.IsPrintable(string(data)) {
			log.Println(n, err)
		}

		if stringsutil.IsCTRLC(string(data)) {
			break
		}
	}
}
