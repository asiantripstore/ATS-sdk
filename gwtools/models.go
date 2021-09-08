package gwtools

//RouteInfo information of a singlie route
type RouteInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Name   string `json:"name"`
}

//ListRoutes ...
type ListRoutes struct {
	AppName    string      `json:"app_name"`
	Host       string      `json:"host"`
	Port       string      `json:"port"`
	ListRoutes []RouteInfo `json:"listRoutes"`
}
