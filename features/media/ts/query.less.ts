
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Media } from "./Media";

/**
 * 查询
 * @method GET
 */
export interface Request {

    /**
     * 存储表名
     */
    name?: string

    /**
     * 存储分区
     */
    region?: int32

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 搜索关键字
     */
    q?: string

    /**
     * 路径前缀
     */
    prefix?: string

    /**
     * 类型,多个逗号分割
     */
    type?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32
    
}

export interface QueryDataPage {
    /**
     * 分页位置
     */
    p: int32
    /**
    * 单页记录数
    */
    n: int32
    /**
     * 总页数
     */
    count: int32
    /**
     * 总记录数
     */
    total: int32
}

export interface QueryData {
    /**
     * 媒体
     */
    items: Media[]

    /**
     * 分页
     */
    page?: QueryDataPage
}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
