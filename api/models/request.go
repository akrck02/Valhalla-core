package models

type Request struct {
	Authorization string      `json:"authorization"`
	IP            string      `json:"ip"`
	UserAgent     string      `json:"userAgent"`
	Params        interface{} `json:"params"`
	Body          []byte      `json:"body"`
}

func (r *Request) GetParams() map[string]interface{} {
	return r.Params.(map[string]interface{})
}

func (r *Request) GetParam(key string) interface{} {
	return r.GetParams()[key]
}

func (r *Request) GetParamString(key string) string {
	return r.GetParam(key).(string)
}
