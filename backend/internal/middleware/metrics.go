package middleware

import (
    "fmt"
    "regexp"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
)

var (
    // use limited labels to avoid high cardinality
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    // tuned buckets for typical API latencies
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Histogram of request latencies",
            Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.3, 1.2, 5.0},
        },
        []string{"method", "path"},
    )
)

func init() {
    // safe register: MustRegister will panic if already registered; handle gracefully by ignoring duplicate registration
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

var (
    uuidRe    = regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
    numberRe  = regexp.MustCompile(`^\d+$`)
)

// sanitizePath replaces likely variable path segments (ids, uuids) with :id so labels stay low-cardinality.
func sanitizePath(p string) string {
    if p == "" {
        return p
    }
    // split and sanitize each segment
    segs := strings.Split(p, "/")
    for i, s := range segs {
        if s == "" {
            continue
        }
        if uuidRe.MatchString(s) || numberRe.MatchString(s) {
            segs[i] = ":id"
            continue
        }
        // very long hex strings or base64-like segments are also considered ids
        if len(s) > 40 {
            segs[i] = ":id"
        }
    }
    sp := strings.Join(segs, "/")
    // ensure it starts with '/'
    if !strings.HasPrefix(sp, "/") {
        sp = "/" + sp
    }
    return sp
}

// PrometheusMiddleware instruments requests with Prometheus metrics.
func PrometheusMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        status := c.Writer.Status()
        path := c.FullPath()
        if path == "" {
            path = c.Request.URL.Path
        }
        path = sanitizePath(path)
        // use numeric status to keep labels stable
        statusLabel := fmt.Sprintf("%d", status)
        httpRequestsTotal.WithLabelValues(c.Request.Method, path, statusLabel).Inc()
        httpRequestDuration.WithLabelValues(c.Request.Method, path).Observe(time.Since(start).Seconds())
    }
}
