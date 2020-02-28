
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Authority } from '../Authority';

/**
 * 验证用户权限
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid:int64

    /**
     * 资源路径
     */
    path: string

}

interface Response extends BaseResponse {
    
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
