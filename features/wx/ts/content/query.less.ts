
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int32 , int64 } from "../lib/less";
import { ContentList } from "../ContentList";

/**
 * 查询
 * @method GET
 */
export interface Request {

    /**
     * ID
     */
    id?: int64

    /**
     * 微信原始数据库的msgId 消息id自增
     *
     */
    msgId : string

    /**
     * 微信原始数据库 的 createTime 消息时间(毫秒)
     */
    createTime: string

    /**
     * 消息发送者的微信id
     */
    msgTalkerUserAlias : string


    /**
     * 消息发送者的微信昵称
     */
    msgTalkerUserNickName : string


    /**
     * 微信原始数据库 的 type 判断消息类型
     */
    type : string


    /**
     * 微信群的微信用户名
     */
    msgGroupName : string


    /**
     * 开始时间
     */
    startTime?: int32

    /**
     * 结束时间
     */
    endTime?: int32


    /**
     * 模糊匹配关键字
     */
    q?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}


export interface Response extends BaseResponse {
    data?: ContentList
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
