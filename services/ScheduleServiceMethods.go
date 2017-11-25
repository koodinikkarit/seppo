package services

import (
	"time"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/managers"
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
	newDB := s.getGormDB()
	defer newDB.Close()

	d1 := time.Unix(in.Start, 0)
	d2 := time.Unix(in.End, 0)

	newSchedule := db.Schedule{
		Name:  in.Name,
		Start: &d1,
		End:   &d2,
	}
	newDB.Create(&newSchedule)
	res.Schedule = generators.NewSchedule(&newSchedule)

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
	tx := s.getGormDB()
	defer tx.Close()

	var schedule db.Schedule
	tx.First(&schedule, in.ScheduleId)

	if schedule.ID == 0 {
		res.Success = true
		return res, nil
	}

	if in.Name != "" {
		schedule.Name = in.Name
	}

	tx.Save(&schedule)

	if len(in.AddSongIds) > 0 {
		var newScheduleVariations []db.ScheduleVariation
		for _, newSongID := range in.AddSongIds {
			newScheduleVariations = append(
				newScheduleVariations,
				db.ScheduleVariation{
					VariationID: newSongID,
					ScheduleID:  schedule.ID,
				},
			)
		}
		managers.BatchCreateScheduleVariations(
			tx,
			newScheduleVariations,
		)
	}

	if len(in.RemoveSongIds) > 0 {
		tx.Where("schedule_variations.variation_id in (?)", in.RemoveSongIds).
			Delete(&db.ScheduleVariation{})
	}
	res.Schedule = generators.NewSchedule(&schedule)
	res.Success = true
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
	newDB := s.getGormDB()
	defer newDB.Close()

	var schedule db.Schedule
	newDB.First(&schedule, in.ScheduleId)

	if schedule.ID == 0 {
		res.Success = false
		return res, nil
	}

	newDB.Delete(&schedule)
	res.Success = true
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
	newDB := s.getGormDB()
	defer newDB.Close()

	schedules := []db.Schedule{}

	query := newDB.Table("schedules")

	if in.Offset > 0 {
		query = query.Offset(in.Offset)
	}
	if in.Limit > 0 {
		query = query.Limit(in.Limit)
	}

	query.Count(&res.MaxSchedules)
	query.Find(&schedules)

	for _, schedule := range schedules {
		res.Schedules = append(
			res.Schedules,
			generators.NewSchedule(&schedule),
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
	newDB := s.getGormDB()
	defer newDB.Close()

	var schedules []db.Schedule

	newDB.Where("id in (?)", in.ScheduleIds).
		Find(&schedules)

	for _, scheduleID := range in.ScheduleIds {
		found := false
		for _, schedule := range schedules {
			if schedule.ID != scheduleID {
				continue
			}
			found = true
			res.Schedules = append(
				res.Schedules,
				generators.NewSchedule(&schedule),
			)
		}
		if found == false {
			res.Schedules = append(
				res.Schedules,
				&SeppoService.Schedule{},
			)
		}
	}

	return res, nil
}
