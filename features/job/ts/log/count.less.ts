
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { CountData } from '../Query';

/**
 * 日志数量
 * @method GET
 */
export interface Request {

    /**
     * 工作ID
     */
    jobId: int64 

    /**
     * 日志类型 多个都会分割
     */
    type?: string

    /**
     * 关键字
     */
    q?: string

}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
