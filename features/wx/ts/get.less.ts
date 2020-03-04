
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { UserType } from './UserType';
import { User } from './User';
import { Token } from './Token';

/**
 * 获取用户
 * @method GET
 */
interface Request {

    /**
     * 类型,多个逗号分割
     */
    type: UserType

    /**
     * appid
     */
    appid: string

    /**
     * openid
     */
    openid: string

    /**
     * 是否更新用户信息
     */
    update?: boolean

}

interface Response extends BaseResponse {
    token?: Token
    data?: User
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
