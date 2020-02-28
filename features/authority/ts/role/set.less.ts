
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Role } from '../Role';

/**
 * 修改角色
 * @method POST
 */
interface Request {

    /**
     * 资源ID
     */
    id: int64

    /**
     * 角色名
     */
    name?: string

    /**
     * 说明
     */
    title?: string

    /**
     * 其他选项 JSON 叠加
     */
    options?: string
}

interface Response extends BaseResponse {
    data?: Role
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
