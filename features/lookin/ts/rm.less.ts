
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";

/**
 * 取消赞
 * @method POST
 */
interface Request {

    /**
     * 目标ID
     */
    tid: int64

    /**
     * 项ID 默认 0
     */
    iid?: int64
    
    /**
     * 用户ID
     */
    uid?: int64

     /**
     * 用户ID
     */
    fuid?: int64

    /**
     * 好友级别，多个逗号分割
     */
    flevel?: string
    
}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
