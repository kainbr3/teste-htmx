package handlers

type HealthStatus struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type Result struct {
	Result string `json:"result"`
}

type OperationsPageInputs struct {
	Types       map[string]string `json:"types"`
	Domains     map[string]string `json:"domains"`
	Blockchains map[string]string `json:"blockchains"`
	Tokens      map[string]string `json:"tokens"`
}
