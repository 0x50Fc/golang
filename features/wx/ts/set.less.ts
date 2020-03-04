
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { UserType } from './UserType';
import { User, UserState } from './User';
import { Token } from './Token';

/**
 * 修改用户关注状态
 * @method POST
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
     * 关注状态
     */
    state: UserState

}

interface Response extends BaseResponse {
    data?: User
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
