
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';

/**
 * 开发平台授权确认
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * state
     */
    state?: string

    /**
     * code
     */
    code: string

}


export interface Response extends BaseResponse {
    data?: User
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
