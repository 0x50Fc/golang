
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { QueryDataPage } from "../Query";
import { Fans } from "../Fans";

/**
 * 查询粉丝
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 好友ID，多个逗号分割
     */
    in?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface FansQueryData {

    /**
     * 粉丝
     */
    items: Fans[]

    /**
     * 分页
     */
    page?: QueryDataPage
}


export interface Response extends BaseResponse {
    data?: FansQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
