package configs

import (
	c "crypto-braza-tokens-admin/constants"
	d "crypto-braza-tokens-admin/domain"
	db "crypto-braza-tokens-admin/repositories/mongo"
	kvs "crypto-braza-tokens-admin/utils/keys-values"
	l "crypto-braza-tokens-admin/utils/logger"
	"os"
)

type Configs struct {
	SettingsDomain   *d.SettingsDomain
	OperationsDomain *d.OperationsDomain
}

func InitRequirements() {
	validateRequiredEnvs()
	l.NewLogger(os.Getenv(c.LOG_LEVEL))
	db.Start(os.Getenv(c.MONGO_DATABASE))
	kvs.Start()
}
