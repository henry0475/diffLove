package config

// Config for auto config
type Config struct {
	Hosts struct {
		Database struct {
			Address  string `json:"address"`
			Port     int    `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
			Db       string `json:"db"`
			Charset  string `json:"charset"`
		} `json:"database"`
		Cache struct {
			Address string `json:"address"`
			Port    int    `json:"port"`
			Auth    string `json:"auth"`
		} `json:"cache"`
	} `json:"hosts"`
	Urls struct {
		Servers struct {
			Domain string `json:"domain"`
		} `json:"servers"`
	} `json:"urls"`
	Time struct {
		TimeZone   string `json:"time_zone"`
		TimeLayout string `json:"time_layout"`
	} `json:"time"`
	Security struct {
		Salt string `json:"salt"`
	} `json:"security"`

	Validation struct {
		Token struct {
			App int `json:"app"`
			Web int `json:"web"`
		} `json:"token"`
	} `json:"validation"`

	Pagination struct {
		MsgBoard int `json:"msgboard"`
	} `json:"pagination"`
}
