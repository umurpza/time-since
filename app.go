package main

import (
	"context"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Time since wedding
func (a *App) Wedding(name string) string {
	since := time.Since(time.Date(2023, time.August, 23, 0, 0, 0, 0, time.UTC))

	return fmt.Sprintf("It's been %s since %s", since, name)
}

// CalculateElapsedTime calculates the elapsed time for an event given its UTC time.
func (a *App) CalculateElapsedTime(eventTime string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, eventTime)
	if err != nil {
		return "", fmt.Errorf("invalid time format: %w", err)
	}

	now := time.Now().UTC()
	duration := now.Sub(parsedTime)

	// Years and months are approximate as not all years and months have the same length
	years := int(duration.Hours()) / (24 * 365)
	remainingHoursAfterYears := int(duration.Hours()) % (24 * 365)

	months := remainingHoursAfterYears / (24 * 30)
	remainingHoursAfterMonths := remainingHoursAfterYears % (24 * 30)

	weeks := remainingHoursAfterMonths / (24 * 7)
	remainingHoursAfterWeeks := remainingHoursAfterMonths % (24 * 7)

	days := remainingHoursAfterWeeks / 24
	hours := remainingHoursAfterWeeks % 24

	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// Build the result string
	result := ""
	if years > 0 {
		result += fmt.Sprintf("%dy ", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%dmo ", months)
	}
	if weeks > 0 {
		result += fmt.Sprintf("%dw ", weeks)
	}
	if days > 0 {
		result += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 || result == "" { // Always show seconds if no larger unit is available
		result += fmt.Sprintf("%ds ", seconds)
	}

	return result + "ago", nil
}
