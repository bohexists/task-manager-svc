package middleware

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
)

// NatsPublisher интерфейс для отправки логов в NATS
type NatsPublisher interface {
	Publish(subject string, data []byte) error
}

// LoggingInterceptor is a unified interceptor for both Unary and Streaming RPCs that logs requests, responses, and errors.
// It also sends logs to NATS via the provided NatsPublisher.
type LoggingInterceptor struct {
	natsPublisher NatsPublisher
}

// NewLoggingInterceptor creates a new instance of LoggingInterceptor with NATS integration
func NewLoggingInterceptor(natsPublisher NatsPublisher) *LoggingInterceptor {
	return &LoggingInterceptor{
		natsPublisher: natsPublisher,
	}
}

// UnaryInterceptor logs incoming and outgoing unary RPCs, and sends logs to NATS
func (l *LoggingInterceptor) UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Логируем входящий запрос
	log.Printf("Incoming Unary request - Method: %s, Time: %v, Request: %+v", info.FullMethod, start, req)

	// Обрабатываем запрос с помощью хендлера
	resp, err := handler(ctx, req)

	// Логируем ответ
	end := time.Now()
	log.Printf("Outgoing Unary response - Method: %s, Time: %v, Response: %+v, Error: %v", info.FullMethod, end, resp, err)

	// Формируем лог-сообщение для отправки в NATS
	logEntry := map[string]interface{}{
		"method":    info.FullMethod,
		"request":   req,
		"response":  resp,
		"error":     err,
		"timestamp": end,
	}
	logData, _ := json.Marshal(logEntry)

	// Отправляем лог в NATS
	l.natsPublisher.Publish("logs", logData)

	return resp, err
}

// StreamInterceptor logs streaming RPCs and sends logs to NATS
func (l *LoggingInterceptor) StreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()

	// Логируем начало стриминга
	log.Printf("Incoming Stream request - Method: %s, Time: %v", info.FullMethod, start)

	// Передаем обработку запроса хендлеру
	err := handler(srv, ss)

	// Логируем завершение стриминга
	end := time.Now()
	log.Printf("Completed Stream request - Method: %s, Time: %v, Error: %v", info.FullMethod, end, err)

	// Формируем лог-сообщение для отправки в NATS
	logEntry := map[string]interface{}{
		"method":    info.FullMethod,
		"error":     err,
		"timestamp": end,
	}
	logData, _ := json.Marshal(logEntry)

	// Отправляем лог в NATS
	l.natsPublisher.Publish("logs", logData)

	return err
}
