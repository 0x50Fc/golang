
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Job } from '../Job';

/**
 * 修改工作
 * @method POST
 */
interface Request {

    /**
     * 应用ID
     */
    id: int64

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
     * 总任务数
     */
    maxCount?: int32

    /**
     * 已执行任务数
     */
    count?: int32

    /**
     * 错误任务数
     */
    errCount?: int32

    /**
     * 增加已执行任务数
     */
    addCount?: int32

    /**
     * 增加错误任务数
     */
    addErrCount?: int32

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
