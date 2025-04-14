package configuration

type Configuration struct {
	Application Application `json:"application"`
}

type Application struct {
	Server Server `json:"server"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
