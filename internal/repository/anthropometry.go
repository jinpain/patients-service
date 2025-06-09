package repository

type Anthropometry struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewAnthropometryRepository(DbHandler DbHandler, sqlStore SqlStore) *Anthropometry {
	return &Anthropometry{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
