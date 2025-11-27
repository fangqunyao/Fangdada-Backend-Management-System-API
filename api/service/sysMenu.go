//菜单服务层

package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

// ISysMenuService 接口定义
type ISysMenuService interface {
	QuerySysMenuVoList(c *gin.Context)
	CreateSysMenu(c *gin.Context, menu entity.SysMenu)
	GetSysMenu(c *gin.Context, Id int)
	UpdateSysMenu(c *gin.Context, menu entity.SysMenu)
	DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto)
	GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string)
	GetSysMenuTree(c *gin.Context, MenuName string, MenuStatus string)
}

type SysMenuServiceImpl struct{}

// 新增菜单
func (s SysMenuServiceImpl) CreateSysMenu(c *gin.Context, sysMenu entity.SysMenu) {
	bool := dao.CreateSysMenu(sysMenu)
	if !bool {
		result.Failed(c, int(result.ApiCode.MENUISEXIST),
			result.ApiCode.GetMessage(result.ApiCode.MENUISEXIST))
		return
	}
	result.Success(c, true)
}

// QuerySysMenuVoList 查询新增选项列表
func (s SysMenuServiceImpl) QuerySysMenuVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysMenuVoList())
}

// GetSysMenu 获取详情
func (s SysMenuServiceImpl) GetSysMenu(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysMenu(Id))
}

// UpdateSysMenu 修改菜单
func (s SysMenuServiceImpl) UpdateSysMenu(c *gin.Context, menu entity.SysMenu) {
	result.Success(c, dao.UpdateSysMenu(menu))
}

// DeleteSysMenu 删除菜单
func (s SysMenuServiceImpl) DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto) {
	bool := dao.DeleteSysMenu(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.DELSYSMENUFAILED),
			result.ApiCode.GetMessage(result.ApiCode.DELSYSMENUFAILED))
		return
	}
	result.Success(c, true)
}

// GetSysMenuList 查询菜单列表（扁平）
func (s SysMenuServiceImpl) GetSysMenuList(c *gin.Context, menuName, menuStatus string) {
	menus, err := dao.GetSysMenuList(menuName, menuStatus)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "查询菜单失败")
		return
	}
	result.Success(c, menus)
}

// GetSysMenuTree 查询菜单树
// func (s SysMenuServiceImpl) GetSysMenuTree(c *gin.Context, menuName, menuStatus string) {
// 	// 1. 调用 DAO 获取扁平菜单
// 	flatMenus, err := dao.GetSysMenuList(menuName, menuStatus)
// 	if err != nil {
// 		result.Failed(c, int(result.ApiCode.FAILED), "查询菜单失败")
// 		return
// 	}
// 	// 2. 构建树形结构
// 	tree := buildMenuTree(flatMenus)
// 	// 3. 返回结果
// 	result.Success(c, tree)
// }

func (s *SysMenuServiceImpl) GetSysMenuTree(c *gin.Context, menuName, menuStatus string) {
	flatMenus, err := dao.GetSysMenuList(menuName, menuStatus)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "查询菜单失败")
		return
	}

	tree := entity.SysMenuList(flatMenus).BuildTree()

	result.Success(c, tree)
}

// buildMenuTree 将扁平菜单列表构建成树形结构
// func buildMenuTree(flatMenus []*entity.SysMenu) []*entity.SysMenu {
// 	if len(flatMenus) == 0 {
// 		return nil
// 	}

// 	// 创建 ID 到菜单的映射，并初始化 Children
// 	menuMap := make(map[uint]*entity.SysMenu, len(flatMenus))
// 	for _, menu := range flatMenus {
// 		menu.Children = []*entity.SysMenu{} // 避免 nil 切片
// 		menuMap[menu.ID] = menu
// 	}

// 	// 构建父子关系并收集根节点
// 	var rootMenus []*entity.SysMenu
// 	for _, menu := range flatMenus {
// 		if menu.ParentId == 0 {
// 			// 明确约定：ParentId = 0 表示根节点
// 			rootMenus = append(rootMenus, menu)
// 		} else if parent, exists := menuMap[menu.ParentId]; exists {
// 			parent.Children = append(parent.Children, menu)
// 		}
// 		// 注意：如果 ParentId 非 0 但父节点不存在，该菜单会被忽略（或可记录日志）
// 	}

// 	return rootMenus
// }

var sysMenuService = SysMenuServiceImpl{}

func SysMenuService() ISysMenuService {
	return &sysMenuService
}
