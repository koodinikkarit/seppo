package services

import (
	"time"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s SeppoServiceServer) CreateSchedule(
	ctx context.Context,
	in *SeppoService.CreateScheduleRequest,
) (
	*SeppoService.CreateScheduleResponse,
	error,
) {
	res := &SeppoService.CreateScheduleResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	d1 := time.Unix(in.Start, 0)
	d2 := time.Unix(in.End, 0)

	newSchedule := db.Schedule{
		Name:  in.Name,
		Start: &d1,
		End:   &d2,
	}

	newDb.Create(&newSchedule)

	res.Schedule = NewSchedule(&newSchedule)

	return res, nil
}

func (s SeppoServiceServer) UpdateSchedule(
	ctx context.Context,
	in *SeppoService.UpdateScheduleRequest,
) (
	*SeppoService.UpdateScheduleResponse,
	error,
) {
	res := &SeppoService.UpdateScheduleResponse{}
	tx := s.getDB().Begin()
	defer tx.Close()

	var schedule db.Schedule
	tx.First(&schedule, in.ScheduleId)

	if schedule.ID > 0 {
		schedule.Name = in.Name

		if len(in.AddSongIds) > 0 {
			for i := 0; i < len(in.AddSongIds); i++ {
				tx.Create(&db.ScheduleVariation{
					ScheduleID:  in.ScheduleId,
					VariationID: in.AddSongIds[i],
				})
			}
		}

		if len(in.RemoveSongIds) > 0 {
			tx.Where("variation_version_id in (?)", in.RemoveSongIds).
				Delete(&db.ScheduleVariation{})
		}
		res.Schedule = NewSchedule(&schedule)
		res.Success = true
	} else {
		res.Success = false
	}

	tx.Commit()

	return res, nil
}

func (s SeppoServiceServer) RemoveSchedule(
	ctx context.Context,
	in *SeppoService.RemoveScheduleRequest,
) (
	*SeppoService.RemoveScheduleResponse,
	error,
) {
	res := &SeppoService.RemoveScheduleResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	var schedule db.Schedule
	newDb.First(&schedule, in.ScheduleId)

	if schedule.ID > 0 {
		newDb.Delete(&schedule)
		res.Success = true
	} else {
		res.Success = false
	}

	return res, nil
}

func (s *SeppoServiceServer) SearchSchedules(
	ctx context.Context,
	in *SeppoService.SearchSchedulesRequest,
) (
	*SeppoService.SearchSchedulesResponse,
	error,
) {
	res := &SeppoService.SearchSchedulesResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	schedules := []db.Schedule{}

	query := newDb.Table("schedules")

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}
	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query.Count(&res.MaxSchedules)
	query.Find(&schedules)

	for i := 0; i < len(schedules); i++ {
		res.Schedules = append(
			res.Schedules,
			NewSchedule(&schedules[i]),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchScheduleById(
	ctx context.Context,
	in *SeppoService.FetchScheduleByIdRequest,
) (
	*SeppoService.FetchScheduleByIdResponse,
	error,
) {
	res := &SeppoService.FetchScheduleByIdResponse{}
	newDb := s.getDB()
	defer newDb.Close()

	schedules := []db.Schedule{}
	newDb.Where("id in (?)", in.ScheduleIds).Find(&schedules)

	for i := 0; i < len(in.ScheduleIds); i++ {
		found := false
		for j := 0; j < len(schedules); j++ {
			if in.ScheduleIds[i] == schedules[j].ID {
				found = true
				res.Schedules = append(
					res.Schedules,
					NewSchedule(&schedules[j]),
				)
			}
		}
		if found == false {
			res.Schedules = append(
				res.Schedules,
				&SeppoService.Schedule{
					Id: 0,
				},
			)
		}
	}

	return res, nil
}
