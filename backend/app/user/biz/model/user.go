package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"unique;unique_index"`
	PasswordHashed string `json:"passwordHashed" gorm:"type:varchar(255) not null"`
	Username       string `json:"username"`
	Avator         string `json:"avator"`
}

func (User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func GetByEmailWithCache(db *gorm.DB, redisCli *redis.Client, ctx context.Context, email string) (user *User, err error) {
	cachedKey := fmt.Sprintf("user_%v", email)
	cachedResult := redisCli.Get(ctx, cachedKey)
	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		bytes, err1 := cachedResult.Bytes()
		if err1 != nil {
			return err1
		}
		err1 = json.Unmarshal(bytes, &user)
		if err1 != nil {
			return err1
		}
		return nil
	}()
	if err == nil {
		return
	}
	// 缓存穿透
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func GetById(db *gorm.DB, ctx context.Context, user_id int64) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id = ? ", user_id).First(&user).Error
	return
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func UpdateUser(db *gorm.DB, ctx context.Context, user *User) error {
	updates := map[string]interface{}{}
	f := false
	if user.Username != "" {
		updates["username"] = user.Username
		f = true
	}
	if user.Avator != "" {
		updates["avator"] = user.Avator
		f = true
	}
	if user.PasswordHashed != "" {
		updates["password_hashed"] = user.PasswordHashed
		f = true
	}
	if !f {
		return fmt.Errorf("no update")
	}
	updates["updated_at"] = gorm.Expr("now()")
	return db.WithContext(ctx).
		Model(&User{}).
		Where("id = ? and deleted_at IS NULL", user.ID).
		Updates(updates).Error
}

func DeleteUser(db *gorm.DB, ctx context.Context, userid int64) error {
	return db.WithContext(ctx).Where("id = ?", userid).Unscoped().Delete(&User{}).Error
}

func QueryUser(db *gorm.DB, ctx context.Context, userid int64) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id = ?", userid).First(&user).Error
	return
}

// func QueryUser(db *gorm.DB, ctx context.Context, userid int64) (user *User, err error)  {
// 	db := DB.Model(model.User{})
// 	if keyword != nil && len(*keyword) != 0 {
// 		db = db.Where(DB.Or("name like ?", "%"+*keyword+"%").
// 			Or("introduce like ?", "%"+*keyword+"%"))
// 	}
// 	var total int64
// 	if err := db.Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	var res []*model.User
// 	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	return res, total, nil

// }
