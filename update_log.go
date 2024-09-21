package winter

type UpdateLog struct {
	Property string `json:"property"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

func NewUpdateLog(property, oldValue, newValue string) *UpdateLog {
	return &UpdateLog{
		Property: property,
		OldValue: oldValue,
		NewValue: newValue,
	}
}
