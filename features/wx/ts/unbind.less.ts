
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { UserType } from './UserType';
import { User } from './User';

/**
 * 解绑用户
 * @method GET
 */
interface Request {

    /**
     * 类型
     */
    type?: UserType

    /**
     * appid
     */
    appid?: string

    /**
     * openid
     */
    openid?: string

    /**
     * unionid
     */
    unionid?: string

    /**
     * 用户ID
     */
    uid: int64

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
