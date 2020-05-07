
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Member } from "./Member";

/**
 * 删除成员信息
 * @method POST
 */
interface Request {

    /**
     * 商户ID
     */
    bid: int64

    /**
     * 成员ID
     */
    uid: int64

}

interface Response extends BaseResponse {
    data?: Member
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
