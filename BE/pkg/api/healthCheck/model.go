package healthCheck

type GetAliveResponse struct {
	Api   string     `json:"api"`
	Version string  `json:"version"`
	Environment string  `json:"environment"`
}