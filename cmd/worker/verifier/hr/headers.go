package hr

import (
	"github.com/skirkyn/dcw/cmd/common"
)

type SimpleHeadersSupplier[In any] struct {
	defaultHeaders map[string]string
}

func NewSimpleHeadersSupplier[In any](defaultHeaders map[string]string) common.Function[In, map[string]string] {
	return &SimpleHeadersSupplier[In]{defaultHeaders: defaultHeaders}
}

func (sf *SimpleHeadersSupplier[In]) Apply(in In) (map[string]string, error) {
	return sf.defaultHeaders, nil
}