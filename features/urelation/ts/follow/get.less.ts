
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Follow } from "../Follow";
import { QueryDataPage } from "../Query";

/**
 * 获取关注的好友
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 好友ID
     */
    fuid: int64

}

export interface Response extends BaseResponse {
    data?: Follow
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
