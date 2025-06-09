package repository

type InsurancePolicy struct {
	dbHandler DbHandler
	sqlStore  SqlStore
}

func NewInsurancePolicyRepository(DbHandler DbHandler, sqlStore SqlStore) *InsurancePolicy {
	return &InsurancePolicy{
		dbHandler: DbHandler,
		sqlStore:  sqlStore,
	}
}
