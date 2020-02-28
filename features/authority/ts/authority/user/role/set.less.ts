
import { BaseResponse, ErrCode } from "../../../lib/BaseResponse"
import { int64, int32 } from "../../../lib/less";
import { Authority } from '../../../Authority';

/**
 * 用户设置角色
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 角色名, 多个逗号分割
     */
    role: string

}


interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
