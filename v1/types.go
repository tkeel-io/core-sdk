package v1

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type Result struct {
	Err  error
	Data []byte
}

func (r *Result) Parse(v interface{}) error {
	if nil != r.Err {
		return r.Err
	}

	var err error
	var val interface{}
	if err = json.Unmarshal(r.Data, &val); nil != err {
		return err
	} else if err = mapstructure.Decode(val, v); nil != err {
		return err
	}

	return nil
}
