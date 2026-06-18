package configs

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type GormLogger struct {
	log                  *logrus.Logger
	level                logger.LogLevel
	slowThreshold        time.Duration
	ignoreRecordNotFound bool
}

func NewGormLogger(log *logrus.Logger, level logger.LogLevel) logger.Interface {
	return &GormLogger{
		log:                  log,
		level:                level,
		slowThreshold:        200 * time.Millisecond,
		ignoreRecordNotFound: true,
	}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	clone := *l
	clone.level = level
	return &clone
}

func (l *GormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if l.level >= logger.Info {
		l.log.WithContext(ctx).Infof(msg, args...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if l.level >= logger.Warn {
		l.log.WithContext(ctx).Warnf(msg, args...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if l.level >= logger.Error {
		l.log.WithContext(ctx).Errorf(msg, args...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	entry := l.log.WithContext(ctx).WithFields(logrus.Fields{
		"component":  "gorm",
		"elapsed":    elapsed.String(),
		"elapsed_ms": float64(elapsed.Microseconds()) / 1000,
		"rows":       rows,
		"sql":        sql,
	})

	switch {
	case err != nil && l.level >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.ignoreRecordNotFound):
		entry.WithError(err).Error("database query failed")
	case elapsed > l.slowThreshold && l.level >= logger.Warn:
		entry.WithField("slow_threshold", l.slowThreshold.String()).Warn("slow database query")
	case l.level >= logger.Info:
		entry.Debug("database query")
	}
}
