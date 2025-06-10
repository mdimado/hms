package patient

import (
	"hospital-management/internal/models"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePatient(req models.CreatePatientRequest, createdBy uint) (*models.Patient, error) {
	patient := &models.Patient{
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Email:           req.Email,
		Phone:           req.Phone,
		DateOfBirth:     req.DateOfBirth,
		Gender:          req.Gender,
		Address:         req.Address,
		EmergencyContact: req.EmergencyContact,
		BloodGroup:      req.BloodGroup,
		Allergies:       req.Allergies,
		InsuranceNumber: req.InsuranceNumber,
		RegistrationDate: time.Now(),
		CreatedBy:       createdBy,
	}

	if err := s.repo.Create(patient); err != nil {
		return nil, err
	}

	return s.repo.GetByID(patient.ID)
}

func (s *Service) GetAllPatients() ([]models.Patient, error) {
	return s.repo.GetAll()
}

func (s *Service) GetPatientByID(id uint) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdatePatient(id uint, req models.UpdatePatientRequest) (*models.Patient, error) {
	patient, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.FirstName != "" {
		patient.FirstName = req.FirstName
	}
	if req.LastName != "" {
		patient.LastName = req.LastName
	}
	// ... update other fields

	if err := s.repo.Update(patient); err != nil {
		return nil, err
	}

	return s.repo.GetByID(id)
}

func (s *Service) DeletePatient(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) UpdateMedicalInfo(id uint, req models.UpdateMedicalInfoRequest) (*models.Patient, error) {
	patient, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	patient.MedicalHistory = req.MedicalHistory
	patient.CurrentMedications = req.CurrentMedications
	patient.Allergies = req.Allergies

	if err := s.repo.Update(patient); err != nil {
		return nil, err
	}

	return s.repo.GetByID(id)
}