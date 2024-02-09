package handlers

import (
	"fmt"
	"github.com/Gamilkarr/stattrack/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints_UpdateMetrics(t *testing.T) {
	tests := []struct {
		name, url  string
		statusCode int
	}{
		{
			name:       "200 gauge",
			url:        "/update/gauge/gauge_name/46.4",
			statusCode: http.StatusOK,
		},
		{
			name:       "404 gauge",
			url:        "/update/gauge/46.4",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "200 counter",
			url:        "/update/counter/counter_name/56",
			statusCode: http.StatusOK,
		},
		{
			name:       "404 counter",
			url:        "/update/counter/56",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "400 unknown metric type",
			url:        "/update/rnjfgkjwe/gauge_name/46.4",
			statusCode: http.StatusBadRequest,
		},
	}
	e := &Handler{
		Repo: &repository.MemStorage{
			Gauge:   make(map[string]float64),
			Counter: make(map[string]int64),
		},
	}
	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", e.UpdateMetrics)
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, fmt.Sprint(ts.URL, tt.url), nil)
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, tt.statusCode, resp.StatusCode)
			_ = resp.Body.Close()
		})
	}
}
