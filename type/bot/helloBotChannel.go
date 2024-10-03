package bot

type TestCallType struct {
	Content    string      `json:"content"`
	Components []Component `json:"components"`
}

type Component struct {
	Type       int64 `json:"type"`
	Components []struct {
		Type     int64  `json:"type"`
		Label    string `json:"label"`
		Style    int64  `json:"style"`
		CustomID string `json:"custom_id"`
	} `json:"components"`
}
