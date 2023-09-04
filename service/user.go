package service

import (
	m "WheelChair-tiktok/model"
	"WheelChair-tiktok/utils"
	"errors"
	"gorm.io/gorm"
)

func AddUser(u m.User) (userInter m.User, err error) {
	// 判断用户是否注册
	if !errors.Is(m.DB.Where("user_name = ?", u.UserName).First(&userInter).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("user has already registered")
	}

	// 否则 hash密码 注册
	u.Password, err = utils.Encrypt(u.Password)
	if err != nil {
		return userInter, err
	}

	err = m.DB.Create(&u).Error
	return u, err
}

// Login 注册
func Login(u m.User) (userInter m.User, err error) {
	//判断用户是否为i
	err = m.DB.Where("ID = ?", uint(1)).First(&userInter).Error
	if err == nil {
		return userInter, errors.New("reserve user")
	}
	// 检测用户是否在数据库中
	err = m.DB.Where("user_name = ?", u.UserName).First(&userInter).Error

	// 处理查询错误
	if err != nil {
		// 用户不存在
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userInter, errors.New("the user is not exist")
		}
		return userInter, err
	}

	// 验证密码是否正确
	if ok := utils.BcryptCheck(u.Password, userInter.Password); !ok {
		return userInter, errors.New("incorrect password")
	}

	// 登录成功
	return userInter, nil
}

func GetUserInfo(id uint) (userInter m.User, err error) {
	// 根据id 检测用户是否存在数据库中
	err = m.DB.Where("ID = ?", id).First(&userInter).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userInter, errors.New("the user is not exist")
		}
		return userInter, err
	}
	return userInter, err
}

func IsFollowing(followerID uint, followedID uint) bool {
	var follow m.Follow
	err := m.DB.Where("follower_id = ? AND followed_id = ?", followerID, followedID).First(&follow).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return false
	}
	return true
}
