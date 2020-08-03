package models

type AccountIdDetail struct {
	ID    int64  `json:"id,omitempty"`
	Email string `json:"email,omitempty"`

	Dashboard                 bool `json:"dashboard"`
	Spc                       bool `json:"spc"`
	AdjustmentGuiding         bool `json:"adjustment_guiding"`
	ParameterFailurePediction bool `json:"parameter_failure_pediction"`
	ProcessStatusAnalysis     bool `json:"process_status_analysis"`

	Company string `json:"company,omitempty"`
}

type Err struct {
	Error_ string `json:"error,omitempty"`
}
