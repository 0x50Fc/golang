
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { CountData } from '../Query';

/**
 * 应用数量
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

}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
