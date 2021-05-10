package internal

import "go.uber.org/fx"

var Module = fx.Provide(
	newLogger,
)
