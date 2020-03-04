package wx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/http"
	"github.com/hailongz/golang/micro"
)

func MP_Send(app micro.IContext, mp interface{}, method string, u string, data interface{}) (interface{}, error) {
	return MP_SendWithType(app, mp, method, u, data, http.OptionTypeJson)
}

func MP_SendWithType(app micro.IContext, mp interface{}, method string, u string, data interface{}, dataType string) (interface{}, error) {

	var result interface{} = nil
	var err error = micro.NewError(ERROR_CONFIG, "未找到 baseURL 配置")

	dynamic.Each(dynamic.Get(mp, "baseURL"), func(_ interface{}, baseURL interface{}) bool {

		options := http.Options{}
		options.Method = method
		options.Url = fmt.Sprintf("%s%s", baseURL, u)
		options.Data = data
		options.ResponseType = http.OptionResponseTypeJson
		options.Type = dataType

		app.Println("[MP_Send]", options.Url, options.Data)

		result, err = http.Send(&options)

		app.Println("[MP_Send]", result, err)

		if err != nil || result == nil {
			return true
		}

		errcode := dynamic.IntValue(dynamic.Get(result, "errcode"), 0)

		if errcode == -1 {
			return true
		}

		if errcode != 0 {
			err = micro.NewError(int(errcode), dynamic.StringValue(dynamic.Get(result, "errmsg"), "未知微信服务错误"))
			return false
		}

		return false
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func MP_GetSecret(mp interface{}, appid string) (string, error) {
	v := dynamic.GetWithKeys(mp, []string{"keys", appid})
	if v == nil {
		return "", micro.NewError(ERROR_CONFIG, "未找到 secret 配置 "+appid)
	}
	return dynamic.StringValue(v, ""), nil
}

func MP_GetAccessToken(app micro.IContext, userType int32, appid string, forceUpdate bool) (string, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	key := fmt.Sprintf("mp_token_%d_%s", userType, appid)

	secret, err := MP_GetSecret(mp, appid)

	if err != nil {
		return "", err
	}

	if !forceUpdate {

		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			text, err := redis.Get(prefix + key).Result()

			if err == nil {
				return text, nil
			}
		}
	}

	v := Token{}

	if !forceUpdate {
		conn, prefix, err := app.GetDB("rd")

		if err == nil {

			p, err := db.Get(conn, &v, prefix, " WHERE type=? AND appid=?", userType, appid)

			now := time.Now().Unix()

			if err == nil && p != nil && v.Etime > now {

				{
					redis, fix, err := app.GetRedis("default")

					if err == nil {
						redis.Set(fix+key, v.AccessToken, time.Second*time.Duration(v.Etime-now)).Result()
					}
				}

				return v.AccessToken, nil

			}
		}
	}

	ret, err := MP_Send(app, mp, "GET", "/cgi-bin/token", map[string]interface{}{"grant_type": "client_credential", "appid": appid, "secret": secret})

	if err != nil {
		return "", err
	}

	token := dynamic.StringValue(dynamic.Get(ret, "access_token"), "")
	expires_in := dynamic.IntValue(dynamic.Get(ret, "expires_in"), 0)
	now := time.Now().Unix()

	{
		conn, prefix, err := app.GetDB("wd")

		if err == nil {

			if v.Id == 0 {
				_, err = db.Get(conn, &v, prefix, " WHERE type=? AND appid=?", userType, appid)
				if err != nil {
					return "", err
				}
			}

			v.Type = UserType_MP
			v.Appid = appid
			v.AccessToken = token
			v.Etime = now + expires_in - 30

			if v.Id != 0 {

				_, err = db.Update(conn, &v, prefix)

				if err != nil {
					return "", err
				}

			} else {

				_, err = db.Insert(conn, &v, prefix)

				if err != nil {
					return "", err
				}
			}

			{
				redis, fix, err := app.GetRedis("default")

				if err == nil {
					redis.Set(fix+key, v.AccessToken, time.Second*time.Duration(v.Etime-now)).Result()
				}
			}

		}
	}

	return token, nil
}

func MP_GetTicket(app micro.IContext, appid string, ticketType string) (string, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	key := "mp_ticket_" + ticketType + "_" + appid

	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			text, err := redis.Get(prefix + key).Result()

			if err == nil {
				return text, nil
			}
		}
	}

	v := Ticket{}

	{
		conn, prefix, err := app.GetDB("rd")

		if err == nil {

			p, err := db.Get(conn, &v, prefix, " WHERE type=? AND appid=?", ticketType, appid)

			now := time.Now().Unix()

			if err == nil && p != nil && v.Etime > now {

				{
					redis, fix, err := app.GetRedis("default")

					if err == nil {
						redis.Set(fix+key, v.Ticket, time.Second*time.Duration(v.Etime-now)).Result()
					}
				}

				return v.Ticket, nil

			}
		}
	}

	token, err := MP_GetAccessToken(app, UserType_MP, appid, false)

	if err != nil {
		return "", err
	}

	ret, err := MP_Send(app, mp, "GET", "/cgi-bin/ticket/getticket", map[string]interface{}{"access_token": token, "type": ticketType})

	if err != nil {
		return "", err
	}

	ticket := dynamic.StringValue(dynamic.Get(ret, "ticket"), "")
	expires_in := dynamic.IntValue(dynamic.Get(ret, "expires_in"), 0)
	now := time.Now().Unix()

	{
		conn, prefix, err := app.GetDB("wd")

		if err == nil {

			if v.Id == 0 {
				_, err = db.Get(conn, &v, prefix, " WHERE type=? AND appid=?", ticketType, appid)
				if err != nil {
					return "", err
				}
			}

			v.Type = ticketType
			v.Appid = appid
			v.Ticket = ticket
			v.Etime = now + expires_in - 30

			if v.Id != 0 {

				_, err = db.Update(conn, &v, prefix)

				if err != nil {
					return "", err
				}

			} else {

				_, err = db.Insert(conn, &v, prefix)

				if err != nil {
					return "", err
				}
			}

			{
				redis, fix, err := app.GetRedis("default")

				if err == nil {
					redis.Set(fix+key, v.Ticket, time.Second*time.Duration(v.Etime-now)).Result()
				}
			}

		}
	}

	return token, nil
}

