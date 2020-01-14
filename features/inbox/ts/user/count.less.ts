
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";

/**
 * 用户数量
 * @method GET
 */
export interface Request {

    /**
     * 类型
     */
    type: int64

    /**
     * 内容ID
     */
    mid?: int64

    /**
     * 内容项ID
     */
    iid?: int64

}


export interface UserCountData {
    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: UserCountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
