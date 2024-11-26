package entity

// CamLiveStreamRequest represents the request body for starting a camera livestream
type CamLiveStreamRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	CameraID string `json:"camera_id" binding:"required"`
}
type CamLiveStreamResponse struct {
	StreamUrlSuffix string `json:"stream_url_suffix"`
	CameraID        string `json:"camera_id"`
	Status          int    `json:"status"`
}
