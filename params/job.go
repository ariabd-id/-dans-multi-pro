package params

type GetJob struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Fulltime    string `json:"username,omitempty"`
	Page        int    `json:"page,omitempty"`
}
