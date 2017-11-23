package services

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (s *SeppoServiceServer) SearchLogs(
	ctx context.Context,
	in *SeppoService.SearchLogsRequest,
) (
	*SeppoService.SearchLogsResponse,
	error,
) {
	res := &SeppoService.SearchLogsResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	if in.MessageType > 0 {
		queryMods = append(
			queryMods,
			qm.Where("log_type = ?", in.MessageType),
		)
	}

	if in.StartDate > 0 {
		queryMods = append(
			queryMods,
			qm.Where("message_date >= ?", in.StartDate),
		)
	}

	if in.EndDate > 0 {
		queryMods = append(
			queryMods,
			qm.Where("message_date <= ?", in.EndDate),
		)
	}

	if in.SearchWord != "" {
		queryMods = append(
			queryMods,
			qm.Where("message LIKE ?", "%"+in.SearchWord+"%"),
		)
	}

	c, _ := models.Logs(
		newDb,
		queryMods...,
	).Count()

	res.MaxLogs = uint64(c)

	if in.Offset > 0 {
		queryMods = append(
			queryMods,
			qm.Offset(int(in.Offset)),
		)
	} else {
		queryMods = append(
			queryMods,
			qm.Offset(10000),
		)
	}

	if in.Limit > 0 {
		queryMods = append(
			queryMods,
			qm.Limit(int(in.Limit)),
		)
	}

	logs, _ := models.Logs(
		newDb,
		queryMods...,
	).All()

	for _, log := range logs {
		res.Logs = append(
			res.Logs,
			generators.NewLog(log),
		)
	}

	return res, nil
}
