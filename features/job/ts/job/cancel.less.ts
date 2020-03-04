
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Job } from '../Job';

/**
 * 取消工作
 * @method POST
 */
interface Request {

    /**
     * 应用ID
     */
    id: int64
    
}

interface Response extends BaseResponse {
    data?: Job
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
