// Package common 登录相关的Service
// @author: kbj
// @date: 2023/2/3
package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"react-admin-server/entity/domain"
	"react-admin-server/entity/vo/system"
	"react-admin-server/global/consts"
	"react-admin-server/global/g"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
	"time"
)

type LoginService struct {
}

// Login 登录的Service
func (l *LoginService) Login(ctx *fiber.Ctx, param *system.LoginRequest) error {
	// MD5密码
	param.Password = tool.Md5Encode(param.Username+param.Password, 512)

	// 查询对应用户信息
	var user domain.User
	if err := g.DbClient.Where("username = ? and password = ? and enabled = '1'", param.Username, param.Password).First(&user).Error; err != nil {
		_ = tool.LogDbError(err)
		return consts.NewServiceError("用户名或密码错误")
	}

	// 生成Token
	userResp := system.LoginUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
	token, err := l.generateJwtToken(&userResp)
	if err != nil {
		return err
	}

	// 返回登录成功
	ctx.Set(fiber.HeaderAuthorization, token)
	return r.Ok(ctx, r.Msg("登录成功"))
}

// Info 查询登录用户信息
func (l *LoginService) Info(ctx *fiber.Ctx) error {
	user := g.LoginUser.User(ctx)
	resp := map[string]any{
		"roles": l.GetRoleKeys(user.ID), // 查询所有的角色信息
		"user": &system.LoginUserResponse{
			ID:       user.ID,
			CreateAt: user.CreateAt,
			UpdateAt: user.UpdateAt,
			Username: user.Username,
			Mobile:   user.Mobile,
			Gender:   user.Gender,
			Avatar:   user.Avatar,
			DeptId:   user.DeptId,
			NickName: user.NickName,
			Email:    user.Email,
		},
		"permissions": l.GetPermissions(user),
	}
	return r.Ok(ctx, r.Data(resp))
}

// Menus 用户菜单信息
func (l *LoginService) Menus(ctx *fiber.Ctx) error {
	user := g.LoginUser.User(ctx)

	var menus []*domain.Menu
	if user.IsAdmin() {
		// 管理员角色查询全部
		g.DbClient.Model(domain.Menu{}).Where("menu_type <> 'B'").Find(&menus)
	} else {
		sql := `
			WITH RECURSIVE test AS (
				SELECT
					t1.* 
				FROM
					t_menu t1
					JOIN ( SELECT DISTINCT tt2.menu_id FROM t_user_role tt1 JOIN t_role_menu tt2 ON tt1.role_id = tt2.role_id WHERE tt1.user_id = ? ) t2 ON t2.menu_id = t1.id 
				WHERE
					t1.delete_at = 0 
					AND t1.menu_type <> 'B' 
					AND t1.enabled = '1' UNION ALL
				SELECT
					t1.* 
				FROM
					t_menu t1, test 
				WHERE
					t1.parent_id = test.id 
					AND t1.delete_at = 0 
					AND t1.menu_type <> 'B' 
					AND t1.enabled = '1' 
			)
			SELECT * FROM test ORDER BY test.order_num
		`

		g.DbClient.Raw(sql, user.ID).Scan(&menus)
	}

	return r.Ok(ctx, r.Data(menus))
}

// RefreshToken 刷新Token
func (l *LoginService) RefreshToken(id uint) string {
	// 查询用户信息
	var user domain.User
	if err := g.DbClient.Find(&user, id).Error; err != nil {
		g.Logger.Error("刷新用户Token失败", zap.Error(err))
		return ""
	}

	userResp := system.LoginUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	// 生成Token
	token, _ := l.generateJwtToken(&userResp)
	return token
}

// 生成Token
func (*LoginService) generateJwtToken(users *system.LoginUserResponse) (string, error) {
	// 生成Token
	claims := jwt.MapClaims{
		"name": g.Env.Jwt.Name,
		"exp":  time.Now().Add(time.Minute * time.Duration(g.Env.Jwt.Expire)).Unix(),
		"user": users,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(g.Env.Jwt.Key))
	if err != nil {
		g.Logger.Error("Token生成失败", zap.Error(err))
		return "", consts.NewServiceError("Token生成失败")
	}
	return t, nil
}

// GetRoleKeys 查询用户的角色Key
func (*LoginService) GetRoleKeys(userId uint) *[]string {
	var roleKeys []string
	err := g.DbClient.Model(&domain.Role{}).Distinct("t_role.role_key").
		Joins("join t_user_role on t_user_role.role_id = t_role.id and t_user_role.user_id = ?", userId).
		Where("t_role.enabled = '1'").Scan(&roleKeys).Error
	if err != nil {
		g.Logger.Error("查询角色Key失败", zap.Error(err))
	}
	return &roleKeys
}

// RolesList 全体角色列表
func (*LoginService) RolesList(ctx *fiber.Ctx) error {
	var list []domain.Role
	if err := tool.LogDbError(g.DbClient.Where("enabled = '1'").Order("order_num").Find(&list).Error); err != nil {
		return consts.NewServiceError("查询失败")
	}
	return r.Ok(ctx, r.Data(&list))
}

// GetPermissions 查询用户的权限字符
func (*LoginService) GetPermissions(user *domain.User) *[]string {
	var permissions []string
	if user.IsAdmin() {
		permissions = append(permissions, "*:*:*")
	} else {
		// 非管理员查询权限字符
		err := g.DbClient.Model(&domain.Menu{}).Distinct("t_menu.permission_flag").
			Joins("join t_role_menu on t_role_menu.menu_id = t_menu.id").
			Joins("join t_role on t_role.id = t_role_menu.role_id and t_role.enabled = '1' and t_role.delete_at = 0").
			Joins("join t_user_role on t_user_role.role_id = t_role.id and t_user_role.user_id = ?", user.ID).
			Where("t_menu.enabled = '1'").Scan(&permissions).Error
		_ = tool.LogDbError(err)
	}
	return &permissions
}
