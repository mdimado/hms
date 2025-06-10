package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null;check:role IN ('receptionist','doctor')"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Patient struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	FirstName      string    `json:"first_name" gorm:"not null"`
	LastName       string    `json:"last_name" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique"`
	Phone          string    `json:"phone" gorm:"not null"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender" gorm:"check:gender IN ('male','female','other')"`
	Address        string    `json:"address"`
	EmergencyContact string  `json:"emergency_contact"`
	BloodGroup     string    `json:"blood_group"`
	Allergies      string    `json:"allergies"`
	MedicalHistory string    `json:"medical_history"`
	CurrentMedications string `json:"current_medications"`
	InsuranceNumber string   `json:"insurance_number"`
	RegistrationDate time.Time `json:"registration_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy      uint      `json:"created_by"`
	CreatedByUser  User      `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role" binding:"required,oneof=receptionist doctor"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone"`
}

type CreatePatientRequest struct {
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	Email          string    `json:"email" binding:"omitempty,email"`
	Phone          string    `json:"phone" binding:"required"`
	DateOfBirth    time.Time `json:"date_of_birth" binding:"required"`
	Gender         string    `json:"gender" binding:"required,oneof=male female other"`
	Address        string    `json:"address"`
	EmergencyContact string  `json:"emergency_contact"`
	BloodGroup     string    `json:"blood_group"`
	Allergies      string    `json:"allergies"`
	InsuranceNumber string   `json:"insurance_number"`
}

type UpdatePatientRequest struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" binding:"omitempty,email"`
	Phone          string    `json:"phone"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender" binding:"omitempty,oneof=male female other"`
	Address        string    `json:"address"`
	EmergencyContact string  `json:"emergency_contact"`
	BloodGroup     string    `json:"blood_group"`
	Allergies      string    `json:"allergies"`
	InsuranceNumber string   `json:"insurance_number"`
}

type UpdateMedicalInfoRequest struct {
	MedicalHistory     string `json:"medical_history"`
	CurrentMedications string `json:"current_medications"`
	Allergies          string `json:"allergies"`
}