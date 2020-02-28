
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Notice } from "./Notice";
import { TopPage } from './Query';

/**
 * 查询
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface QueryData {
    /**
     * 通知
     */
    items: Notice[]

    /**
     * 分页
     */
    page?: TopPage
}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
