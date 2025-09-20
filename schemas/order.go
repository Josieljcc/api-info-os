package schemas

import (
	"time"

	"gorm.io/gorm"
)

// OrderStatus representa os possíveis status de uma ordem de serviço
type OrderStatus string

const (
	StatusOpen       OrderStatus = "open"        // OS recém criada
	StatusInProgress OrderStatus = "in_progress" // OS em execução
	StatusWaiting    OrderStatus = "waiting"     // Aguardando peças ou cliente
	StatusCompleted  OrderStatus = "completed"   // OS finalizada com sucesso
	StatusCancelled  OrderStatus = "cancelled"   // OS cancelada
	StatusSuspended  OrderStatus = "suspended"   // OS temporariamente suspensa
)

// IsValid verifica se o status é válido
func (s OrderStatus) IsValid() bool {
	switch s {
	case StatusOpen, StatusInProgress, StatusWaiting, StatusCompleted, StatusCancelled, StatusSuspended:
		return true
	default:
		return false
	}
}

type Order struct {
	gorm.Model
	OpeningDate  time.Time   `gorm:"not null"`
	ForecastDate time.Time   `gorm:"not null"`
	ClosingDate  *time.Time  `gorm:"null"`
	Status       OrderStatus `gorm:"not null;type:varchar(20)"`
	Comment      string      `gorm:"size:255"`
	ClientID     uint        `gorm:"not null"`
	Client       Client      `gorm:"foreignKey:ClientID;references:ID"`
	TechnicianID uint        `gorm:"not null"`
	Technician   Technician  `gorm:"foreignKey:TechnicianID;references:ID"`
	Services     []Service   `gorm:"many2many:order_services;"`
	Parts        []Part      `gorm:"many2many:order_parts;"`
}

type OrderResponse struct {
	ID           uint               `json:"id"`
	OpeningDate  time.Time          `json:"openingDate"`
	ForecastDate time.Time          `json:"forecastDate"`
	ClosingDate  *time.Time         `json:"closingDate"`
	Status       OrderStatus        `json:"status"`
	Comment      string             `json:"comment"`
	ClientID     uint               `json:"clientId"`
	Client       ClientResponse     `json:"client"`
	TechnicianID uint               `json:"technicianId"`
	Technician   TechnicianResponse `json:"technician"`
	Services     []ServiceResponse  `json:"services"`
	Parts        []PartResponse     `json:"parts"`
}

func (o Order) ToResponse() OrderResponse {
	var partsResponse []PartResponse
	for _, part := range o.Parts {
		partsResponse = append(partsResponse, part.ToResponse())
	}

	var servicesResponse []ServiceResponse
	for _, service := range o.Services {
		servicesResponse = append(servicesResponse, service.ToResponse())
	}

	return OrderResponse{
		ID:           o.ID,
		OpeningDate:  o.OpeningDate,
		ForecastDate: o.ForecastDate,
		ClosingDate:  o.ClosingDate,
		Status:       o.Status,
		Comment:      o.Comment,
		ClientID:     o.ClientID,
		Client:       o.Client.ToResponse(),
		TechnicianID: o.TechnicianID,
		Technician:   o.Technician.ToResponse(),
		Services:     servicesResponse,
		Parts:        partsResponse,
	}
}
