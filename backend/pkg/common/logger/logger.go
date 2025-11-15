package logger

import (
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(cfg *config.Config) Logger {
	z, _ := zap.NewProduction()
	s := z.Sugar()
	return Logger{s}
}
