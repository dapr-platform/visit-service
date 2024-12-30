package entity

// CamLiveStreamRequest represents the request body for starting a camera livestream
type CamLiveStreamRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	CameraID       string `json:"camera_id" binding:"required"`
	VisitRecordID  string `json:"visit_record_id" binding:"required"`
	DisableSaveMp4 bool   `json:"disable_save_mp4"`
}
type CamLiveStreamResponse struct {
	StreamUrlSuffix string `json:"stream_url_suffix"`
	Mp4StreamUrlSuffix string `json:"mp4_stream_url_suffix"`
	CameraID        string `json:"camera_id"`
	Status          int    `json:"status"`
}

