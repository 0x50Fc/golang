
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { User } from "./User";
import { int64 } from "./lib/less";

/**
 * 获取用户
 * @method GET
 */
interface Request {

    /**
     * 用户ID
     */
    id?: int64

    /**
     * 用户名
     */
    name?: string

    /**
     * 昵称
     */
    nick?: string

    /**
     * 是否自动创建, name 必须存在
     */
    autocreate?: boolean
}

interface Response extends BaseResponse {
    data?: User
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
