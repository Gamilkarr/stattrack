package service

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Metrics struct {
	CounterMetrics map[string]int64
	GaugeMetrics   map[string]float64
}

func (m *Metrics) UpdateGaugeMetrics() {
	x := &runtime.MemStats{}
	runtime.ReadMemStats(x)

	m.GaugeMetrics = map[string]float64{
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
}

func (m *Metrics) UpdateCounterMetrics() {
	if m.CounterMetrics == nil {
		m.CounterMetrics = map[string]int64{"PollCount": 0}
	}
	m.CounterMetrics["PollCount"] += 1
}

func (m *Metrics) UpdateMetrics() {
	time.Sleep(2 * time.Second)
	m.UpdateGaugeMetrics()
	m.UpdateCounterMetrics()
}

func (m *Metrics) GetAllMetrics() []map[string]string {
	result := make([]map[string]string, 0, len(m.CounterMetrics)+len(m.GaugeMetrics))
	for name, value := range m.CounterMetrics {
		result = append(result, map[string]string{"type": "counter", "name": name, "value": fmt.Sprintf("%d", value)})
	}
	for name, value := range m.GaugeMetrics {
		result = append(result, map[string]string{"type": "gauge", "name": name, "value": fmt.Sprintf("%g", value)})
	}
	return result
}
