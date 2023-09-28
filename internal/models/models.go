package models

import "github.com/SKF/go-utility/v2/uuid"

type GetIntegrationsResponse struct {
	Integrations []GetIntegrationsResponseIntegration `json:"integrations"`
}

type GetIntegrationsResponseIntegration struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type GetIntegrationResponse struct {
	IntegrationDetails GetIntegrationResponseDetails `json:"integrationDetails"`
	IntegrationSecret  GetIntegrationResponseSecret  `json:"integrationSecrets"`
}

type GetIntegrationResponseDetails struct {
	ID               string                       `json:"id"`
	Name             string                       `json:"name"`
	CompanyName      string                       `json:"companyName"`
	CreatedTimestamp int64                        `json:"createdTimestamp"`
	Type             string                       `json:"type"`
	Version          string                       `json:"version"`
	AgentID          string                       `json:"agentId"`
	Config           GetIntegrationResponseConfig `json:"config"`
	Status           GetIntegrationResponseStatus `json:"status"`
}

type GetIntegrationResponseConfig struct {
	RootNodeID   string `json:"hierarchyRootId"`
	RootNodeType string `json:"hierarchyRootType"`
}

type GetIntegrationResponseStatus struct {
	Status string `json:"status"`
}

type GetIntegrationResponseSecret struct {
	Type     string `json:"type"`
	Username string `json:"username"`
}
