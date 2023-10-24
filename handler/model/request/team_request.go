package request

type TeamRequest struct {
	Name       string `json:"name"`
	BadgePhoto string `json:"badgePhoto"`
	City       string `json:"city"`
	Coach      string `json:"coach"`
	Website    string `json:"website"`
}
