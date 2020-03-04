
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from '../Query';
import { Job } from '../Job';

/**
 * 查询工作
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface JobQueryData {
    /**
     * 工作
     */
    items: Job[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: JobQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
