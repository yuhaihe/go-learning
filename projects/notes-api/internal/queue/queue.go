package queue

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeNoteCreated = "note:created"

type Payload struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type Enqueuer interface {
	EnqueueNoteCreated(ctx context.Context, id uint, title string) error
}

type AsynqEnqueuer struct {
	client *asynq.Client
}

func NewAsynqEnqueuer(client *asynq.Client) *AsynqEnqueuer {
	return &AsynqEnqueuer{client: client}
}

func (e *AsynqEnqueuer) EnqueueNoteCreated(ctx context.Context, id uint, title string) error {
	body, err := json.Marshal(Payload{ID: id, Title: title})
	if err != nil {
		return err
	}
	_, err = e.client.EnqueueContext(ctx, asynq.NewTask(TypeNoteCreated, body))
	return err
}
