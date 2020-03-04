
import { BaseResponse, ErrCode } from "../../lib/BaseResponse"
import { int64, int32 } from "../../lib/less";
import { Job } from '../../Job';

/**
 * 主机更新工作进度
 * @method POST
 */
interface Request {

    /**
     * 主机 token
     */
    token: string

    /**
     * 主机超时时间 秒
     */
    expires: int64
    
    /**
     * 工作ID
     */
    jobId: int64

    /**
     * 工作是否已完成
     */
    done: boolean

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
}

interface Response extends BaseResponse {
    data?: Job
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
