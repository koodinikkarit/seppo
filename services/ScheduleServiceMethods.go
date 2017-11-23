package services

import (
	"time"

	null "gopkg.in/volatiletech/null.v6"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
	"github.com/volatiletech/sqlboiler/queries/qm"
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

	newSchedule := models.Schedule{
		Name:  in.Name,
		Start: null.NewTime(d1, true),
		End:   null.NewTime(d2, true),
	}
	newSchedule.Insert(newDb)
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
	newDb := s.getDB()
	defer newDb.Close()
	tx, _ := newDb.Begin()

	schedule, _ := models.FindSchedule(
		newDb,
		in.ScheduleId,
	)

	if schedule == nil {
		res.Success = true
		return res, nil
	}

	if in.Name != "" {
		schedule.Name = in.Name
	}

	if len(in.AddSongIds) > 0 {
		for _, newSongID := range in.AddSongIds {
			newScheduleVariation := models.ScheduleVariation{
				VariationID: newSongID,
			}

			schedule.AddScheduleVariations(
				tx,
				true,
				&newScheduleVariation,
			)
		}
	}

	if len(in.RemoveSongIds) > 0 {
		schedule.ScheduleVariations(
			tx,
			qm.WhereIn("schedule_variations.variation_id in ?", in.RemoveSongIds),
		).DeleteAll()
	}
	res.Schedule = generators.NewSchedule(schedule)
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
	newDb := s.getDB()
	defer newDb.Close()

	schedule, _ := models.FindSchedule(
		newDb,
		in.ScheduleId,
	)

	if schedule == nil {
		res.Success = false
		return res, nil
	}

	schedule.DeletedAt = null.NewTime(time.Now(), true)
	schedule.Update(newDb)
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
	newDb := s.getDB()
	defer newDb.Close()

	var queryMods []qm.QueryMod

	if in.Offset > 0 {
		queryMods = append(
			queryMods,
			qm.Offset(int(in.Offset)),
		)
	}

	if in.Limit > 0 {
		queryMods = append(
			queryMods,
			qm.Limit(int(in.Limit)),
		)
	}

	c, _ := models.Schedules(
		newDb,
		queryMods...,
	).Count()
	res.MaxSchedules = uint64(c)

	schedules, _ := models.Schedules(
		newDb,
		queryMods...,
	).All()

	for _, schedule := range schedules {
		res.Schedules = append(
			res.Schedules,
			generators.NewSchedule(schedule),
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

	schedules, _ := models.Schedules(
		newDb,
		qm.WhereIn("id in ?", in.ScheduleIds),
	).All()

	for _, scheduleID := range in.ScheduleIds {
		found := false
		for _, schedule := range schedules {
			if schedule.ID != scheduleID {
				continue
			}
			found = true
			res.Schedules = append(
				res.Schedules,
				generators.NewSchedule(schedule),
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
