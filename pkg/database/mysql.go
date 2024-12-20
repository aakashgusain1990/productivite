package database


type MySQLDatabase struct {
	DB *bun.DB
}

var (
	once sync.Once
	dbInstance *MySQLDatabase
)

func NewMySQLDatabase(conf *entities.Config) *MySQLDatabase {
	once.Do(func() {
		sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName))
		if err != nil {
			panic(err)
		}
	})
	return dbInstance
}

func (MySQLDatabase *MySQLDatabase) GetDB() *bun.DB {
	return dbInstance.DB
}