func MP_UpdateUser(app micro.IContext, user *User) error {

	if user.Type == UserType_MP {

		mp := dynamic.Get(app.GetConfig(), "mp")

		if user.AccessToken != "" {

			data, err := MP_Send(app, mp, "GET", "/sns/userinfo", map[string]interface{}{"access_token": user.AccessToken, "openid": user.Openid, "lang": "zh_CN"})

			if err == nil {
				user.Gender = int32(dynamic.IntValue(dynamic.Get(data, "sex"), int64(user.Gender)))
				user.Nick = dynamic.StringValue(dynamic.Get(data, "nickname"), user.Nick)
				user.Logo = dynamic.StringValue(dynamic.Get(data, "headimgurl"), user.Logo)
				user.Lang = dynamic.StringValue(dynamic.Get(data, "language"), user.Lang)
				user.Province = dynamic.StringValue(dynamic.Get(data, "province"), user.Province)
				user.City = dynamic.StringValue(dynamic.Get(data, "city"), user.City)
				user.Country = dynamic.StringValue(dynamic.Get(data, "country"), user.Country)
				user.Unionid = dynamic.StringValue(dynamic.Get(data, "unionid"), user.Unionid)
				return nil
			}

		}

		token, err := MP_GetAccessToken(app, UserType_MP, user.Appid, false)

		if err != nil {
			return err
		}

		data, err := MP_Send(app, mp, "GET", "/cgi-bin/user/info", map[string]interface{}{"access_token": token, "openid": user.Openid, "lang": "zh_CN"})

		if err != nil {
			return err
		}

		user.Gender = int32(dynamic.IntValue(dynamic.Get(data, "sex"), int64(user.Gender)))
		user.Nick = dynamic.StringValue(dynamic.Get(data, "nickname"), user.Nick)
		user.Logo = dynamic.StringValue(dynamic.Get(data, "headimgurl"), user.Logo)
		user.Lang = dynamic.StringValue(dynamic.Get(data, "language"), user.Lang)
		user.Province = dynamic.StringValue(dynamic.Get(data, "province"), user.Province)
		user.City = dynamic.StringValue(dynamic.Get(data, "city"), user.City)
		user.Country = dynamic.StringValue(dynamic.Get(data, "country"), user.Country)
		user.Unionid = dynamic.StringValue(dynamic.Get(data, "unionid"), user.Unionid)

		return nil

	} else {
		return micro.NewError(ERROR_NOT_FOUND, "无法更新用户信息")
	}

}

