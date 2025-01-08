package types

type Task struct {
	BaseModel
	Title       string
	Description string
	ProjectID   string
	Done        bool
	Priority    uint8
}
