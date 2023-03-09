package wc_analytics

import (
	"os"

	"github.com/segmentio/analytics-go/v3"
)

func GetAnalyticsClient() analytics.Client {
	var SegmentClient = analytics.New(os.Getenv("SEGMENT_API_KEY"))
	return SegmentClient
}
