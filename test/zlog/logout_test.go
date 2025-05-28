package zlog

import (
	"GoChat/pkg/zlog"
	"testing"

	"go.uber.org/zap"
)

func TestInfo(t *testing.T) {
	zlog.Info("this is a info", zap.String("name", "apylee"))
}
