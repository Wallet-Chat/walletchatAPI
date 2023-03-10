package wc_analytics

import (
	"os"

	"github.com/segmentio/analytics-go/v3"
)

//to send the data, .Close has to be called on the instance of SegmentClient
func GetAnalyticsClient() analytics.Client {
	var SegmentClient = analytics.New(os.Getenv("SEGMENT_API_KEY"))
	return SegmentClient
}
