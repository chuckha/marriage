package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/publicsuffix"
)

const unavailable = "unavailable"

func main() {
	start := time.Now().Add(2 * 24 * time.Hour).Format("2006-01-02")
	end := time.Now().Add(150 * 24 * time.Hour).Format("2006-01-02")
	fmt.Printf("Checking from %s to %s\n", start, end)
	urlFormat := "https://calendly.com/api/booking/event_types/BFEN3HYGXBSSUYOC/calendar/range?timezone=America%%2FNew_York&diagnostics=false&range_start=%s&range_end=%s&single_use_link_uuid=&embed_domain=projectcupid.cityofnewyork.us&embed_type=Inline"

	j, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(j)
	}
	client := http.Client{
		Jar: j,
	}
	u, err := url.Parse(fmt.Sprintf(urlFormat, start, end))
	if err != nil {
		panic(err)
	}
	fmt.Println(u.String())
	headers := http.Header{}
	headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	headers.Set("Host", "calendly.com")
	headers.Set("User-Agent", " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Safari/605.1.15")
	headers.Set("Accept-Language", "en-us")
	// Don't do this or it returns brotli lol
	//	headers.Set("Accept-Encoding", "gzip, deflate, br")
	req := &http.Request{
		Method: "GET",
		URL:    u,
		Header: headers,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	responseData := &response{}
	fmt.Fprintln(os.Stderr, string(data))
	if err := json.Unmarshal(data, responseData); err != nil {
		panic(err)
	}
	found := false
	for _, day := range responseData.Days {
		if day.Status != unavailable {
			found = true
			fmt.Printf("%s has availability! %v\n", day.Date, day.Spots)
		}
	}
	if !found {
		fmt.Println("Nothing available")
	}
}

type response struct {
	InviteePublisherError bool
	Today                 string
	AvailabilityTimezone  string
	Days                  []struct {
		Date          string
		Status        string
		Spots         []interface{}
		InviteeEvents []interface{}
	}
}
