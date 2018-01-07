package service

import (
	"golang.org/x/net/context"

	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
)

func (s *SeppoServiceServer) SearchLogs(
	ctx context.Context,
	in *SeppoService.SearchLogsRequest,
) (
	*SeppoService.SearchLogsResponse,
	error,
) {
	res := &SeppoService.SearchLogsResponse{}
	db := s.getDB()

	query := db.Table("logs")

	if in.MessageType > 0 {
		query = query.Where("log_type = ?", in.MessageType)
	}

	if in.StartDate > 0 {
		query = query.Where("message_date > ?", in.StartDate)
	}

	if in.EndDate > 0 {
		query = query.Where("message_date > ?", in.EndDate)
	}

	if in.SearchWord != "" {
		query = query.Where("message LIKE ?", "%"+in.SearchWord+"%")
	}

	query.Count(&res.MaxLogs)

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}

	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	logs := []models.Log{}
	query.Find(&logs)

	for i := 0; i < len(logs); i++ {
		res.Logs = append(
			res.Logs,
			generators.NewLog(&logs[i]),
		)
	}

	return res, nil
}
