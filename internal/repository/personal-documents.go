package repository

type PersonalDocument struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewPersonalDocumentRepository(DbHandler DbHandler, sqlStore SqlStore) *PersonalDocument {
	return &PersonalDocument{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
