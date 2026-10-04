package storage

import (
	"time"
)

// IpSecurity 记录 IP 的信任状态：登录成功为 true，失败为 false
type IpSecurity struct {
	Ip       string `gorm:"column:ip;primaryKey;size:64"`
	Security bool   `gorm:"column:security"`
}

func (IpSecurity) TableName() string { return "ip_security" }

// Captcha 一次性验证码记录
type Captcha struct {
	Key      string    `gorm:"column:key;primaryKey;size:64"`
	Value    string    `gorm:"column:value;size:16"`
	ExpireAt time.Time `gorm:"column:expire_at"`
}

func (Captcha) TableName() string { return "captcha" }

func SetIpSecurity(ip string, security bool) {
	_ = db.Save(&IpSecurity{Ip: ip, Security: security}).Error
}

// GetIpSecurity 默认（无记录）为不安全状态，需要验证码
func GetIpSecurity(ip string) (result bool) {
	var row IpSecurity
	if err := db.Where("ip = ?", ip).First(&row).Error; err != nil {
		return false
	}
	return row.Security
}

func SetCaptcha(key string, value string, duration time.Duration) {
	_ = db.Save(&Captcha{
		Key:      key,
		Value:    value,
		ExpireAt: time.Now().Add(duration),
	}).Error
}

// ValidCaptcha 校验验证码并立即删除（一次性使用）
func ValidCaptcha(key string, target string) (result bool) {
	var row Captcha
	if err := db.Where("key = ?", key).First(&row).Error; err != nil {
		return false
	}
	_ = db.Where("key = ?", key).Delete(&Captcha{}).Error
	return row.Value == target && row.ExpireAt.After(time.Now())
}

// CleanExpiredCaptcha 清理已过期的验证码记录
func CleanExpiredCaptcha() {
	_ = db.Where("expire_at < ?", time.Now()).Delete(&Captcha{}).Error
}
