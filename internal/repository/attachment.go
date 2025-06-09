package repository

type Attahment struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewAttahmentRepository(DbHandler DbHandler, sqlStore SqlStore) *Attahment {
	return &Attahment{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
