
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Follow } from "../Follow";
import { QueryDataPage } from "../Query";
import { Fans } from '../../../../market-less/despatch/feed/kk/urelation/ObjectSet';

/**
 * 获取粉丝
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
    data?: Fans
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
