package examples_internal

type Conf struct {
	DB DatabaseConf `toml:"database" json:"database"`
}

type DatabaseConf struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
}
