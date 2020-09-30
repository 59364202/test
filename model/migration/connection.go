package migrate

import (
	"haii.or.th/api/util/pqx"
)

const pgpool_Nhc = "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/nhc?sslmode=disable"
const pgpool_Staging = "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/staging?sslmode=disable"

const (
	Cmd                 = "regen"
	Cmd_RegenDataImg    = "RegenDataImg"
	Cmd_RegenData       = "RegenData"
	Cmd_RegenMasterData = "RegenMasterData"
)

func OpenNhc() (*pqx.DB, error) {
	return pqx.OpenDb(pgpool_Nhc)
}

func OpenStaging() (*pqx.DB, error) {
	return pqx.OpenDb(pgpool_Staging)
}
