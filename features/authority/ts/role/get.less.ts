
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Role } from '../Role';

/**
 * 获取角色
 * @method GET
 */
interface Request {

    /**
     * 角色ID
     */
    id?: int64

    /**
     * 角色名称
     */
    name?: string
}

interface Response extends BaseResponse {
    data?: Role
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
