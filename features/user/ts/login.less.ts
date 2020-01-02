
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { User } from "./User";

/**
 * 登录
 * @method POST
 */
interface Request {

    /**
     * 用户名
     */
    name: string

    /**
     * 密码
     */
    password: string
}

interface Response extends BaseResponse {
    data?: User
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
