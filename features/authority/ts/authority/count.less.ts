
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";

/**
 * 授权数量
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 角色ID
     */
    roleId?: int64

    /**
     * 资源ID
     */
    resId?: int64

}


export interface AuthorityCountData {
    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: AuthorityCountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
