
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 删除
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
    * 查询条件
    */
    where?: string

    /**
     * 排序
     */
    orderBy?: string

    /**
     * 限制条件
     */
    limit?: string

    /**
     * 参数 JSON 
     * ["",123,245]
     */
    args?: string
}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
