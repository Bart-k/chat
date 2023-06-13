package chat

import (
	"context"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/discoveryregistry"
	"github.com/OpenIMSDK/chat/pkg/common/config"
	"github.com/OpenIMSDK/chat/pkg/common/mctx"
	"github.com/OpenIMSDK/chat/pkg/eerrs"
	"github.com/OpenIMSDK/chat/pkg/proto/admin"
)

func NewAdminClient(zk discoveryregistry.SvcDiscoveryRegistry) *AdminClient {
	return &AdminClient{
		zk: zk,
	}
}

type AdminClient struct {
	zk discoveryregistry.SvcDiscoveryRegistry
}

func (o *AdminClient) client(ctx context.Context) (admin.AdminClient, error) {
	conn, err := o.zk.GetConn(ctx, config.Config.RpcRegisterName.OpenImAdminName)
	if err != nil {
		return nil, err
	}
	return admin.NewAdminClient(conn), nil
}

func (o *AdminClient) GetConfig(ctx context.Context) (map[string]string, error) {
	client, err := o.client(ctx)
	if err != nil {
		return nil, err
	}
	conf, err := client.GetClientConfig(ctx, &admin.GetClientConfigReq{})
	if err != nil {
		return nil, err
	}
	if conf.Config == nil {
		return map[string]string{}, nil
	}
	return conf.Config, nil
}

func (o *AdminClient) CheckInvitationCode(ctx context.Context, invitationCode string) error {
	client, err := o.client(ctx)
	if err != nil {
		return err
	}
	resp, err := client.FindInvitationCode(ctx, &admin.FindInvitationCodeReq{Codes: []string{invitationCode}})
	if err != nil {
		return err
	}
	if len(resp.Codes) == 0 {
		return eerrs.ErrInvitationNotFound.Wrap()
	}
	if resp.Codes[0].UsedUserID != "" {
		return eerrs.ErrInvitationCodeUsed.Wrap()
	}
	return nil
}

func (o *AdminClient) CheckRegister(ctx context.Context, ip string) error {
	client, err := o.client(ctx)
	if err != nil {
		return err
	}
	_, err = client.CheckRegisterForbidden(ctx, &admin.CheckRegisterForbiddenReq{Ip: ip})
	return err
}

func (o *AdminClient) CheckLogin(ctx context.Context, userID string, ip string) error {
	client, err := o.client(ctx)
	if err != nil {
		return err
	}
	_, err = client.CheckLoginForbidden(ctx, &admin.CheckLoginForbiddenReq{Ip: ip, UserID: userID})
	return err
}

func (o *AdminClient) UseInvitationCode(ctx context.Context, userID string, invitationCode string) error {
	client, err := o.client(ctx)
	if err != nil {
		return err
	}
	_, err = client.UseInvitationCode(ctx, &admin.UseInvitationCodeReq{UserID: userID, Code: invitationCode})
	return err
}

func (o *AdminClient) CheckNilOrAdmin(ctx context.Context) (bool, error) {
	if !mctx.HaveOpUser(ctx) {
		return false, nil
	}
	_, err := mctx.CheckAdmin(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *AdminClient) CreateToken(ctx context.Context, userID string, userType int32) (*admin.CreateTokenResp, error) {
	client, err := o.client(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateToken(ctx, &admin.CreateTokenReq{UserID: userID, UserType: userType})
}

func (o *AdminClient) GetDefaultFriendUserID(ctx context.Context) ([]string, error) {
	client, err := o.client(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.FindDefaultFriend(ctx, &admin.FindDefaultFriendReq{})
	if err != nil {
		return nil, err
	}
	return resp.UserIDs, nil
}

func (o *AdminClient) GetDefaultGroupID(ctx context.Context) ([]string, error) {
	client, err := o.client(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.FindDefaultGroup(ctx, &admin.FindDefaultGroupReq{})
	if err != nil {
		return nil, err
	}
	return resp.GroupIDs, nil
}