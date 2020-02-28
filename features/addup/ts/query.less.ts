
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Addup } from './Addup';

/**
 * 查询
 * @method POST
 */
export interface Request {

    /**
     * 统计表名
     */
    name: string

    /**
     * 分区 默认无分区
     */
    region?: int32

    /**
     * 查询字段 SQL , 默认 * 
     */
    fields?: string

    /**
     * 查询条件
     */
    where?: string

    /**
     * 排序
     */
    orderBy?: string

    /**
     * 分组
     */
    groupBy?: string

    /**
     * 分组筛选条件
     */
    having?: string

    /**
     * 限制条件
     */
    limit?: string

    /**
     * 参数 JSON 
     * ["",123,245]
     */
    args?: string

    /**
     * 缓存
     */
    cacheKey?: string

}

export interface QueryData {
    /**
     * 记录
     */
    items: any[]
}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
