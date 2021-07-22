package config

type sqlConf struct {
	Dialect string
	Name    string
	Pwd     string
	Host    string
	Port    string
	Db      string
}

type redisConf struct {
	Host string
	Port string
	DB   int
	Pwd  string
}

var SqlConfig = func() *sqlConf {
	return &sqlConf{
		Dialect: "mysql",
		Name:    "typecho",
		Pwd:     "3e1e1758",
		Host:    "127.0.0.1",
		Port:    "3306",
		Db:      "typecho",
	}
}()

var RedisConfig = func() *redisConf {
	return &redisConf{
		Host: "127.0.0.1",
		Port: "6379",
		DB:   0,
		Pwd:  "123456",
	}
}()

func init() {

}
