
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Top } from './Top';
import { QueryData } from './Query';

/**
 * 查询
 * @method POST
 */
export interface Request {

    /**
     * 推荐表名
     */
    name: string

    /**
     * 搜索关键字
     */
    q?: string

    /**
     * 目标ID,多个逗号分割
     */
    tids?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

    /**
     * 顶部ID
     */
    topId?: int64
}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
