package utils

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var GrpcRequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "grpc_request_duration",
	Help:    "Duration of grpc requests in seconds",
	Buckets: []float64{0.005, 0.01, 0.1, 1}})
