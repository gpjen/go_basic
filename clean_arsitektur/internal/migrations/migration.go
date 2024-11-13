// internal/migrations/migration.go
package migrations

import (
	"clean_arsitektur/internal/domain/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// MigrationHistory tracks which migrations have been run
type MigrationHistory struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex"`
	BatchNo   int
	CreatedAt time.Time
}

// Migration represents a single migration
type Migration struct {
	Name string
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}

// List of all migrations
var migrations = []Migration{
	{
		Name: "create_users_table",
		Up: func(db *gorm.DB) error {
			return db.AutoMigrate(&entity.User{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&entity.User{})
		},
	},
	// Add more migrations here
	{
		Name: "add_role_to_users",
		Up: func(db *gorm.DB) error {
			// Example of adding a column
			type User struct {
				Role string `gorm:"type:varchar(50);default:'user'"`
			}
			return db.AutoMigrate(&User{})
		},
		Down: func(db *gorm.DB) error {
			// Remove the column
			return db.Migrator().DropColumn(&entity.User{}, "role")
		},
	},
}

// MigrationManager handles database migrations
type MigrationManager struct {
	db *gorm.DB
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(db *gorm.DB) *MigrationManager {
	return &MigrationManager{db: db}
}

// Initialize prepares the migration system
func (m *MigrationManager) Initialize() error {
	// Create migration history table if it doesn't exist
	return m.db.AutoMigrate(&MigrationHistory{})
}

// GetLastBatchNo gets the last batch number
func (m *MigrationManager) GetLastBatchNo() (int, error) {
	var lastBatch MigrationHistory
	result := m.db.Order("batch_no desc").First(&lastBatch)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, result.Error
	}
	return lastBatch.BatchNo, nil
}

// Up runs all pending migrations
func (m *MigrationManager) Up() error {
	// Get the last batch number
	lastBatchNo, err := m.GetLastBatchNo()
	if err != nil {
		return err
	}

	currentBatch := lastBatchNo + 1

	// Get list of completed migrations
	var completed []string
	if err := m.db.Model(&MigrationHistory{}).Pluck("name", &completed).Error; err != nil {
		return err
	}

	// Create a map for O(1) lookup
	completedMap := make(map[string]bool)
	for _, name := range completed {
		completedMap[name] = true
	}

	// Run pending migrations
	for _, migration := range migrations {
		if !completedMap[migration.Name] {
			fmt.Printf("Running migration: %s\n", migration.Name)

			// Run migration within transaction
			err := m.db.Transaction(func(tx *gorm.DB) error {
				if err := migration.Up(tx); err != nil {
					return err
				}

				// Record successful migration
				return tx.Create(&MigrationHistory{
					Name:      migration.Name,
					BatchNo:   currentBatch,
					CreatedAt: time.Now(),
				}).Error
			})

			if err != nil {
				return fmt.Errorf("error running migration %s: %v", migration.Name, err)
			}

			fmt.Printf("Completed migration: %s\n", migration.Name)
		}
	}

	return nil
}

// Down rolls back the last batch of migrations
func (m *MigrationManager) Down() error {
	lastBatchNo, err := m.GetLastBatchNo()
	if err != nil {
		return err
	}

	if lastBatchNo == 0 {
		fmt.Println("No migrations to roll back")
		return nil
	}

	// Get migrations from the last batch
	var histories []MigrationHistory
	if err := m.db.Where("batch_no = ?", lastBatchNo).Find(&histories).Error; err != nil {
		return err
	}

	// Create map of migration names to Migration objects
	migrationMap := make(map[string]Migration)
	for _, migration := range migrations {
		migrationMap[migration.Name] = migration
	}

	// Roll back each migration in reverse order
	for i := len(histories) - 1; i >= 0; i-- {
		history := histories[i]
		migration, exists := migrationMap[history.Name]
		if !exists {
			return fmt.Errorf("migration %s not found", history.Name)
		}

		fmt.Printf("Rolling back migration: %s\n", history.Name)

		err := m.db.Transaction(func(tx *gorm.DB) error {
			if err := migration.Down(tx); err != nil {
				return err
			}

			// Remove migration history
			return tx.Delete(&MigrationHistory{}, "name = ?", history.Name).Error
		})

		if err != nil {
			return fmt.Errorf("error rolling back migration %s: %v", history.Name, err)
		}

		fmt.Printf("Rolled back migration: %s\n", history.Name)
	}

	return nil
}

// Status prints the current migration status
func (m *MigrationManager) Status() error {
	var histories []MigrationHistory
	if err := m.db.Order("created_at").Find(&histories).Error; err != nil {
		return err
	}

	fmt.Println("\nMigration Status:")
	fmt.Println("------------------")

	completedMap := make(map[string]MigrationHistory)
	for _, history := range histories {
		completedMap[history.Name] = history
	}

	for _, migration := range migrations {
		history, completed := completedMap[migration.Name]
		status := "Pending"
		batch := ""
		if completed {
			status = "Completed"
			batch = fmt.Sprintf("(Batch %d)", history.BatchNo)
		}
		fmt.Printf("%-30s %s %s\n", migration.Name, status, batch)
	}

	return nil
}
