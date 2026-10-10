package taskx

import "github.com/hibiken/asynq"

const TaskNoteCreated = "note:created"

// NewNoteTask 生成 note:created 任务，载荷为 JSON {"id":1,"title":"..."}。
func NewNoteTask(id uint, title string) (*asynq.Task, error) {
	panic("TODO")
}
