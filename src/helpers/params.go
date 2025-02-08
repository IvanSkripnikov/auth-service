package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"authenticator/logger"
)

const (
	ParamsGet  = 0
	ParamsPost = 1
)

type RequestParams struct {
	r      *http.Request
	method int
	err    error
}

func NewRequestParams(r *http.Request, method int) *RequestParams {
	return &RequestParams{
		r:      r,
		method: method,
		err:    nil,
	}
}

func (rp *RequestParams) setErr(err error) {
	if rp.err == nil {
		rp.err = err
	}
}

func (rp *RequestParams) Err() error {
	return rp.err
}

func (rp *RequestParams) GetValueFromRequestData(name string) (string, error) {
	var values url.Values

	if rp.method == ParamsGet {
		values = rp.r.URL.Query()
	} else {
		values = rp.r.PostForm
		logger.Debug(values.Encode())
	}

	if values.Has(name) {
		if values.Get(name) != "" {
			return values.Get(name), nil
		} else {
			return "", fmt.Errorf("'%s' parameter is required in this request", name)
		}
	} else {
		return "", fmt.Errorf("'%s' parameter is required in this request", name)
	}
}

func (rp *RequestParams) GetString(name string, required bool) string {
	valueString, err := rp.GetValueFromRequestData(name)

	if err != nil && required {
		rp.setErr(err)
		return ""
	}

	return valueString
}

func (rp *RequestParams) GetInt64(name string, required bool) int64 {
	valueString, err := rp.GetValueFromRequestData(name)

	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	value, err := strconv.Atoi(valueString)
	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	return int64(value)
}

func (rp *RequestParams) GetUint(name string, required bool) uint {
	valueString, err := rp.GetValueFromRequestData(name)

	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	value, err := strconv.Atoi(valueString)
	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	return uint(value)
}

func (rp *RequestParams) GetUint8(name string, required bool) uint8 {
	valueString, err := rp.GetValueFromRequestData(name)

	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	value, err := strconv.Atoi(valueString)
	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	return uint8(value)
}

func (rp *RequestParams) GetFloat32(name string, required bool) float32 {
	valueString, err := rp.GetValueFromRequestData(name)

	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	value, err := strconv.ParseFloat(valueString, 32)
	if err != nil && required {
		rp.setErr(err)
		return 0
	}

	return float32(value)
}
