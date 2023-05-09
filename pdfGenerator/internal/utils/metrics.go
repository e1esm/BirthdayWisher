package utils

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var GrpcRequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "pdf_generation_duration",
	Help:    "Duration of pdf generation",
	Buckets: []float64{0.1, 1, 1.5, 2, 2.5, 3, 3.5}})
