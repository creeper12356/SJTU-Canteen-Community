package app

import "fmt"

func InitAll() error {
	err := InitMySQL()
	if err != nil {
		return fmt.Errorf("MySQL initialization failed: %w", err)
	}
	return nil
}
