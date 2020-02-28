
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from '../Query';
import { Authority } from '../Authority';

/**
 * 查询授权
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface AuthorityQueryData {

    /**
     * 授权
     */
    items: Authority[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: AuthorityQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
