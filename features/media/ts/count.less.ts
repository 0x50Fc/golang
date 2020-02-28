
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 数量
 * @method GET
 */
export interface Request {

    /**
     * 存储表名
     */
    name?: string

    /**
     * 存储分区
     */
    region?: int32

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 搜索关键字
     */
    q?: string

    /**
     * 路径前缀
     */
    prefix?: string

    /**
     * 类型,多个逗号分割
     */
    type?: string

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
