package winter

type UpdateLog struct {
	Property string `json:"property"`
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}

func NewUpdateLog(property, oldValue, newValue string) *UpdateLog {
	return &UpdateLog{
		Property: property,
		OldValue: oldValue,
		NewValue: newValue,
	}
}
