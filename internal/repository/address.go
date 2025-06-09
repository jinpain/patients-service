package repository

type Address struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewAddressRepository(DbHandler DbHandler, sqlStore SqlStore) *Address {
	return &Address{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
