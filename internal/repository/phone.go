package repository

type Phone struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewPhoneRepository(DbHandler DbHandler, sqlStore SqlStore) *Phone {
	return &Phone{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
