
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Role } from "../Role";

/**
 * 删除角色
 * @method POST
 */
interface Request {

    /**
     * 角色ID
     */
    id: int64

}

interface Response extends BaseResponse {
    data?: Role
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
