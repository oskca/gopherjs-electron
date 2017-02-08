// +build ignore

package main

type Foo []struct {
	ConstructorMethod struct {
		Parameters []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Properties  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Parameters  []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"parameters"`
				PossibleValues []struct {
					Description string `json:"description"`
					Value       string `json:"value"`
				} `json:"possibleValues"`
				Properties []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Properties  []struct {
						Description string `json:"description"`
						Name        string `json:"name"`
						Type        string `json:"type"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"properties"`
			Required bool        `json:"required"`
			Type     interface{} `json:"type"`
		} `json:"parameters"`
		Signature string `json:"signature"`
	} `json:"constructorMethod"`
	Description string `json:"description"`
	Events      []struct {
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Platforms   []string `json:"platforms"`
		Returns     []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Parameters  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Type        string `json:"type"`
			} `json:"parameters"`
			Properties []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Type        string `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"returns"`
	} `json:"events"`
	InstanceEvents []struct {
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Platforms   []string `json:"platforms"`
		Returns     []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Parameters  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Type        string `json:"type"`
			} `json:"parameters"`
			PossibleValues []struct {
				Description string `json:"description"`
				Value       string `json:"value"`
			} `json:"possibleValues"`
			Properties []struct {
				Description    string `json:"description"`
				Name           string `json:"name"`
				PossibleValues []struct {
					Value string `json:"value"`
				} `json:"possibleValues"`
				Properties []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"returns"`
	} `json:"instanceEvents"`
	InstanceMethods []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Parameters  []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Parameters  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Parameters  []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Properties  []struct {
						Description string `json:"description"`
						Name        string `json:"name"`
						Type        string `json:"type"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"parameters"`
				Properties []struct {
					Description string        `json:"description"`
					Name        string        `json:"name"`
					Properties  []interface{} `json:"properties"`
					Type        string        `json:"type"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"parameters"`
			PossibleValues []struct {
				Description string `json:"description"`
				Value       string `json:"value"`
			} `json:"possibleValues"`
			Properties []struct {
				Description    string `json:"description"`
				Name           string `json:"name"`
				PossibleValues []struct {
					Description string `json:"description"`
					Value       string `json:"value"`
				} `json:"possibleValues"`
				Properties []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"properties"`
			Required bool   `json:"required"`
			Type     string `json:"type"`
		} `json:"parameters"`
		Platforms []string `json:"platforms"`
		Returns   struct {
			Description string `json:"description"`
			Properties  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Type        string `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"returns"`
		Signature string `json:"signature"`
	} `json:"instanceMethods"`
	InstanceName       string `json:"instanceName"`
	InstanceProperties []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"instanceProperties"`
	Methods []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Parameters  []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Parameters  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Parameters  []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Properties  []struct {
						Description string `json:"description"`
						Name        string `json:"name"`
						Properties  []struct {
							Description string `json:"description"`
							Name        string `json:"name"`
							Type        string `json:"type"`
						} `json:"properties"`
						Type string `json:"type"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"parameters"`
				Properties []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"parameters"`
			PossibleValues []struct {
				Description string `json:"description"`
				Value       string `json:"value"`
			} `json:"possibleValues"`
			Properties []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Parameters  []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"parameters"`
				Properties []interface{} `json:"properties"`
				Type       string        `json:"type"`
			} `json:"properties"`
			Required bool   `json:"required"`
			Type     string `json:"type"`
		} `json:"parameters"`
		Platforms []string `json:"platforms"`
		Returns   struct {
			Description string `json:"description"`
			Properties  []struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Type        string `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"returns"`
		Signature string `json:"signature"`
	} `json:"methods"`
	Name    string `json:"name"`
	Process struct {
		Main     bool `json:"main"`
		Renderer bool `json:"renderer"`
	} `json:"process"`
	Properties []struct {
		Description    string        `json:"description"`
		Name           string        `json:"name"`
		Parameters     []interface{} `json:"parameters"`
		PossibleValues []struct {
			Description string `json:"description"`
			Value       string `json:"value"`
		} `json:"possibleValues"`
		Properties []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Parameters  []struct {
				Description    string `json:"description"`
				Name           string `json:"name"`
				PossibleValues []struct {
					Value string `json:"value"`
				} `json:"possibleValues"`
				Required bool        `json:"required"`
				Type     interface{} `json:"type"`
			} `json:"parameters"`
			Platforms []string `json:"platforms"`
			Returns   struct {
				Description string `json:"description"`
				Type        string `json:"type"`
			} `json:"returns"`
			Signature string `json:"signature"`
			Type      string `json:"type"`
		} `json:"properties"`
		Type string `json:"type"`
	} `json:"properties"`
	RepoURL       string `json:"repoUrl"`
	Slug          string `json:"slug"`
	StaticMethods []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Parameters  []struct {
			Name     string `json:"name"`
			Required bool   `json:"required"`
			Type     string `json:"type"`
		} `json:"parameters"`
		Platforms []string `json:"platforms"`
		Returns   struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"returns"`
		Signature string `json:"signature"`
	} `json:"staticMethods"`
	Type       string `json:"type"`
	Version    string `json:"version"`
	WebsiteURL string `json:"websiteUrl"`
}
