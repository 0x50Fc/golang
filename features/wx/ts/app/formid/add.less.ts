
import { BaseResponse, ErrCode } from "../../lib/BaseResponse"
import { int64, int32 } from "../../lib/less";
import { User } from '../../User';

/**
 * 小程序 添加 formid
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * openid
     */
    openid: string

    /**
     * JSON
     * [
     *   {"formid":"123","etime":1}
     * ]
     */
    items: string
}


export interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
