package pit

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
	"time"
	"io/ioutil"
)

func PitFour() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		var fixedUrl = url
		if !strings.HasPrefix(url, "http://") {
			fixedUrl = "http://" + url
		}
		resp, err := http.Get(fixedUrl)
		if err == nil {
			// dataLen, err := io.Copy(os.Stdout, resp.Body)
			dataLen, err := io.Copy(ioutil.Discard, resp.Body)
			// respData, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err == nil {
				fmt.Printf("Succesfully wrote %d data for url %s\n", dataLen, url)
			}
		}
	}
	fmt.Printf("Total time taken is %.2fs\n", time.Since(start).Seconds())
}
