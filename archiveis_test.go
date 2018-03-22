package archiveis

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const page = "https://yro.slashdot.org/story/18/03/21/2112247/russia-secretly-helped-venezuela-launch-a-cryptocurrency-to-evade-us-sanctions#comments"

func TestCapture1(t *testing.T) {
	randSleep(t, 60)

	// Link which has been submitted before.
	url, err := Capture(page)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resolved URL=%q", url)
}

func TestCapture2(t *testing.T) {
	randSleep(t, 60)

	// Link which has likely not been submitted before.
	url, err := Capture(fmt.Sprintf("%v?%v", page, time.Now().Unix()))
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
