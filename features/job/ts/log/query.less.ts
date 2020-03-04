
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from '../Query';
import { Job } from '../Job';
import { Log } from '../Log';

/**
 * 查询日志
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface LogQueryData {

    /**
     * 工作
     */
    items: Log[]

    /**
     * 分页
     */
    page?: Page

}

export interface Response extends BaseResponse {
    data?: LogQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
