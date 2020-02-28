
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Like } from "./Like";

/**
 * 点赞
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
    uid: int64

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Like
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
