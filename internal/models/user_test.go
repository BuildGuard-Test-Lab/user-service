package models

import (
	"testing"
	"time"
)

func TestUserStruct(t *testing.T) {
	now := time.Now()
	u := User{
		ID:        1,
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if u.ID != 1 {
		t.Errorf("expected ID 1, got %d", u.ID)
	}
	if u.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", u.Email)
	}
	if u.Name != "Test User" {
		t.Errorf("expected name Test User, got %s", u.Name)
	}
	if u.CreatedAt.IsZero() {
		t.Error("expected non-zero CreatedAt")
	}
	if u.UpdatedAt.IsZero() {
		t.Error("expected non-zero UpdatedAt")
	}
}

func TestUserZeroValue(t *testing.T) {
	var u User
	if u.ID != 0 {
		t.Errorf("expected zero ID, got %d", u.ID)
	}
	if u.Email != "" {
		t.Errorf("expected empty email, got %s", u.Email)
	}
}
