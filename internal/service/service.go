package service

import (
	"math/rand"
	"runtime"
)

type Service struct {
}

func (s *Service) GetGaugeMetricsMap() map[string]float64 {
	x := &runtime.MemStats{}
	runtime.ReadMemStats(x)

	metricsMap := map[string]float64{
		"Alloc":         float64(x.Alloc),
		"TotalAlloc":    float64(x.TotalAlloc),
		"Sys":           float64(x.Sys),
		"Lookups":       float64(x.Lookups),
		"Mallocs":       float64(x.Mallocs),
		"Frees":         float64(x.Frees),
		"HeapAlloc":     float64(x.HeapAlloc),
		"HeapSys":       float64(x.HeapSys),
		"HeapIdle":      float64(x.HeapIdle),
		"HeapInuse":     float64(x.HeapInuse),
		"HeapReleased":  float64(x.HeapReleased),
		"HeapObjects":   float64(x.HeapObjects),
		"StackInuse":    float64(x.StackInuse),
		"StackSys":      float64(x.StackSys),
		"MSpanInuse":    float64(x.MSpanInuse),
		"MSpanSys":      float64(x.MSpanSys),
		"MCacheInuse":   float64(x.MCacheInuse),
		"MCacheSys":     float64(x.MCacheSys),
		"BuckHashSys":   float64(x.BuckHashSys),
		"GCSys":         float64(x.GCSys),
		"OtherSys":      float64(x.OtherSys),
		"NextGC":        float64(x.NextGC),
		"LastGC":        float64(x.LastGC),
		"PauseTotalNs":  float64(x.PauseTotalNs),
		"NumGC":         float64(x.NumGC),
		"NumForcedGC":   float64(x.NumForcedGC),
		"GCCPUFraction": x.GCCPUFraction,
		"RandomValue":   rand.Float64(),
	}

	return metricsMap
}
