
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';
import { MessageType } from '../MessageType';

/**
 * 发送消息
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * openid
     */
    openid: string

    /**
     * 消息类型
     */
    type: MessageType

    /**
     * 消息内容:
     * MessageType.Text:
     * {
         "content":"Hello World"
       }
     * MessageType.Image:
     * MessageType.Voice:
     * {
         "media_id":"MEDIA_ID"
       }
     * MessageType.Video:
     * {
         "media_id":"MEDIA_ID",
         "thumb_media_id":"MEDIA_ID",
         "title":"TITLE",
         "description":"DESCRIPTION"
       }
     * MessageType.Music:
     * {
         "title":"MUSIC_TITLE",
         "description":"MUSIC_DESCRIPTION",
         "musicurl":"MUSIC_URL",
         "hqmusicurl":"HQ_MUSIC_URL",
         "thumb_media_id":"THUMB_MEDIA_ID" 
       }
     * MessageType.News:
     * {
         "articles": [
         {
             "title":"Happy Day",
             "description":"Is Really A Happy Day",
             "url":"URL",
             "picurl":"PIC_URL"
         }
         ]
       }
     * MessageType.Mpnews:
     * {
         "media_id":"MEDIA_ID"
       }
     * MessageType.Msgmenu:
     * {
         "head_content": "您对本次服务是否满意呢? "
        "list": [
        {
            "id": "101",
            "content": "满意"
        },
        {
            "id": "102",
            "content": "不满意"
        }
        ],
        "tail_content": "欢迎再次光临"
       }
     * MessageType.Wxcard:
     * {           
        "card_id":"123dsdajkasd231jhksad"        
       }
     * MessageType.Miniprogrampage:
     * {
            "title":"title",
            "appid":"appid",
            "pagepath":"pagepath",
            "thumb_media_id":"thumb_media_id"
       }
     * MessageType.Template:
     * {
         "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
           "url":"http://weixin.qq.com/download",  
           "miniprogram":{
             "appid":"xiaochengxuappid12345",
             "pagepath":"index?foo=bar"
           },          
           "data":{
                   "first": {
                       "value":"恭喜你购买成功！",
                       "color":"#173177"
                   },
                   "keyword1":{
                       "value":"巧克力",
                       "color":"#173177"
                   },
                   "keyword2": {
                       "value":"39.8元",
                       "color":"#173177"
                   },
                   "keyword3": {
                       "value":"2014年9月22日",
                       "color":"#173177"
                   },
                   "remark":{
                       "value":"欢迎再次购买！",
                       "color":"#173177"
                   }
           }
       }
     */
    
    body: string

    /**
     * 客服账号
     */
    kf_account?: string

}


export interface Response extends BaseResponse {
  
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
