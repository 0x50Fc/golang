
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Addup } from './Addup';

/**
 * 上传统计数据
 * @method POST
 */
interface Request {

    /**
     * 统计表名
     */
    name: string

    /**
     * 分区 默认无分区
     */
    region?: int32

    /**
     * 统计项ID
     */
    iid: int64

    /**
     * iid + time + unionKeys唯一
     * JSON 对象
     */
    unionKeys?: string

    /**
     * 设置数据项的值 JSON
     * { "key" : "value" }
     */
    set?: string

    /**
     * 增加数据项的值 JSON
     * { "key" : "value" }
     */
    add?: string

    /**
     * 时间 (秒)
     */
    time: int64
}

interface Response extends BaseResponse {
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
