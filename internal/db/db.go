package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Manager struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	client    *mongo.Client
}

func New() *Manager {
	m := &Manager{}
	m.ctx, m.ctxCancel = context.WithTimeout(context.Background(), 10*time.Second)

	var err error
	m.client, err = mongo.Connect(m.ctx, options.Client().ApplyURI("mongodb://localhost:21017"))
	if err != nil {
		log.Fatal(err)
	}

	return m
}

// Disconnect should be called before Manager release
func (m *Manager) Disconnect() {
	defer m.ctxCancel()
	err := m.client.Disconnect(m.ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Manager) AddUser() error {
	return nil
}
