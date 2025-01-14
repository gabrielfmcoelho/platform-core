package parser

import (
	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal"
)

// Parse Service to PublicService
func ToPublicService(s domain.Service) domain.PublicService {
	return domain.PublicService{
		ID:         s.ID,
		Name:       s.Name,
		AppUrl:     s.AppUrl,
		LastUpdate: s.LastUpdate,
		Status:     s.Status,
	}
}

// Parse Service to HubService
func ToHubService(s domain.Service) domain.HubService {
	return domain.HubService{
		ID:            s.ID,
		Name:          s.Name,
		IconUrl:       s.IconUrl,
		Description:   s.Description,
		ScreenshotUrl: s.ScreenshotUrl,
		LastUpdate:    s.LastUpdate,
		Status:        s.Status,
		Price:         s.Price,
	}
}

// Parse Service to MarketingService
func ToMarketingService(s domain.Service) domain.MarketingService {
	return domain.MarketingService{
		ID:            s.ID,
		IconUrl:       s.IconUrl,
		MarketingName: s.MarketingName,
		TagLine:       s.TagLine,
		Description:   s.Description,
		Benefits:      internal.ParseDelimitedStrings(s.Benefits),
		Features:      internal.ParseDelimitedStrings(s.Features),
		Tags:          internal.ParseDelimitedStrings(s.Tags),
	}
}

// Parse Service to UseService
func ToUseService(s domain.Service) domain.UseService {
	return domain.UseService{
		Service: ToPublicService(s),
	}
}
