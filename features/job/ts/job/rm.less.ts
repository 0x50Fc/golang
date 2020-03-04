
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Job } from '../Job';

/**
 * 删除工作
 * @method POST
 */
interface Request {

    /**
     * 工作ID
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
