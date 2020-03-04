import { int64, int32 } from "./lib/less";

/**
 * 群内容
 * @type db
 */
export class Content {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 微信原始数据库 的 talker
     */
    talker : string = ''

    /**
     * 微信原始数据库 的 content 消息内容
     * @length 512
     */
    content : string = ''

    /**
     * 微信原始数据库 的 createTime 消息时间(毫秒)
     */
    createTime: string = ''

    /**
     * 处理后的消息时间
     */
    ctime: int32 = 0

    /**
     * 入库时间
     */
    etime: int64 = 0

    /**
     * 微信原始数据库 的 type 判断消息类型
     */
    type : string = ''

    /**
     * 微信原始数据库 的 isSend 判断是否是自己发送的消息
     */
    isSend : string = ''

    /**
     * 微信原始数据库的msgId 消息id自增
     * @index ASC
     */
    msgId : string = ''

    /**
     * 机器人id
     * @index ASC
     */
    robotUserAlias : string = ''

    /**
     * 机器人的微信用户名
     */
    robotUserName : string = ''

    /**
     * 机器人的微信昵称
     */
    robotUserNickName : string = ''

    /**
     * 消息发送者的微信用户名
     */
    msgTalkerUserName : string = ''

    /**
     * 消息发送者的微信id
     */
    msgTalkerUserAlias : string = ''

    /**
     * 消息发送者的微信昵称
     */
    msgTalkerUserNickName : string = ''

    /**
     * 消息发送者的微信头像大图
     */
    msgTalkerUserReserved1 : string = ''

    /**
     * 消息发送者的微信头像小图
     */
    msgTalkerUserReserved2 : string = ''

    /**
     * 微信群的微信用户名
     */
    msgGroupName : string = ''

    /**
     * 微信群的微信昵称
     */
    msgGroupNickName : string = ''

    /**
     * 处理过的聊天内容
     */
    msgContent : string = ''

}