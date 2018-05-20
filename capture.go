package archiveis

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

var (
	BaseURL               = "https://archive.is"                                                                                                       // Overrideable default package value.
	HTTPHost              = "archive.is"                                                                                                               // Overrideable default package value.
	UserAgent             = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36" // Overrideable default package value.
	DefaultRequestTimeout = 10 * time.Second                                                                                                           // Overrideable default package value.
	DefaultPollInterval   = 5 * time.Second                                                                                                            // Overrideable default package value.

	jsLocationExpr = regexp.MustCompile(`document\.location\.replace\(["']([^"']+)`)
)

// Config settings for page capture client behavior.
type Config struct {
	Wait           bool          // Wait until the crawl has been completed.
	WaitTimeout    time.Duration // Max time to wait for crawl completion.  Default is unlimited.
	PollInterval   time.Duration // Interval between crawl completion checks.  Defaults to 5s.
	RequestTimeout time.Duration // Overrides default request timeout.
	SubmitID       string        // Accepts a user-provided submitid.
}

// Capture archives the provided URL using the archive.is service.
func Capture(u string, cfg ...Config) (string, error) {
	timeout := DefaultRequestTimeout
	if len(cfg) > 0 && cfg[0].RequestTimeout > time.Duration(0) {
		timeout = cfg[0].RequestTimeout
	}

	var (
		submitID string
		body     []byte
		resp     *http.Response
		final    string
		err      error
	)

	if len(cfg) > 0 && len(cfg[0].SubmitID) > 0 {
		submitID = cfg[0].SubmitID
		log.Debugf("Will use caller-provided submitid=%v", submitID)
	} else if submitID, err = newSubmitID(timeout); err != nil {
		return "", err
	}

	content := fmt.Sprintf("submitid=%v&url=%v", url.QueryEscape(submitID), url.QueryEscape(u))

	resp, body, err = doRequest("POST", BaseURL+"/submit/", ioutil.NopCloser(bytes.NewBufferString(content)), timeout)
	if err != nil {
		return "", err
	}

	if resp.StatusCode/100 == 3 {
		// Page has already been archived.
		log.Debug("Detected redirect to archived page")

		if loc := resp.Header.Get("Location"); len(loc) == 0 {
			return "", fmt.Errorf("received a redirect status-code %v with an empty Location header", resp.StatusCode)
		} else {
			final = loc
		}
	} else {
		// log.Debugf("body: %+v\n", string(body))
		// log.Debugf("headers: %+v\n", resp.Header)
		// log.Debugf("trailers: %+v\n", resp.Trailer)

		doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
		if err != nil {
			return "", fmt.Errorf("constructing goquery doc from submission response: %s", err)
		}

		if script := doc.Find("script").First(); script != nil {
			js := strings.Trim(script.Text(), "\r\n\t ")
			if match := jsLocationExpr.FindStringSubmatch(js); len(match) > 1 {
				final = match[1]
			}
		}

		if len(final) == 0 {
			input := doc.Find("input[name=id]").First()
			if input == nil {
				return "", errors.New("page archive ID not found in submission response content")
			}
			id, exists := input.Attr("value")
			if !exists {
				log.Debugf("No page archive ID value detected, here was the page content: %v", string(body))
				return "", errors.New("no page archive ID value available")
			}

			final = fmt.Sprintf("%v/%v", BaseURL, id)
		}
	}

	log.Debugf("Capture for url=%v -> %v", u, final)

	if len(cfg) > 0 && cfg[0].Wait {
		var (
			waitTimeout  = cfg[0].WaitTimeout
			pollInterval = DefaultPollInterval
		)

		if cfg[0].PollInterval > time.Duration(0) {
			pollInterval = cfg[0].PollInterval
		}

		if err := waitForCrawlToFinish(final, body, timeout, waitTimeout, pollInterval); err != nil {
			return final, err
		}
	}

	return final, nil
}

// newSubmitID gets the index page and extracts the form submission identifier.
func newSubmitID(timeout time.Duration) (string, error) {
	_, body, err := doRequest("", BaseURL, nil, timeout)
	if err != nil {
		return "", err
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

func waitForCrawlToFinish(url string, body []byte, requestTimeout time.Duration, waitTimeout time.Duration, pollInterval time.Duration) error {
	var (
		expr  = regexp.MustCompile(`<html><body>`)
		until = time.Now().Add(waitTimeout)
		d     = time.Now().Sub(until)
		err   error
	)

	if body != nil && !expr.Match(body) {
		log.WithField("url", url).WithField("wait-timeout", waitTimeout).WithField("poll-interval", pollInterval).Debugf("Detected crawl completion after %s", d)
		if err := checkCrawlResult(body); err != nil {
			return err
		}
		return nil
	}

	log.WithField("url", url).WithField("wait-timeout", waitTimeout).WithField("poll-interval", pollInterval).Debug("Waiting for crawl to finish")
	for {
		if waitTimeout != time.Duration(0) && time.Now().After(until) {
			return fmt.Errorf("timed out after %s waiting for crawl to complete", waitTimeout)
		}

		_, body, err = doRequest("", url, nil, requestTimeout)

		d = time.Now().Sub(until)

		if err != nil {
			log.Debugf("Non-fatal error while polling for crawl completion: %s (continuing on, waiting for %s so far)", err, d)
		} else if !expr.Match(body) {
			log.WithField("url", url).WithField("wait-timeout", waitTimeout).WithField("poll-interval", pollInterval).Debugf("Detected crawl completion after %s", d)
			break
		}

		time.Sleep(pollInterval)
	}
	return nil
}

// checkCrawlResult searches for known archive.is errors in HTML content.
func checkCrawlResult(body []byte) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("crawl result check gq new doc: %s", err)
	}
	if block := doc.Find("html > body > div").First(); block != nil {
		if text := strings.Trim(block.Text(), "\r\n\t "); text == "Error: Network error." {
			return fmt.Errorf("archive.is crawl result: Network Error")
		}
	}
	return nil
}
