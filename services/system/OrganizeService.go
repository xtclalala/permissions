package system

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	"permissions/model/common"
	"permissions/model/system"
)

type OrganizeService struct{}

var AppOrganizeService = new(OrganizeService)

// Register 注册组织
func (s *OrganizeService) Register(dto system.SysOrganize) (err error) {
	if errors.Is(s.CheckRepeat(dto.Pid, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新组织
func (s *OrganizeService) Update(dto system.SysOrganize) (err error) {
	var old system.SysOrganize
	err = global.Db.Where("id = ?", dto.ID).Find(&old).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.Pid != dto.Pid || old.Name != dto.Name {
		if errors.Is(s.CheckRepeat(dto.Pid, dto.Name), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}
	err = global.Db.Save(dto).Error
	return
}

// Search 搜索组织
func (s *OrganizeService) Search(dto SearchOrganize) (err error, list []system.SysOrganize, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system.SysOrganize{})
	var menus []system.SysOrganize

	if dto.Pid != 0 {
		db = db.Where("pid = ?", dto.Pid)
	}
	if dto.Name != "" {
		db = db.Where("name like ?", "%"+dto.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return err, menus, total
	}
	db = db.Limit(limit).Offset(offset)

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "sort", Raw: true},
		Desc:   dto.Desc,
	}

	err = db.Order(oc).Find(&list).Error
	return err, list, total
}

// CheckRepeat 检查 pid 和 name 是否存在
func (s *OrganizeService) CheckRepeat(pid uint, name string) (err error) {
	var temp system.SysOrganize
	err = global.Db.Where("pid = ? and name = ?", pid, name).First(&temp).Error
	return
}

// GetAll 查所有组织
func (s *OrganizeService) GetAll() (err error, dos []system.SysOrganize) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查组织
func (s *OrganizeService) GetById(id uint) (err error, do system.SysOrganize) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}

// GetOrgByUserId 根据 用户id 查组织
func (s *OrganizeService) GetOrgByUserId(userId uuid.UUID) (err error, orgs []system.SysOrganize) {
	rows, err := global.Db.Where(&system.M2mUserOrganize{}, userId).Rows()
	defer rows.Close()
	if err != nil {
		return err, orgs
	}
	for rows.Next() {
		var userOrg system.M2mUserOrganize
		global.Db.ScanRows(rows, &userOrg)
		_, org := s.GetById(userOrg.SysOrganizeId)
		orgs = append(orgs, org)
	}
	return
}

type SearchMenu struct {
	common.BasePage
	Menu
}

type Menu struct {
	Pid       uint   `json:"pid"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
}
