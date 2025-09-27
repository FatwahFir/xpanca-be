package dto

type ImageResponse struct {
	ID          uint   `json:"id"`
	URL         string `json:"url"`
	IsThumbnail bool   `json:"is_thumbnail"`
}
