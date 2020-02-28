
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";

/**
 * 资源数量
 * @method GET
 */
export interface Request {

    /**
     * 前缀
     */
    prefix?: string

}


export interface ResCountData {
    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: ResCountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
