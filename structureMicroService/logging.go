package main

import (
	"context"
	"fmt"
	"time"
)

// Custom logger for service, or we can just use log
// This can be useful for benchmarking functions

type LoggingService struct {
	next Service
}

func newLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) getCatFact(ctx context.Context) (fact *CatFact, err error){
    defer func(start time.Time){
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.getCatFact(ctx)
}
