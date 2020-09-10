package VBOSS

import (
	"fmt"
	"vb-server/utils/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket*oss.Bucket;

func InitData(){

	//get config
	ossInfo:=config.VBGetConfigDBOSS()

	//oss arg
	fmt.Println("对象存储[OSS]参数:")
	fmt.Println("访问域名:",ossInfo.EndPoint)
	if ossInfo.EndPointHost==1{
		//启用自定义域名
		fmt.Println("[自定义域名已经启用]")
	}
	fmt.Println("默认Bucket:",ossInfo.Bucket)
	fmt.Println("默认目录:",ossInfo.Dirctory)

	//init
	// 创建OSSClient实例。
	client, err := oss.New(ossInfo.EndPoint, ossInfo.AccessKeyId, ossInfo.AccessKeySecret)
	if err != nil {
		panic(err)

	}

	// 获取存储空间。
	var bucketError error;
	Bucket, bucketError = client.Bucket(ossInfo.Bucket)
	if bucketError != nil {
		panic(err)

	}

}


