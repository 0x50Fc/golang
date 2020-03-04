
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";

/**
 * 删除评论
 * @method POST
 */
interface Request {

    /**
     * 内容ID
     */
    id: int64

    /**
     * 用户ID
     */
    uid?: string


    /**
     * 群ID
     */
    groupid?: string


}


interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
