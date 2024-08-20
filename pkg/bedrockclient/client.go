package bedrockclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/dacsang97/go-claude-bedrock/pkg/message"
	"github.com/dacsang97/go-claude-bedrock/pkg/model"
	"github.com/goccy/go-json"
)

type Client struct {
	model         model.ClaudeBedrockModel
	bedrockClient *bedrockruntime.Client
}

func NewClient(region string, model model.ClaudeBedrockModel) (*Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	return &Client{
		model:         model,
		bedrockClient: bedrockruntime.NewFromConfig(sdkConfig),
	}, nil
}

func (c *Client) CreateMessages(ctx context.Context, request message.MessagesRequest) (*message.MessagesResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	output, err := c.bedrockClient.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(string(c.model)),
		ContentType: aws.String("application/json"),
		Body:        body,
	})

	if err != nil {
		return nil, err
	}

	var response message.MessagesResponse
	err = json.Unmarshal(output.Body, &response)

	return &response, err
}
