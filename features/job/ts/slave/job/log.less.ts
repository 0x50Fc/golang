
import { BaseResponse, ErrCode } from "../../lib/BaseResponse"
import { int64, int32 } from "../../lib/less";
import { Job } from '../../Job';
import { LogType } from "../../Log";

/**
 * 主机更新工作日志
 * @method POST
 */
interface Request {

    /**
     * 主机 token
     */
    token: string

    /**
     * 工作ID
     */
    jobId: int64

    /**
     * 应用ID
     */
    appid: int64

    /**
     * 类型
     */
    type: LogType

    /**
     * 日志内容
     */
    body: string

}

interface Response extends BaseResponse {
  
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
