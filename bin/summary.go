package bin

import "github.com/dmin12/zgrab_tls1.3"

// Summary holds the results of a run of a ZGrab2 binary.
type Summary struct {
	StatusesPerModule map[string]*zgrab2.State `json:"statuses"`
	StartTime         string                   `json:"start"`
	EndTime           string                   `json:"end"`
	Duration          string                   `json:"duration"`
}
