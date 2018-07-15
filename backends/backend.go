package backends

type StandardBackender interface{
	GetResult(taskID string)
}