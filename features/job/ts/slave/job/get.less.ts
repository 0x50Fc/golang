
import { BaseResponse, ErrCode } from "../../lib/BaseResponse"
import { int64 } from "../../lib/less";
import { Job } from '../../Job';

/**
 * 主机获取可用工作
 * @method GET
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
    
}

interface Response extends BaseResponse {
    data?: Job
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
