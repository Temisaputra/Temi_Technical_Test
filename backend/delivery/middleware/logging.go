package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			requestID := uuid.New().String()
			r = r.WithContext(context.WithValue(r.Context(), "request_id", requestID))

			next.ServeHTTP(rw, r)

			latency := time.Since(start)

			fields := []zap.Field{
				zap.String("request_id", requestID),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status_code", rw.statusCode),
				zap.Duration("latency", latency),
				zap.String("client_ip", r.RemoteAddr),
			}

			switch {
			case rw.statusCode >= 500:
				logger.Error("server error", fields...)
			case rw.statusCode >= 400:
				logger.Warn("client error", fields...)
			default:
				logger.Info("request handled", fields...)
			}
		})
	}
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
