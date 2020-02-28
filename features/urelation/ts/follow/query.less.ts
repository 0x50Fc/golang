
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Follow } from "../Follow";
import { QueryDataPage } from "../Query";

/**
 * 查询关注的好友
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
     * 模糊匹配关键字
     */
    q?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface FollowQueryData {

    /**
     * 关系
     */
    items: Follow[]

    /**
     * 分页
     */
    page?: QueryDataPage
}


export interface Response extends BaseResponse {
    data?: FollowQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
