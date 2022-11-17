package model

import "database/sql"

type Application struct {
	ID          uint `gorm:"primaryKey"`
	LeaderID    int
	Leader      User `gorm:"foreignKey:LeaderID;"`
	Name        string
	Description sql.NullString
	Servers     []Server `gorm:"many2many:Cluster"`
	Status      int
}

type Server struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description sql.NullString
	IpAddress   sql.NullString
	Apps        []Application `gorm:"many2many:Cluster"`
	IsExternal  bool
}

type Cluster struct {
	ApplicationID uint `gorm:"primaryKey"`
	ServerID      uint `gorm:"primaryKey"`
	Name          string
	Description   sql.NullString
}
