package service

type MeetingCreateRequest struct {
	Name     string `json:"name, omitempty"`
	CreateAt int    `json:"create_at"`
	EndAt    int64  `json:"end_at"`
}
