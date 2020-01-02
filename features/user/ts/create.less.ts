
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { User } from "./User";
import { int64 } from "./lib/less";

/**
 * 创建用户
 * @method POST
 */
interface Request {

    /**
     * 用户名
     */
    name: string

    /**
     * 昵称
     */
    nick?: string

    /**
     * 密码
     */
    password?: string

}

interface Response extends BaseResponse {
    data?: User
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
