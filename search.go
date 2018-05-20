package archiveis

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

const snapshotLayout = "2 Jan 2006 15:04"

// Snapshot represents an instance of a URL page snapshot on archive.is.
type Snapshot struct {
	URL          string
	ThumbnailURL string
	Timestamp    time.Time
}

// Search for URL snapshots.
func Search(url string, timeout time.Duration) ([]Snapshot, error) {
	searchURL := fmt.Sprintf("%v/%v", BaseURL, url)
	resp, body, err := doRequest("", searchURL, nil, timeout)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("received non-2xx status code=%v while checking for archived snapshots of %v", resp.StatusCode, url)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("constructing goquery doc from archive.is search response: %s", err)
	}

	snaps := []Snapshot{}

	doc.Find(".THUMBS-BLOCK > div").Each(func(_ int, s *goquery.Selection) {
		var (
			u     = s.Find("a").AttrOr("href", "")
			th    = s.Find("img").AttrOr("src", "")
			tsStr = s.Find("div").Text()
		)

		if u == "" && strings.HasSuffix(tsStr, " more") {
			// Skip it, as it's a "# more" button.
			return
		}

		ts, err := time.Parse(snapshotLayout, tsStr)
		if err != nil {
			log.WithField("url", url).Errorf("Parsing snapshot timestamp %q: %s (skipping snapshot entry u=%s)", tsStr, err, u)
			return
		}

		snap := Snapshot{
			URL:          u,
			ThumbnailURL: th,
			Timestamp:    ts,
		}
		snaps = append(snaps, snap)
	})

	sort.Slice(snaps, func(i, j int) bool {
		return snaps[i].Timestamp.After(snaps[j].Timestamp)
	})

	return snaps, nil
}
