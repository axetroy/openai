package fine_tune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

type CreateFineTuneParams struct {
	TrainingFile                 string   `json:"training_file"`
	ValidationFile               *string  `json:"validation_file,omitempty"`
	Model                        *string  `json:"model,omitempty"`
	NEpochs                      *int     `json:"n_epochs,omitempty"`
	BatchSize                    *int     `json:"batch_size,omitempty"`
	LearningRateMultiplier       *int     `json:"learning_rate_multiplier,omitempty"`
	PromptLossWeight             *int     `json:"prompt_loss_weight,omitempty"`
	ComputeClassificationMetrics *bool    `json:"compute_classification_metrics,omitempty"`
	ClassificationNClasses       *int     `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  *string  `json:"classification_positive_class,omitempty"`
	ClassificationBetas          []string `json:"classification_betas,omitempty"`
	Suffix                       *string  `json:"suffix,omitempty"`
}

// docs: https://platform.openai.com/docs/api-reference/embeddings/create
func (this *FineTune) CreateFineTune(params CreateFineTuneParams) (*CreateFineTuneResponse, error) {
	url := fmt.Sprintf("%s/v1/fine-tunes", this.domain)

	body, err := json.Marshal(params)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	source, err := event_source.NewEventSource[any](url, "POST", http.Header{
		"Authorization": []string{"Bearer " + this.apiKey},
		"Content-Type":  []string{"application/json"},
	}, bytes.NewBuffer(body))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer source.Close()

	b, err := io.ReadAll(source.Response().Body)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var data CreateFineTuneResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
