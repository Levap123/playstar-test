package postgres

import (
	"github.com/Levap123/playstar-test/logging_service/entities"
	"github.com/Levap123/playstar-test/logging_service/logs"
	"github.com/jmoiron/sqlx"
)

type LogsRepo struct {
	DB     *sqlx.DB
	logger *logs.Logger
}

// type LogMessage struct {
// 	Level   string `json:"level"`
// 	Service string `json:"service"`
// 	Error   string `json:"error"`
// 	Time    string `json:"time"`
// 	Caller  string `json:"caller"`
// 	Message string `json:"message"`
// }

func (lr *LogsRepo) Insert(logMsg entities.LogMessage) {
	query := `INSERT INTO logs (level, service, error, time, caller, message) 
	VALUES (:level, :service, :error, :time, :caller, :message);`
	
	if _, err := lr.DB.NamedExec(query, logMsg); err != nil {
		lr.logger.Err(err).Msg("error in inserting log message to DB")
	}
}
