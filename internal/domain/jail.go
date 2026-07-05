package domain

type JailStatus string

const (
	JailStatusRunning JailStatus = "running"
	JailStatusStopped JailStatus = "stopped"
)

type Jail struct {
	Name   string
	Status JailStatus
}
