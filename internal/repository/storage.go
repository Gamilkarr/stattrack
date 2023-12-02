package repository

type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func (m *MemStorage) UpdateGaugeMetrics(name string, val float64) error {
	m.Gauge[name] = val
	return nil
}
func (m *MemStorage) UpdateCounterMetrics(name string, val int64) error {
	m.Counter[name] += val
	return nil
}
