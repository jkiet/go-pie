package comet

import (
	"encoding/json"
	"fmt"
	"github.com/jkiet/go-pie/driver"
	"io/ioutil"
	"net/http"
)

type CometClient struct {
	Section       *driver.Section
	etag          string
	last_modified string
}

func NewCometClient(s *driver.Section) *CometClient {
	return &CometClient{Section: s, etag: "0", last_modified: "Thu, 1 Jan 1970 00:00:00 GMT"}
}

func (c *CometClient) Listen(url string) {
	client := &http.Client{}
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("\nmalformed url %v, err: %v", url, err)
			return
		}
		req.Header.Add("If-None-Match", c.etag)
		req.Header.Add("If-Modified-Since", c.last_modified)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("\nrequest to %v failed - err: %v", url, err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("\nwrong response status code from %v: %v", url, resp.StatusCode)
			continue
		}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			fmt.Printf("\nread body error: %v", err)
			continue
		}
		resp.Body.Close()

		var data map[string]string
		err = json.Unmarshal(respBody, &data)
		if err != nil {
			fmt.Printf("\nunmarshal error: %v", err)
			continue
		}
		c.Section.Reload(data)
		c.last_modified = resp.Header.Get("Last-Modified")
		c.etag = resp.Header.Get("Etag")
	}
}
