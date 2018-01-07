package service

import (
	"golang.org/x/net/context"

	SeppoService "github.com/koodinikkarit/go-clientlibs/seppo"
	"github.com/koodinikkarit/seppo/generators"
	"github.com/koodinikkarit/seppo/models"
)

func (s *SeppoServiceServer) FetchMatiasClient(
	ctx context.Context,
	in *SeppoService.FetchMatiasClientRequest,
) (
	*SeppoService.FetchMatiasClientResponse,
	error,
) {
	res := &SeppoService.FetchMatiasClientResponse{}
	db := s.getDB()
	var matiasClients []models.MatiasClient

	db.Where("id in (?)", in.MatiasClientIds).Find(&matiasClients)

	for _, matiasClientId := range in.MatiasClientIds {
		found := false
		for _, matiasClient := range matiasClients {
			if matiasClient.ID != matiasClientId {
				continue
			}
			found = true
			res.MatiasClients = append(
				res.MatiasClients,
				generators.NewMatiasClient(&matiasClient),
			)
		}
		if found == false {
			res.MatiasClients = append(
				res.MatiasClients,
				&SeppoService.MatiasClient{},
			)
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) SearchMatiasClients(
	ctx context.Context,
	in *SeppoService.SearchMatiasClientsRequest,
) (
	*SeppoService.SearchMatiasClientsResponse,
	error,
) {
	res := &SeppoService.SearchMatiasClientsResponse{}
	db := s.getDB()
	var matiasClients []models.MatiasClient
	q := db.Find(&matiasClients)

	q.Count(&res.MaxMatiasClients)

	for _, matiasClient := range matiasClients {
		res.MatiasClients = append(
			res.MatiasClients,
			generators.NewMatiasClient(&matiasClient),
		)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchMatiasClientEwDatabases(
	ctx context.Context,
	in *SeppoService.FetchMatiasClientEwDatabasesRequest,
) (
	*SeppoService.FetchMatiasClientEwDatabasesResponse,
	error,
) {
	res := &SeppoService.FetchMatiasClientEwDatabasesResponse{}
	db := s.getDB()

	var ewDatabases []models.EwDatabase
	db.Where("matias_client_id in (?)", in.MatiasClientIds).Find(&ewDatabases)

	res.EwDatabases = make(map[uint32]*SeppoService.FetchMatiasClientEwDatabasesResponse_EwDatabases)

	for _, matiasClientId := range in.MatiasClientIds {
		res.EwDatabases[matiasClientId] = &SeppoService.FetchMatiasClientEwDatabasesResponse_EwDatabases{}
		for _, ewDatabase := range ewDatabases {
			if ewDatabase.MatiasClientID == matiasClientId {
				res.EwDatabases[matiasClientId].EwDatabases = append(
					res.EwDatabases[matiasClientId].EwDatabases,
					generators.NewEwDatabase(&ewDatabase),
				)
			}
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) UpdateMatiasClient(
	ctx context.Context,
	in *SeppoService.UpdateMatiasClientRequest,
) (
	*SeppoService.UpdateMatiasClientResponse,
	error,
) {
	res := &SeppoService.UpdateMatiasClientResponse{}
	db := s.getDB()

	var matiasClient models.MatiasClient
	db.First(&matiasClient, in.MatiasClientId)

	if matiasClient.ID == 0 {
		res.Success = true
		return res, nil
	}

	if in.Name != "" {
		matiasClient.Name = in.Name
	}

	switch in.State {
	case SeppoService.MatiasClientAcceptedState_ACEPTED:
		matiasClient.Accepted = true
	case SeppoService.MatiasClientAcceptedState_DECLINED:
		matiasClient.Accepted = false
	}

	db.Save(&matiasClient)

	res.Success = true
	res.MatiasClient = generators.NewMatiasClient(&matiasClient)

	return res, nil
}
