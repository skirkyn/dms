package sfa

import (
	"errors"
	"github.com/skirkyn/dcw/cmd/common"
	"github.com/skirkyn/dcw/cmd/common/dto"
)

type StringGeneratorHandler struct {
	workSupplier        common.Function[int, []string]
	responseTransformer dto.ResponseTransformer[[]string]
}

func NewGeneratorHandler(supplier common.Function[int, []string],
	responseTransformer dto.ResponseTransformer[[]string]) common.Function[dto.Request[any], []byte] {
	return &StringGeneratorHandler{supplier, responseTransformer}
}

func (gh *StringGeneratorHandler) Apply(req dto.Request[any]) ([]byte, error) {

	result, err := gh.workSupplier.Apply(req.Body.(int))
	resp := dto.Response[[]string]{Done: !errors.Is(err, PotentialResultsExhaustedError), Body: result}
	bytes, err := gh.responseTransformer.ResponseToBytes(resp)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}