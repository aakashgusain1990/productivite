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
		sqlDB, err := sql.Open("mysql", "")
		if err != nil {
			panic(err)
		}
		dbInstance := bun.NewDB(sqlDB, mysqldialect.New())
	})
	return dbInstance
}

func (MySQLDatabase *MySQLDatabase) GetDB() *bun.DB {
	return dbInstance.DB
}