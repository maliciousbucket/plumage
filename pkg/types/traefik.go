package types

type TraefikIngressConfig struct {
	Image        string
	Name         string
	Namespace    string
	AdminEnabled bool
	AdminPort    int
	Ports        []*TraefikPort
}

type TraefikPort struct {
	Protocol string
	Name     string
	Port     int
}
