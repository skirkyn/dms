package sfv

import (
	"context"
	"errors"
	"github.com/skirkyn/dcw/cmd/common"
	"github.com/skirkyn/dcw/cmd/common/dto"
	"golang.org/x/sync/semaphore"
	"log"
)

type Worker struct {
	semaphore   *semaphore.Weighted
	context     context.Context
	transformer dto.ResponseTransformer[[]string]
	verifier    common.Predicate[string]
}

func NewWorker(semaphore *semaphore.Weighted,
	context context.Context,
	transformer dto.ResponseTransformer[[]string],
	verifier common.Predicate[string]) common.Function[[]byte, *dto.Request[string]] {
	return &Worker{semaphore, context, transformer, verifier}
}
func (w *Worker) Apply(work []byte) (*dto.Request[string], error) {
	resp, err := w.transformer.BytesToResponse(work)
	if err != nil {
		log.Printf("can't process response %s because of %s", string(work), err.Error())
		return nil, err
	}
	if !resp.Done {
		return nil, errors.New("done")
	}
	input := resp.Body

	for i := 0; i < len(input); i++ {
		current := input[i]
		if w.verifier.Test(current) {
			return &dto.Request[string]{Type: dto.Result, Body: current}, nil
		}
	}

	return nil, nil
}
