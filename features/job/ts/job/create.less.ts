
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Job } from '../Job';

/**
 * 创建工作
 * @method POST
 */
interface Request {

    /**
     * 类型
     */
    type?: int64

    /**
     * 应用ID
     */
    appid?: int64

    /**
     * 别名
     */
    alias?: string

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

    /**
     * 开始时间
     */
    stime?: int64
}

interface Response extends BaseResponse {
    data?: Job
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
