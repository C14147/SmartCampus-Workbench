package db

import (
    "context"
    "errors"
    "fmt"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// Connect creates a gorm DB with sensible pooling and retries on transient failures.
func Connect(dsn string) (*gorm.DB, error) {
    var lastErr error
    // retry with exponential backoff
    for i := 0; i < 5; i++ {
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Silent),
        })
        if err == nil {
            sqlDB, err := db.DB()
            if err != nil {
                return nil, fmt.Errorf("failed to obtain sql.DB: %w", err)
            }
            // pooling defaults
            sqlDB.SetMaxOpenConns(25)
            sqlDB.SetMaxIdleConns(25)
            sqlDB.SetConnMaxLifetime(5 * time.Minute)

            // quick ping to verify connectivity
            ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
            defer cancel()
            if err := sqlDB.PingContext(ctx); err != nil {
                lastErr = err
                dbCloseErr := sqlDB.Close()
                if dbCloseErr != nil {
                    // ignore
                }
            } else {
                return db, nil
            }
        } else {
            lastErr = err
        }

        // backoff before next attempt
        time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
    }
    if lastErr == nil {
        lastErr = errors.New("unknown error connecting to database")
    }
    return nil, fmt.Errorf("db connect retries exhausted: %w", lastErr)
}