func MP_NewNonceStr() string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(m.Sum(nil))
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func MP_AppDecrypt(app micro.IContext, sessionKey string, encryptedData string, iv string) ([]byte, error) {

	data, err := base64.StdEncoding.DecodeString(encryptedData)

	if err != nil {
		app.Println("[EncryptedData]", err)
		return nil, err
	}

	key, err := base64.StdEncoding.DecodeString(sessionKey)

	if err != nil {
		app.Println("[session_key]", err, sessionKey)
		return nil, err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		app.Println("[aes.NewCipher]", err)
		return nil, err
	}

	iv_b, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		app.Println("[iv]", err, iv)
		return nil, err
	}

	if len(iv_b) < block.BlockSize() {
		app.Println("[iv]", err, iv)
		return nil, micro.NewError(ERROR_NOT_FOUND, "错误的 iv")
	}

	cdc := cipher.NewCBCDecrypter(block, iv_b[0:block.BlockSize()])

	ret := make([]byte, len(data))

	cdc.CryptBlocks(ret, data)

	ret = PKCS5UnPadding(ret)

	return ret, nil
}

func MP_Open_GetTicket(app micro.IContext) (string, error) {

	appid := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"open", "appid"}), "")

	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {
			v, err := redis.Get(prefix + "open_ticket_" + appid).Result()
			if err != nil && v != "" {
				return v, nil
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return "", err
	}

	v := Open{}

	p, err := db.Get(conn, &v, prefix, " WHERE appid=? ", appid)

	if err != nil {
		return "", err
	}

	if p == nil {
		return "", micro.NewError(ERROR_NOT_FOUND, "未找到 Ticket")
	}

	return v.Ticket, nil
}

func MP_Open_GetAccessToken(app micro.IContext, appid string) (string, error) {

	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {
			v, err := redis.Get(prefix + "open_access_token" + appid).Result()
			if err != nil && v != "" {
				return v, nil
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return "", err
	}

	v := Open{}

	p, err := db.Get(conn, &v, prefix, " WHERE appid=? ", appid)

	if err != nil {
		return "", err
	}

	if p == nil {
		return "", micro.NewError(ERROR_NOT_FOUND, "未找到 Access Token")
	}

	if v.AccessToken == "" || time.Now().Unix() > v.Etime {

		mp := dynamic.Get(app.GetConfig(), "mp")

		secret, err := MP_GetSecret(mp, appid)

		if err != nil {
			return "", err
		}

		ret, err := MP_Send(app, mp, "POST", "/cgi-bin/component/api_component_token", map[string]interface{}{
			"component_appid":         appid,
			"component_appsecret":     secret,
			"component_verify_ticket": v.Ticket,
		})

		if err != nil {
			return "", err
		}

		v.AccessToken = dynamic.StringValue(dynamic.Get(ret, "component_access_token"), v.AccessToken)

		expires_in := dynamic.IntValue(dynamic.Get(ret, "expires_in"), 0)
		v.Etime = time.Now().Unix() + expires_in

		{
			conn, prefix, err := app.GetDB("wd")

			if err != nil {
				return "", err
			}

			_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"access_token": true, "etime": true})

			if err != nil {
				return "", err
			}
		}

		{
			redis, prefix, err := app.GetRedis("default")

			if err == nil {
				redis.Set(prefix+"open_access_token"+appid, v.AccessToken, time.Second*time.Duration(expires_in)).Result()
			}
		}
	}

	return v.AccessToken, nil
}
