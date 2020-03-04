
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { CountData } from '../Query';

/**
 * 工作数量
 * @method GET
 */
export interface Request {

    /**
     * 类型
     */
    type?: string

    /**
     * 别名前缀
     */
    prefix?: string

    /**
     * 别名
     */
    alias?: string

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 应用ID
     */
    appid?: int64

    /**
     * 主机ID
     */
    sid?: int64
}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
