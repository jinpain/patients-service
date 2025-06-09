package repository

type Patient struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewPatientRepository(DbHandler DbHandler, sqlStore SqlStore) *Patient {
	return &Patient{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
