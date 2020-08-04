package user

import (
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/security"
	"context"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

var (
	// DefaultPassword 默认密码.
	DefaultPassword = "123456"
)

// Service user服务
type Service struct {
	Repo IUserRepository
}

// Create 创建用户
func (s Service) Create(ctx context.Context, user *User) error {
	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return err
	}

	dbUser, err := s.Repo.FindByName(ctx, user.UserName)
	if err != nil {
		return err
	}

	if dbUser != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"userName": user.UserName,
			"dbUserID": dbUser.UserID,
		}, web.ErrConflict, "Create User, Duplicate userName")
	}

	user.Password = string(hashPassword)
	user.UserID = uuid.NewV4()
	user.NeedChangePassword = true

	return s.Repo.Create(ctx, user)
}

// Edit 编辑
// 增加userID 参数，让调用更明确此为必传值
func (s Service) Edit(ctx context.Context, userID uuid.UUID, user User) error {

	dbUser, err := s.Repo.FindByName(ctx, user.UserName)
	if err != nil {
		return err
	}

	if dbUser == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID": userID,
		}, web.ErrNotFound, "Edit, User Not Found")
	}

	if dbUser.UserID != userID {
		return web.WriteErrLogWithField(logrus.Fields{
			"userName": user.UserName,
			"dbUserID": dbUser.UserID,
			"userID":   userID,
		}, web.ErrConflict, "Edit User, userID Conflict")
	}

	dbUser.RealName = user.RealName
	dbUser.Roles = user.Roles
	dbUser.UserName = user.UserName

	return s.Repo.Update(ctx, dbUser)
}

// ChangeInitPassword 修改初始密码
func (s Service) ChangeInitPassword(ctx context.Context, userID uuid.UUID, password string) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if dbUser == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID": userID,
		}, web.ErrNotFound, "ChangeInitPassword, User Not Found")
	}

	if !dbUser.NeedChangePassword {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID":             dbUser.UserID,
			"userName":           dbUser.UserName,
			"needChangePassword": dbUser.NeedChangePassword,
		}, web.ErrNotAcceptable, "ChangeInitPassword, Cannot Change Init Password")
	}

	hashPassword, err := security.Hash(password)
	if err != nil {
		return err
	}

	dbUser.Password = string(hashPassword)
	dbUser.NeedChangePassword = false

	return s.Repo.Update(ctx, dbUser)
}

// ResetPassword 重置密码
func (s Service) ResetPassword(ctx context.Context, userID uuid.UUID) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if dbUser == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID": userID,
		}, web.ErrNotFound, "ResetPassword, User Not Found")
	}

	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return err
	}

	dbUser.Password = string(hashPassword)
	dbUser.NeedChangePassword = true

	return s.Repo.Update(ctx, dbUser)
}

// Lock 锁定
func (s Service) Lock(ctx context.Context, userID uuid.UUID) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if dbUser == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID": userID,
		}, web.ErrNotFound, "Lock, User Not Found")
	}

	if dbUser.IsLock {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID":   dbUser.UserID,
			"userName": dbUser.UserName,
			"Lock":     dbUser.IsLock,
		}, web.ErrNotAcceptable, "Lock, User is Locked, Cannot Lock")
	}

	dbUser.IsLock = true

	return s.Repo.Update(ctx, dbUser)
}

// Unlock 解锁
func (s Service) Unlock(ctx context.Context, userID uuid.UUID) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if dbUser == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID": userID,
		}, web.ErrNotFound, "Unlock, User Not Found")
	}

	if !dbUser.IsLock {
		return web.WriteErrLogWithField(logrus.Fields{
			"userID":   dbUser.UserID,
			"userName": dbUser.UserName,
			"Lock":     dbUser.IsLock,
		}, web.ErrNotAcceptable, "Unlock, User is Unlocked, Cannot Unlock")
	}

	dbUser.IsLock = false

	return s.Repo.Update(ctx, dbUser)
}

// Login 用户登陆
func (s Service) Login(ctx context.Context, userName string, password string) (*User, error) {
	dbUser, err := s.Repo.FindByName(ctx, userName)
	if err != nil {
		return nil, err
	}

	if dbUser == nil {
		return nil, web.WriteErrLogWithField(logrus.Fields{
			"userName": userName,
		}, web.ErrNotFound, "Login, User Not Found")
	}

	if dbUser.IsLock {
		return nil, web.WriteErrLogWithField(logrus.Fields{
			"userID":   dbUser.UserID,
			"userName": dbUser.UserName,
			"Lock":     dbUser.IsLock,
		}, web.ErrUserLocked, "Login, User is Locked, Cannot Login")
	}

	err = security.VerifyPassword(dbUser.Password, password)
	if err != nil {
		return nil, web.WriteErrLogWithField(logrus.Fields{
			"userID":   dbUser.UserID,
			"userName": dbUser.UserName,
			"password": password,
		}, web.ErrPasswordIncorrect, "Login, Password Incorrect, %s", err.Error())
	}

	return dbUser, nil
}
