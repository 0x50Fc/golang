
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";

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
     * @index ASC
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


export interface CountData {
    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
