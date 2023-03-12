// Package system 登录相关的Service
// @author: kbj
// @date: 2023/2/3
package system

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gookit/goutil/arrutil"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"react-admin-server/dao"
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
	if err := dao.User.SelectLoginUser(g.DbClient, &user, param.Username, param.Password); err != nil {
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
		"roles": l.GetRoleKeys(user), // 查询所有的角色信息
		"user": &system.LoginUserResponse{
			ID:       user.ID,
			CreateAt: user.CreateAt,
			UpdateAt: user.UpdateAt,
			Username: user.Username,
			Mobile:   user.Mobile,
			Gender:   user.Gender,
			Avatar:   user.Avatar,
		},
	}
	return r.Ok(ctx, r.Data(resp))
}

// Menus 用户菜单信息
func (l *LoginService) Menus(ctx *fiber.Ctx) error {
	user := g.LoginUser.User(ctx)
	roleKeys := l.GetRoleKeys(user)
	isAdmin := arrutil.Contains(*roleKeys, "admin")

	var menus []*domain.Menu
	if isAdmin {
		// 管理员角色查询全部
		g.DbClient.Model(domain.Menu{}).Where("menu_type <> 'B'").Find(&menus)
	} else {
		sql := `
			WITH RECURSIVE test AS (
				SELECT
					t1.* 
				FROM
					t_menu t1
					JOIN ( SELECT tt2.menu_id FROM t_user_role tt1 JOIN t_role_menu tt2 ON tt1.role_id = tt2.role_id WHERE tt1.user_id = ? ) t2 ON t2.menu_id = t1.id 
				WHERE
					t1.delete_at = 0 
					AND t1.menu_type <> 'B' 
					AND t1.enabled = TRUE UNION ALL
				SELECT
					t1.* 
				FROM
					t_menu t1, test 
				WHERE
					t1.parent_id = test.id 
					AND t1.delete_at = 0 
					AND t1.menu_type <> 'B' 
					AND t1.enabled = TRUE 
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
func (*LoginService) GetRoleKeys(user *domain.User) *[]string {
	var roles []*domain.Role
	err := g.DbClient.Model(user).Association("Roles").Find(&roles)
	if err != nil {
		g.Logger.Error("查询角色Key失败", zap.Error(err))
	}

	roleKeys := lo.Map[*domain.Role, string](roles, func(item *domain.Role, index int) string {
		return item.RoleKey
	})
	return &roleKeys
}
