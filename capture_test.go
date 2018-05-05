// +build integration

package archiveis

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	page     = "https://github.com/jaytaylor/archive.is"
	maxSleep = 60
)

func TestCapture1(t *testing.T) {
	randSleep(t, maxSleep)

	// Link which has been submitted before.
	url, err := Capture(page)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resolved URL=%q", url)
}

func TestCapture2(t *testing.T) {
	randSleep(t, maxSleep)

	// Link which has likely not been submitted before.
	url, err := Capture(fmt.Sprintf("%v?%v", page, time.Now().Unix()))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resolved URL=%q", url)
}

func TestCapture3(t *testing.T) {
	randSleep(t, maxSleep)

	cfg := Config{
		Wait:           true,
		WaitTimeout:    180 * time.Second,
		PollInterval:   15 * time.Second,
		RequestTimeout: 15 * time.Second,
	}

	// Link which has likely not been submitted before.
	url, err := Capture(fmt.Sprintf("%v?%v", page, time.Now().Unix()), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resolved URL=%q", url)
}

// randSleep helps smooth out travis running so many builds at once.
func randSleep(t *testing.T, max int) {
	n := rand.Int() % max
	d := time.Duration(n) * time.Second
	t.Logf("Sleeping for randomly selected interval: %s", d)
	time.Sleep(d)
}
