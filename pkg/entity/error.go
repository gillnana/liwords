package entity

import (
	"fmt"
	"strings"

	"github.com/domino14/liwords/rpc/api/proto/realtime"
	"github.com/rs/zerolog/log"
)

type WooglesError struct {
	code realtime.WooglesError
	data []string
}

const WooglesErrorDelimiter = ";"

func NewWooglesError(code realtime.WooglesError, data ...string) *WooglesError {
	log.Debug().Interface("data", data).Int32("code", int32(code)).Msg("NewWooglesError")
	return &WooglesError{
		code: code,
		data: data,
	}
}

func (w *WooglesError) Error() string {
	var errb strings.Builder
	fmt.Fprintf(&errb, "%s%d", WooglesErrorDelimiter, w.code)
	for _, d := range w.data {
		fmt.Fprintf(&errb, "%s%s", WooglesErrorDelimiter, d)
	}
	return errb.String()
}
