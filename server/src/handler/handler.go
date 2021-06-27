package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"workout/src/helper"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	runHandler(w, r, h, handleError)
}

type errFn func(w http.ResponseWriter, r *http.Request, err *helper.HTTPError)

func runHandler(w http.ResponseWriter, r *http.Request, fn Handler, errfn errFn) {
	defer func() {
		if rv := recover(); rv != nil {
			err := errors.New("handler panic")
			helper.LogError(r, err, rv)
			errfn(w, r, helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UnrecognizedError))
		}
	}()

	r.Body = http.MaxBytesReader(w, r.Body, 2048)
	r.ParseForm()
	var buf helper.ResponseBuffer
	err := fn(&buf, r)
	if err == nil || reflect.ValueOf(err).IsNil() {
		buf.WriteTo(w)
	} else if e, ok := err.(*helper.HTTPError); ok {
		if e.StatusCode >= 500 {
			helper.LogError(r, err, nil)
		}
		errfn(w, r, e)
	} else {
		helper.LogError(r, err, nil)
		errfn(w, r, helper.CreateHTTPErrorWithCode(http.StatusInternalServerError, helper.UnrecognizedError))
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err *helper.HTTPError) {
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCodes": err.ErrorCodes,
			"error":      err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(err.StatusCode)
	io.WriteString(w, err.Error())
}
