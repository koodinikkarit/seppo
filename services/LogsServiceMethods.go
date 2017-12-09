package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) SearchLogs(
	ctx context.Context,
	in *SeppoService.SearchLogsRequest,
) (
	*SeppoService.SearchLogsResponse,
	error,
) {
	res := &SeppoService.SearchLogsResponse{}
	newDb := s.getGormDB()
	defer newDb.Close()

	query := newDb.Table("logs")

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

	logs := []db.Log{}
	query.Find(&logs)

	for i := 0; i < len(logs); i++ {
		res.Logs = append(
			res.Logs,
			generators.NewLog(&logs[i]),
		)
	}

	return res, nil
}
