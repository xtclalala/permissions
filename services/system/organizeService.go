package system

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	system2 "permissions/model/system"
)

type OrganizeService struct{}

var AppOrganizeService = new(OrganizeService)

// Register 注册组织
func (s *OrganizeService) Register(dto *system2.SysOrganize) (err error) {
	if errors.Is(s.CheckRepeat(dto.Pid, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新组织
func (s *OrganizeService) Update(dto *system2.SysOrganize) (err error) {
	var old system2.SysOrganize
	err = global.Db.First(&old, dto.ID).Error
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
func (s *OrganizeService) Search(dto *system2.SearchOrganize) (err error, list []system2.SysOrganize, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system2.SysOrganize{})
	var menus []system2.SysOrganize

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
func (s *OrganizeService) CheckRepeat(pid int, name string) (err error) {
	var total int64
	global.Db.Model(&system2.SysOrganize{}).Where(&system2.SysOrganize{Pid: pid, Name: name}).Count(&total)
	if total != 0 {
		err = gorm.ErrRecordNotFound
	} else {
		err = nil
	}
	return
}

// GetAll 查所有组织
func (s *OrganizeService) GetAll() (err error, dos []system2.SysOrganize) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查组织
func (s *OrganizeService) GetById(id int) (err error, do system2.SysOrganize) {
	err = global.Db.First(&do, id).Error
	return
}

// GetOrgByUserId 根据 用户id 查组织
func (s *OrganizeService) GetOrgByUserId(userId uuid.UUID) (err error, orgs []system2.SysOrganize) {
	rows, err := global.Db.Model(&system2.M2mUserOrganize{}).Where(&system2.M2mUserOrganize{SysUserId: userId}).Rows()
	defer rows.Close()
	if err != nil {
		return err, orgs
	}
	for rows.Next() {
		var userOrg system2.M2mUserOrganize
		global.Db.ScanRows(rows, &userOrg)
		_, org := s.GetById(userOrg.SysOrganizeId)
		orgs = append(orgs, org)
	}
	return
}

func (s *OrganizeService) DeleteOrganize(id int) (err error) {
	err = global.Db.Delete(&system2.SysOrganize{}, id).Error
	return
}
