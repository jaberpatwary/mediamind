package database

import (
	"app/src/model"
	"app/src/utils"

	"gorm.io/gorm"
)

// MigrateProjectsTable drops the old projects table and creates a new one
func MigrateProjectsTable(db *gorm.DB) error {
	// Drop the old projects table if it exists
	if db.Migrator().HasTable(&model.Project{}) {
		utils.Log.Info("Dropping old projects table...")
		if err := db.Migrator().DropTable(&model.Project{}); err != nil {
			utils.Log.Errorf("Failed to drop projects table: %v", err)
			return err
		}
	}

	// Create the new projects table
	utils.Log.Info("Creating new projects table...")
	if err := db.AutoMigrate(&model.Project{}); err != nil {
		utils.Log.Errorf("Failed to create projects table: %v", err)
		return err
	}

	utils.Log.Info("Projects table migration completed successfully")
	return nil
}
