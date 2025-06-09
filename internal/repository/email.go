package repository

type Email struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewEmailRepository(DbHandler DbHandler, sqlStore SqlStore) *Email {
	return &Email{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
