package main

import (
	"errors"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(params httprouter.Params) (int64, error) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
