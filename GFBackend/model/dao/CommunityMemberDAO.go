package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var communityMemberDAOLock sync.Mutex
var communityMemberDAO *CommunityMemberDAO

type ICommunityMemberDAO interface {
	Create(communityID int, member string) error
	DeleteByCommunityID(id int) error
	DeleteByMember(member string) error
	GetCommunityIDsByMember(member string) ([]int, error)
	GetMembersByCommunityIDs(id int) ([]string, error)
}

type CommunityMemberDAO struct {
	db *gorm.DB
}

func NewCommunityMemberDAO() *CommunityMemberDAO {
	if communityMemberDAO == nil {
		communityMemberDAOLock.Lock()
		if communityMemberDAO == nil {
			communityMemberDAO = &CommunityMemberDAO{
				db: model.NewDB(),
			}
		}
		communityMemberDAOLock.Unlock()
	}
	return communityMemberDAO
}

func (communityMemberDAO *CommunityMemberDAO) Create(communityID int, member string) error {
	return nil
}

func (communityMemberDAO *CommunityMemberDAO) DeleteByCommunityID(id int) error {
	return nil
}

func (communityMemberDAO *CommunityMemberDAO) DeleteByMember(member string) error {
	return nil
}

func (communityMemberDAO *CommunityMemberDAO) GetCommunityIDsByMember(member string) ([]int, error) {
	return nil, nil
}

func (communityMemberDAO *CommunityMemberDAO) GetMembersByCommunityIDs(id int) ([]string, error) {
	return nil, nil
}
