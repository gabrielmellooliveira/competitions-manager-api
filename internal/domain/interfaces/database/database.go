package interfaces

import "database/sql"

type Database interface {
	Connect() error
	MigrateEntity(entity Entity)
	Create(entity Entity) error
	Update(entity Entity, id any) error
	First(entity Entity, key string, value any) error
	Find(entities any, key string, value any) error
	GetClient() (*sql.DB, error)
}
