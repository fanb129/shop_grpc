package api

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"shop_api/user_web/forms"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"shop_api/user_web/global"
)

// GenerateSmsCode 生成width长度的短信验证码
func GenerateSmsCode(witdh int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	config := sdk.NewConfig()
	aliSmsInfo := global.ServerConfig.AliSmsInfo
	smsCode := GenerateSmsCode(6)
	credential := credentials.NewAccessKeyCredential(aliSmsInfo.ApiKey, aliSmsInfo.ApiSecrect)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := dysmsapi.NewClientWithOptions(aliSmsInfo.RegionId, config, credential)
	if err != nil {
		panic(err)
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.SignName = aliSmsInfo.SignName               //阿里云验证过的项目名 自己设置
	request.TemplateCode = aliSmsInfo.TemplateCode       //阿里云的短信模板号 自己设置
	request.PhoneNumbers = sendSmsForm.Mobile            //手机号
	request.TemplateParam = "{\"code\":" + smsCode + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	zap.S().Infof("aliSms response is %#v\n", response)

	//将验证码保存起来 - redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
