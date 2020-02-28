
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Relation } from "./Relation";
import { Follow } from "./Follow";

/**
 * 取消关注
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 好友ID
     */
    fuid: int64

}

interface Response extends BaseResponse {
    data?: Follow
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
