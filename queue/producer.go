package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

type IProducer interface {
	GetChannel() string

	Produce(m *Message) bool
}

type producer struct {
	name   string
	client redis.Conn
}

func (p *producer) GetChannel() string {
	return p.name
}

func (p *producer) Produce(m *Message) bool {
	prepareName := fmt.Sprintf("%s.prepare", p.GetChannel())
	fmt.Println("the prepareName is", prepareName)
	if mstring, err := jsoniter.Marshal(m); err == nil {
		if _, err := p.client.Do("LPUSH", prepareName, mstring); err == nil {
			return true
		}
	}
	return false
}

func NewProducer(name string, client redis.Conn) *producer {
	return &producer{
		name:   name,
		client: client,
	}
}
