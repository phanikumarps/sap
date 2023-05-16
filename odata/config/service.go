package config

type Host struct {
	address string
	port    string
}

func NewHost(host, port string) *Host {

	return &Host{
		address: host,
		port:    port,
	}
}

type RootPath string

type Resource struct {
	name      string
	host      Host
	rootPath  string
	sapClient string
	authToken string
	csrfToken string
}

func NewResource(r Resource) *Resource {
	return &Resource{
		name:      r.name,
		host:      r.host,
		rootPath:  r.rootPath,
		sapClient: r.sapClient,
		authToken: r.authToken,
		csrfToken: r.csrfToken,
	}
}

func DefaultRootPath(service string) *RootPath {
	r := RootPath("/sap/opu/odata/sap/" + service + "/")
	return &r
}
