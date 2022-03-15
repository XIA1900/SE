package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var communityManageServiceLock sync.Mutex
var communityManageService *CommunityManageService

type ICommunityManageService interface {
	//CreateCommunity(creator string, name string, description string, createTime string) error
	//GetCommunityByName(name string) (model.Community, model.User, error)
	//UpdateCommunity(updateInfo model.Community) error
	//DeleteCommunity(id int) error
}

type CommunityManageService struct {
	communityDAO       dao.ICommunityDAO
	communityMemberDAO dao.ICommunityMemberDAO
}

func NewCommunityManageService(communityDAO dao.ICommunityDAO, communityMemberDAO dao.ICommunityMemberDAO) *CommunityManageService {
	if communityManageService == nil {
		communityManageServiceLock.Lock()
		if communityManageService == nil {
			communityManageService = &CommunityManageService{
				communityDAO:       communityDAO,
				communityMemberDAO: communityMemberDAO,
			}
		}
		communityManageServiceLock.Unlock()
	}
	return communityManageService
}

var CommunityManageServiceSet = wire.NewSet(
	dao.NewCommunityMemberDAO,
	wire.Bind(new(dao.ICommunityMemberDAO), new(*dao.CommunityMemberDAO)),
	dao.NewCommunityDAO,
	wire.Bind(new(dao.ICommunityDAO), new(*dao.CommunityDAO)),
	NewCommunityManageService,
)
