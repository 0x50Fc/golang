
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 获取数量
 * @method GET
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

}

export interface CountData {
    /**
     * 关注数量
     */
    followCount: int32
    /**
     * 粉丝数量
     */
    fansCount: int32
}

interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
