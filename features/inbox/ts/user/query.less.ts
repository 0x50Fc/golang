
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64,int32 } from "../lib/less";
import { Page } from '../Query';

/**
 * 查询订阅用户
 * @method GET
 */
export interface Request {

    /**
     * 类型
     */
    type: int64

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

}

/**
 * 订阅的用户
 */
export class User {

    /**
     * 用户ID
     */
    uid: int64 = 0

    /**
     * 内容ID
     */
    mid: int64 = 0

    /**
     * 内容项ID
     */
    iid: int64 = 0


    /**
     * 最后时间
     */
    ctime: int64 = 0

    /**
     * 发布者ID
     */
    fuid: int64 = 0

}


export interface UserQueryData {

    /**
     * 用户
     */
    items: User[]

    /**
     * 分页
     */
    page?: Page
}



export interface Response extends BaseResponse {
    data?: UserQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
