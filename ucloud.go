package main

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/pathx"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

var (
	cfg        ucloud.Config
	credential auth.Credential
)

func initUcloud() {
	cfg = ucloud.NewConfig()
	cfg.Region = "cn-bj2"
	cfg.ProjectId = config.Ucloud.ProjectID

	// replace the public/private key by your own
	credential = auth.NewCredential()
	credential.PrivateKey = config.Ucloud.PrivateKey
	credential.PublicKey = config.Ucloud.PublicKey
}

func newGlobalSSH(ip, area string) (domain, instance string, err error) {
	pathxClient := pathx.NewClient(&cfg, &credential)
	req := pathxClient.NewCreateGlobalSSHInstanceRequest()
	req.Port = ucloud.Int(22)
	req.TargetIP = ucloud.String(ip)
	req.Area = ucloud.String(area)
	req.Remark = ucloud.String(fmt.Sprintf("git_ssh_%s_%s", area, ip))
	resp, err := pathxClient.CreateGlobalSSHInstance(req)
	return resp.AcceleratingDomain, resp.InstanceId, err
}

func deleteGlobalSSH(instanceID string) {
	pathxClient := pathx.NewClient(&cfg, &credential)
	req := pathxClient.NewDeleteGlobalSSHInstanceRequest()
	req.InstanceId = ucloud.String(instanceID)
	pathxClient.DeleteGlobalSSHInstance(req)
}

var areaMap = map[string]string{
	"洛杉矶":  "美国",
	"香港":   "中国",
	"新加坡":  "新加坡",
	"东京":   "日本",
	"法兰克福": "德国",
}
