package archiveis

import (
	"fmt"
	"testing"
	"time"
)

const page = "https://yro.slashdot.org/story/18/03/21/2112247/russia-secretly-helped-venezuela-launch-a-cryptocurrency-to-evade-us-sanctions#comments"

func TestCapture1(t *testing.T) {
	// Link which has been submitted before.
	url, err := Capture(page)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resolved URL=%q", url)
}

func TestCapture2(t *testing.T) {
	// Link which has likely not been submitted before.
	url, err := Capture(fmt.Sprintf("%v?%v", page, time.Now().Unix()))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Resolved URL=%q", url)
}
