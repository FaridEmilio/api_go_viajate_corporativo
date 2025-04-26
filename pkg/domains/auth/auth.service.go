package auth

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/commons"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/util"
)

type service struct {
	repository               Repository
	util                     util.UtilService
	commons                  commons.Commons
	//firebaseRemoteRepository storage.FirebaseRemoteRepository
}

func NewUsuarioService(repo Repository, util util.UtilService, commons commons.Commons, firebaseRemoteRepo storage.FirebaseRemoteRepository) UsuarioService {
	return &service{
		repository:               repo,
		util:                     util,
		commons:                  commons,
		firebaseRemoteRepository: firebaseRemoteRepo,
	}
}
