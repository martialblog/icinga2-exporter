package icinga

type Perfdata struct {
	IsCounter bool    `json:"counter"`
	Label     string  `json:"label"`
	Value     float64 `json:"value"`
}

type APIResult struct {
	Results []struct {
		Name     string     `json:"name"`
		Perfdata []Perfdata `json:"perfdata,omitempty"`
	} `json:"results"`
}

type CIBResult struct {
	Results []struct {
		Name   string             `json:"name"`
		Status map[string]float64 `json:"status,omitempty"`
	} `json:"results"`
}
