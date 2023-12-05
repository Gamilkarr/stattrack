package endpoints

import (
	"github.com/Gamilkarr/stattrack/internal/repository"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints_UpdateMetrics(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}

	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}

	tests := []struct {
		name    string
		request string
		fields  fields
		want    want
	}{
		{
			name:    "200 gauge",
			request: "http://localhost:8080/update/gauge/gauge_name/46.4",
			fields: fields{
				gauge: make(map[string]float64),
			},
			want: want{
				contentType: "text/plain",
				statusCode:  200,
			},
		},
		{
			name:    "404 gauge",
			request: "http://localhost:8080/update/gauge/46.4",
			fields: fields{
				gauge: make(map[string]float64),
			},
			want: want{
				contentType: "text/plain",
				statusCode:  404,
			},
		},
		{
			name:    "200 counter",
			request: "http://localhost:8080/update/counter/counter_name/56",
			fields: fields{
				counter: make(map[string]int64),
			},
			want: want{
				contentType: "text/plain",
				statusCode:  200,
			},
		},
		{
			name:    "404 counter",
			request: "http://localhost:8080/update/counter/56",
			fields: fields{
				counter: make(map[string]int64),
			},
			want: want{
				contentType: "text/plain",
				statusCode:  404,
			},
		},
		{
			name:    "400 unknown metric type",
			request: "http://localhost:8080/update/rnjfgkjwe/gauge_name/46.4",
			fields: fields{
				gauge: make(map[string]float64),
			},
			want: want{
				contentType: "text/plain",
				statusCode:  400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoints{
				Repo: &repository.MemStorage{
					Gauge:   tt.fields.gauge,
					Counter: tt.fields.counter,
				}}

			request := httptest.NewRequest(http.MethodPost, tt.request, nil)

			w := httptest.NewRecorder()
			e.UpdateMetrics(w, request)

			res := w.Result()
			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			res.Body.Close()
		})
	}
}
