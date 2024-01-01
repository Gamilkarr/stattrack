package repository

type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func NewRepo() (*MemStorage, error) {
	return &MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}, nil
}

func (m *MemStorage) UpdateGaugeMetrics(name string, val float64) error {
	m.Gauge[name] = val
	return nil
}
func (m *MemStorage) UpdateCounterMetrics(name string, val int64) error {
	m.Counter[name] += val
	return nil
}

func (m *MemStorage) GetGaugeMetricValue(name string) (float64, bool) {
	if val, ok := m.Gauge[name]; ok {
		return val, true
	}
	return 0, false
}
func (m *MemStorage) GetCounterMetricValue(name string) (int64, bool) {
	if val, ok := m.Counter[name]; ok {
		return val, true
	}
	return 0, false
}

func (m *MemStorage) GetCounterMetrics() map[string]int64 {
	return m.Counter
}

func (m *MemStorage) GetGaugeMetrics() map[string]float64 {
	return m.Gauge
}
