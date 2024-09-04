package winter

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewKeyValue(key string, value string) *KeyValue {
	return &KeyValue{
		Key:   key,
		Value: value,
	}
}
