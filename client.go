package integrations

import (
	"context"
	"fmt"

	"github.com/SKF/go-integrations-client/internal/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

type Client struct {
	c *rest.Client
}

func New(opts ...rest.Option) *Client {
	return &Client{
		c: rest.NewClient(
			append([]rest.Option{
				rest.WithBaseURL("https://api-integration.integration.enlight.skf.com"),
			}, opts...)...,
		),
	}
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api-integration.integration.enlight.skf.com")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api-integration.%s.integration.enlight.skf.com", stage))
}

func (c *Client) GetIntegrations(ctx context.Context) ([]Integration, error) {
	request := rest.Get("integrations")

	var response models.GetIntegrationsResponse

	if err := c.c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return nil, err
	}

	integrations := make([]Integration, 0, len(response.Integrations))

	for i := range response.Integrations {
		integration, err := c.GetIntegration(ctx, response.Integrations[i].ID)
		if err != nil {
			return nil, err
		}

		integrations = append(integrations, integration)
	}

	return integrations, nil
}

func (c *Client) GetIntegration(ctx context.Context, integrationID uuid.UUID) (Integration, error) {
	request := rest.Get("integrations/{id}").
		SetHeader("Accept", "application/json").
		Assign("id", integrationID)

	var response models.GetIntegrationResponse

	if err := c.c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return Integration{}, err
	}

	var integration Integration

	integration.FromInternal(response)

	return integration, nil
}
