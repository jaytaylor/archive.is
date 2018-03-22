package archiveis

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gigawattio/errorlib"
	"github.com/parnurzeal/gorequest"
)

const (
	baseURL   = "https://archive.is"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36"
)

var jsLocationExpr = regexp.MustCompile(`document\.location\.replace\(["']([^"']+)`)

// Capture archives the provided URL using the archive.is service.
func Capture(u string) (string, error) {
	submitID, err := newSubmitID()
	if err != nil {
		return "", err
	}

	// return id, nil

	content := fmt.Sprintf("submitid=%v&url=%v", url.QueryEscape(submitID), url.QueryEscape(u))
	fmt.Printf("content=%v\n", content)
	resp, body, errs := newRequest().Post(baseURL+"/submit/").Send(content).Set("content-type", "application/x-www-form-urlencoded").EndBytes()
	if err := errorlib.Merge(errs); err != nil {
		return "", err
	}
	if resp.StatusCode/100 != 2 {
		return "", fmt.Errorf("form submit received unhappy response status-code=%v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("constructing goquery doc from submission response: %s", err)
	}

	if script := doc.Find("script").First(); script != nil {
		js := strings.Trim(script.Text(), "\r\n\t ")
		if match := jsLocationExpr.FindStringSubmatch(js); len(match) > 1 {
			return match[1], nil
		}
	}

	fmt.Printf("body: %+v\n", string(body))
	fmt.Printf("headers: %+v\n", resp.Header)
	fmt.Printf("trailers: %+v\n", resp.Trailer)

	input := doc.Find("input[name=id]").First()
	if input == nil {
		return "", errors.New("page archive ID not found in submission response content")
	}
	id, exists := input.Attr("value")
	if !exists {
		return "", errors.New("no page archive ID value available")
	}

	final := fmt.Sprintf("%v/%v", baseURL, id)
	return final, nil
}

// newSubmitID gets the index page and extracts the form submission identifier.
func newSubmitID() (string, error) {
	resp, body, errs := newRequest().Get(baseURL).EndBytes()
	if err := errorlib.Merge(errs); err != nil {
		return "", err
	}
	if resp.StatusCode/100 != 2 {
		return "", fmt.Errorf("index retrieval received unhappy response status-code=%v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("constructing goquery doc from index: %s", err)
	}

	input := doc.Find("input[name=submitid]").First()
	if input == nil {
		return "", errors.New("no submitid element found")
	}
	id, exists := input.Attr("value")
	if !exists {
		return "", errors.New("no submitid value available")
	}
	return id, nil
}

func newRequest() *gorequest.SuperAgent {
	r := gorequest.New().
		Set("host", strings.Split(baseURL, "://")[1]).
		Set("user-agent", userAgent).
		Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8").
		Set("referer", baseURL+"/")
	return r
}
