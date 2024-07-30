package entities

type Config struct {
	Env Env 
}

type Env struct {
	Local LocalConf 
}

type LocalConf struct {
	Port int 
	Database Database
}
 
type Database struct {
	Host string
	Port int
	User string
	Password string
	DB string
}