package ResponseModel

import "encoding/json"

type BaseResponseModel struct {
	Success    bool
	StatusCode int
	Data       interface{}
}

type BaseErrorModel struct {
	Success    bool
	StatusCode int
	Message    string
}

func (b *BaseErrorModel) JsonEncode() []byte {
	data, err := json.Marshal(b)
	if err != nil {
		return make([]byte, 0)
	}
	return data
}

func (b *BaseResponseModel) JsonEncode() []byte {
	data, err := json.Marshal(b)
	if err != nil {
		return make([]byte, 0)
	}
	return data
}
