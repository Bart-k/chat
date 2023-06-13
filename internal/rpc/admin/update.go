package admin

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/errs"
	"github.com/OpenIMSDK/chat/pkg/proto/admin"
	"time"
)

type Admin struct {
	Account    string    `gorm:"column:account;primary_key;type:char(64)"`
	Password   string    `gorm:"column:password;type:char(64)"`
	FaceURL    string    `gorm:"column:face_url;type:char(64)"`
	Nickname   string    `gorm:"column:nickname;type:char(64)"`
	UserID     string    `gorm:"column:user_id;type:char(64)"` //openIM userID
	Level      int32     `gorm:"column:level;default:1"  `
	CreateTime time.Time `gorm:"column:create_time"`
}

func ToDBAdminUpdate(req *admin.AdminUpdateInfoReq) (map[string]any, error) {
	update := make(map[string]any)
	if req.Account != nil {
		if req.Account.Value == "" {
			return nil, errs.ErrArgs.Wrap("account is empty")
		}
		update["account"] = req.Account.Value
	}
	if req.Password != nil {
		if req.Password.Value == "" {
			return nil, errs.ErrArgs.Wrap("password is empty")
		}
		update["password"] = req.Password.Value
	}
	if req.FaceURL != nil {
		update["face_url"] = req.FaceURL.Value
	}
	if req.Nickname != nil {
		if req.Nickname.Value != "" {
			return nil, errs.ErrArgs.Wrap("nickname is empty")
		}
		update["nickname"] = req.Nickname.Value
	}
	if req.UserID != nil {
		update["user_id"] = req.UserID.Value
	}
	if req.Level != nil {
		update["level"] = req.Level.Value
	}
	if len(update) == 0 {
		return nil, errs.ErrArgs.Wrap("no update info")
	}
	return update, nil
}

func ToDBAdminUpdatePassword(password string) (map[string]any, error) {
	if password == "" {
		return nil, errs.ErrArgs.Wrap("password is empty")
	}
	return map[string]any{"password": password}, nil
}

/*
	Name       *wrapperspb.StringValue
	AppID      *wrapperspb.StringValue
	Icon       *wrapperspb.StringValue
	Url        *wrapperspb.StringValue
	Md5        *wrapperspb.StringValue
	Size       *wrapperspb.Int64Value
	Version    *wrapperspb.StringValue
	Priority   *wrapperspb.UInt32Value
	Status     *wrapperspb.UInt32Value
	CreateTime *wrapperspb.Int64Value
*/
/*
type Applet struct {
	ID         string    `gorm:"column:id;primary_key;size:64"`
	Name       string    `gorm:"column:name;size:64"`
	AppID      string    `gorm:"column:app_id;uniqueIndex;size:255"`
	Icon       string    `gorm:"column:icon;size:255"`
	URL        string    `gorm:"column:url;size:255"`
	MD5        string    `gorm:"column:md5;size:255"`
	Size       int64     `gorm:"column:size"`
	Version    string    `gorm:"column:version;size:64"`
	Priority   uint32    `gorm:"column:priority;size:64"`
	Status     uint8     `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime;size:64"`
}
*/

func ToDBAppletUpdate(req *admin.UpdateAppletReq) (map[string]any, error) {
	update := make(map[string]any)
	if req.Name != nil {
		if req.Name.Value == "" {
			return nil, errs.ErrArgs.Wrap("name is empty")
		}
		update["name"] = req.Name.Value
	}
	if req.AppID != nil {
		if req.AppID.Value == "" {
			return nil, errs.ErrArgs.Wrap("appID is empty")
		}
		update["app_id"] = req.AppID.Value
	}
	if req.Icon != nil {
		update["icon"] = req.Icon.Value
	}
	if req.Url != nil {
		if req.Url.Value == "" {
			return nil, errs.ErrArgs.Wrap("url is empty")
		}
		update["url"] = req.Url.Value
	}
	if req.Md5 != nil {
		if hash, _ := hex.DecodeString(req.Md5.Value); len(hash) != md5.Size {
			return nil, errs.ErrArgs.Wrap("md5 is invalid")
		}
		update["md5"] = req.Md5.Value
	}
	if req.Size != nil {
		if req.Size.Value <= 0 {
			return nil, errs.ErrArgs.Wrap("size is invalid")
		}
		update["size"] = req.Size.Value
	}
	if req.Version != nil {
		if req.Version.Value == "" {
			return nil, errs.ErrArgs.Wrap("version is empty")
		}
		update["version"] = req.Version.Value
	}
	if req.Priority != nil {
		update["priority"] = req.Priority.Value
	}
	if req.Status != nil {
		update["status"] = req.Status.Value
	}
	if len(update) == 0 {
		return nil, errs.ErrArgs.Wrap("no update info")
	}
	return update, nil
}

func ToDBInvitationRegisterUpdate(userID string) map[string]any {
	return map[string]any{"user_id": userID}
}