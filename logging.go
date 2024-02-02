package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) *LoggingService {
	return &LoggingService{
		next: next,
	}
}

func (l *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {

	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{

			"took": time.Since(start),
			"err":  err,
			"fact": fact.Fact,
		}).Info("catfact")

	}(time.Now())

	return l.next.GetCatFact(ctx)
}
