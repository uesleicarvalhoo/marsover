package logger

type Params struct {
	Level          string `json:"level"`
	ServiceName    string `json:"serviceName"`
	ServiceVersion string `json:"serviceVersion"`
	Env            string `json:"env"`
}
