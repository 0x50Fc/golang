
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 数量
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * ID,多个逗号分割
     */
    ids?: string

    /**
     * 类型, 多个逗号分割
     */
    type?: string

    /**
     * 消息来源ID , 多个逗号分割
     */
    fid?: string

    /**
     * 消息来源项ID , 多个逗号分割
     */
    iid?: string

    /**
     * 模糊匹配关键字
     */
    q?: string

    /**
     * 顶部ID
     */
    topId?: int64

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
