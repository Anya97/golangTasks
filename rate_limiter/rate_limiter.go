package main

import (
	"fmt"
	"math"
	"time"
)

type TokenBucket struct {
	token          float64 // количество токенов в корзине
	capacity       float64
	refillRate     float64   // скорость пополнения токенов (2 токена/с)
	lastRefillTime time.Time // последнее пополнение корзины
}

func main() {
	tb := NewTokenBucket(10, 1)
	for i := 0; i < 2000; i++ {
		if result := tb.Request(1); result {
			fmt.Printf("Запрос номер %d будет выполнен\n", i+1)
		} else {
			fmt.Printf("Слишком большое количество запросов. Запрос номер %d не будет выполнен\n", i+1)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		token:          capacity,
		capacity:       capacity,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Refiller() {
	now := time.Now()
	timeLastAdditionBucket := now.Sub(tb.lastRefillTime)
	tokensToAdd := tb.refillRate * timeLastAdditionBucket.Seconds()
	tb.token = math.Min(tb.token+tokensToAdd, tb.capacity)
	tb.lastRefillTime = now
}

func (tb *TokenBucket) Request(tokens float64) bool {
	tb.Refiller()
	if tokens <= tb.token {
		tb.token -= tokens
		return true
	}
	return false
}
