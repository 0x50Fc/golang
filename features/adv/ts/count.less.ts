
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { CountData } from './Query';

/**
 * 获取广告数量
 * @method GET
 */
export interface Request {

    /**
     * 广告ID
     */
    id?: int64

    /**
     * 频道
     */
    channel?: string

    /**
     * 开始时间
     */
    stime?: int64

    /**
     * 结束时间
     */
    etime?: int64


}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
