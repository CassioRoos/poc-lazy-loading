package internal

import v2 "gitlab.com/balance-inc/go-commons/log/v2"

func newLogger() v2.Logger {
	return v2.NewZapLogger()
}
