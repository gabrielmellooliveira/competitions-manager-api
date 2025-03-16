package database

import (
	"database/sql"
	"errors"

	database "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormAdapter struct {
	Url    string
	Client *gorm.DB
}

func NewGormAdapter(url string) database.Database {
	return &GormAdapter{
		Url: url,
	}
}

func (r *GormAdapter) Connect() error {
	db, err := gorm.Open(postgres.Open(r.Url), &gorm.Config{})
	if err != nil {
		return errors.New("falha ao conectar com a base de dados")
	}

	r.Client = db

	return nil
}

func (r *GormAdapter) MigrateEntity(entity database.Entity) {
	r.Client.AutoMigrate(entity)
}

func (r *GormAdapter) Create(entity database.Entity) error {
	return r.Client.Create(entity).Error
}

func (r *GormAdapter) Update(entity database.Entity, id any) error {
	return r.Client.Model(entity).Where("id = ?", id).Updates(entity).Error
}

func (r *GormAdapter) First(entity database.Entity, key string, value any) error {
	return r.Client.First(&entity, key+" = ?", value).Error
}

func (r *GormAdapter) Find(entities any, key string, value any) error {
	return r.Client.Table("supporters").Where(key+" = ?", value).Find(&entities).Error
}

func (r *GormAdapter) GetClient() (*sql.DB, error) {
	return r.Client.DB()
}
