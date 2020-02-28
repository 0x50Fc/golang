
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { QueryData } from './Query';

/**
 * 广告评论
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
