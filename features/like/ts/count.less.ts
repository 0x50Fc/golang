
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 数量
 * @method GET
 */
export interface Request {

    /**
     * 目标
     */
    tid: int64

    /**
     * 项ID 默认 0
     */
    iid?: int64
    
    /**
     * 用户ID
     */
    uid?: int64
    
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
