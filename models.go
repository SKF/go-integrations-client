package integrations

import (
	"strconv"
	"time"

	"github.com/SKF/go-integrations-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

const (
	StatusUnknown = "unknown"
	StatusRunning = "running"
	StatusStopped = "stopped"
)

const (
	IntegrationTypeUnknown  = "unknown"
	IntegrationTypeAnalyst  = "analyst"
	IntegrationTypeObserver = "observer"
)

const (
	IntegrationVersionUnknown = iota
	IntegrationVersion1
	IntegrationVersion2
)

type Integration struct {
	ID          uuid.UUID
	Name        string
	CompanyName string
	Status      string
	CreatedAt   time.Time
	Type        string
	Version     int
}

func (i Integration) IsRunning() bool {
	return i.Status == StatusRunning
}

func (i *Integration) FromInternal(g models.GetIntegrationResponse) {
	i.ID = uuid.UUID(g.IntegrationDetails.ID)
	i.Name = g.IntegrationDetails.Name
	i.CompanyName = g.IntegrationDetails.CompanyName
	i.Status = g.IntegrationDetails.Status.Status
	i.CreatedAt = time.Unix(g.IntegrationDetails.CreatedTimestamp, 0).UTC()
	i.Type = g.IntegrationDetails.Type

	switch i.Status {
	case StatusRunning, StatusStopped:
	default:
		i.Status = StatusUnknown
	}

	switch i.Type {
	case IntegrationTypeAnalyst, IntegrationTypeObserver:
	default:
		i.Type = IntegrationTypeUnknown
	}

	version, err := strconv.ParseInt(g.IntegrationDetails.Version, 10, 64)
	if err != nil {
		version = 0
	}

	i.Version = int(version)
}
