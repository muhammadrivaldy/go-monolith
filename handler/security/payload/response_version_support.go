package payload

type ResponseVersionSupport struct {
	Version string `json:"version"`
	Support bool   `json:"support"`
}
