package constants

const (
	// Port defines HTTP port of the application. In case of change - fix the value in prometheus config.
	Port      = 8080
	TargetURL = "http://app:8080/payload"
)

// HistogramBuckets defines set of histogram buckets.
// 0.5ms, 1ms, 5ms, 10ms, 15ms, 20ms, 25ms.
var HistogramBuckets = []float64{.0005, .001, .005, .01, .015, .02, .025}
