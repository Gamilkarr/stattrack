package repository

//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestMemStorage_UpdateGaugeMetrics(t *testing.T) {
//
//	type storage struct {
//		Gauge   map[string]float64
//		Counter map[string]int64
//	}
//	type args struct {
//		name string
//		val  float64
//	}
//	tests := []struct {
//		name    string
//		storage storage
//		args    args
//		want    map[string]float64
//		wantErr bool
//	}{
//		{
//			name: "base record",
//			storage: storage{
//				Gauge: map[string]float64{"firstRecord": 7.4},
//			},
//			args: args{
//				name: "metrics",
//				val:  1.6,
//			},
//			want:    map[string]float64{"firstRecord": 7.4, "metrics": 1.6},
//			wantErr: false,
//		},
//		{
//			name: "new record",
//			storage: storage{
//				Gauge: map[string]float64{"metrics": 7.5},
//			},
//			args: args{
//				name: "metrics",
//				val:  1.6,
//			},
//			want:    map[string]float64{"metrics": 1.6},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &MemStorage{
//				Gauge:   tt.storage.Gauge,
//				Counter: tt.storage.Counter,
//			}
//			if err := m.UpdateGaugeMetrics(tt.args.name, tt.args.val); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateGaugeMetrics() error = %v, wantErr %v", err, tt.wantErr)
//			}
//			assert.Equal(t, tt.want, m.Gauge)
//		})
//	}
//}
//
//func TestMemStorage_UpdateCounterMetrics(t *testing.T) {
//	type storage struct {
//		Gauge   map[string]float64
//		Counter map[string]int64
//	}
//	type args struct {
//		name string
//		val  int64
//	}
//	tests := []struct {
//		name    string
//		storage storage
//		args    args
//		want    map[string]int64
//		wantErr bool
//	}{
//		{
//			name: "base record",
//			storage: storage{
//				Counter: map[string]int64{"firstRecord": 7},
//			},
//			args: args{
//				name: "metrics",
//				val:  1,
//			},
//			want:    map[string]int64{"firstRecord": 7, "metrics": 1},
//			wantErr: false,
//		},
//		{
//			name: "new record",
//			storage: storage{
//				Counter: map[string]int64{"metrics": 7},
//			},
//			args: args{
//				name: "metrics",
//				val:  1,
//			},
//			want:    map[string]int64{"metrics": 8},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &MemStorage{
//				Gauge:   tt.storage.Gauge,
//				Counter: tt.storage.Counter,
//			}
//			if err := m.UpdateCounterMetrics(tt.args.name, tt.args.val); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateCounterMetrics() error = %v, wantErr %v", err, tt.wantErr)
//			}
//			assert.Equal(t, tt.want, m.Counter)
//		})
//	}
//}
