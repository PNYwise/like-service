package domain

type ExtConf struct {
	App      *App      `json:"app"`
	Database *Database `json:"database"`
	Post     *Post     `json:"post"`
}
type App struct {
	Port int `json:"port"`
}
type Database struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type Post struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
