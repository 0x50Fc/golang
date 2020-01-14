
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { TopPage, GroupBy } from './Query';
import { Inbox } from './Inbox';

/**
 * 查询
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 发布者ID
     */
    fuid?: int64

    /**
     * 类型 type1 | type2 | type3
     */
    type?: int64


    /**
     * 内容ID
     */
    mid?: int64

    /**
     * 内容项ID
     */
    iid?: int64

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

    /**
     * 顶部ID
     */
    topId?: int64

    /**
     * 分组
     */
    groupBy?: GroupBy
}


export interface InboxQueryData {

    /**
     * 收件
     */
    items: Inbox[]

    /**
     * 分页
     */
    page?: TopPage
}


export interface Response extends BaseResponse {
    data?: InboxQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
