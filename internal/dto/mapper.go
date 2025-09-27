package dto

import "github.com/FatwahFir/xpanca-be/internal/domain"

func ToImageResponse(m domain.Image) ImageResponse {
	return ImageResponse{
		ID: m.ID, URL: m.URL, IsThumbnail: m.IsThumbnail,
	}
}

func ToProductResponse(p domain.Product) ProductResponse {
	out := ProductResponse{
		ID: p.ID, Name: p.Name, Category: p.Category, Price: p.Price,
		Description: p.Description, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt,
	}
	if len(p.Images) > 0 {
		out.Images = make([]ImageResponse, 0, len(p.Images))
		for _, img := range p.Images {
			out.Images = append(out.Images, ToImageResponse(img))
		}
	}
	return out
}

func ToProductListResponse(pp []domain.Product) []ProductResponse {
	if len(pp) == 0 {
		return []ProductResponse{}
	}
	out := make([]ProductResponse, 0, len(pp))
	for _, p := range pp {
		out = append(out, ToProductResponse(p))
	}
	return out
}

func ToUserResponse(u domain.User) UserResponse {
	return UserResponse{
		ID: u.ID, Username: u.Username, Role: u.Role,
	}
}
