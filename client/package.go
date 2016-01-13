package client

type PackageService struct {
	client *Client
}

type Package struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	Annotations `json:"annotations"`
	Parameters  `json:"parameters"`
	Binding     `json:"binding"`
}

type Binding struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type PackageListOptions struct {
	Public bool   `url:"public,omitempty"`
	Limit  string `url:"limit,omitempty"`
	Skip   int    `url:"skip,omitempty"`
	Since  int    `url:"since,omitempty"`
	Docs   bool   `url:"docs,omitempty"`
}
