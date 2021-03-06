package log

import (
	"os"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options declare log's configuration
type Options struct {
	Level    string
	Encoding string
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}

	return o, nil
}

// New serve caller to create zap.Logger
func New(o *Options) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte(o.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	enc := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	if o.Encoding == "json" {
		enc = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	cores := make([]zapcore.Core, 0, 2)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	return logger, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New, NewOptions)
