package installer

const (
	statusIdle      = "idle"
	statusSetup     = "setup"
	statusCompleted = "completed"
	statusError     = "error"
)

type Status struct {
	Code      		string `json:"code"`
	ClientConfig    string `json:"client_config"`
	ErrorText 		string `json:"error_text"`
}

func (s *Status) setIdle() {
	s.Code = statusIdle
}

func (s *Status) setSetup() {
	s.Code = statusSetup
}

func (s *Status) setCompleted(clientConfig string) {
	s.Code = statusCompleted
	s.ClientConfig = clientConfig
}

func (s *Status) setError(text string) {
	s.ErrorText = text
	s.Code = statusError
}

func (s Status) IsCompleted() bool {
	return s.Code == statusCompleted
}