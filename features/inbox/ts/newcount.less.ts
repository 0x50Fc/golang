
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { GroupBy } from "./Query";

/**
 * 最新数量
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
     * 顶部ID
     */
    topId: int64

    /**
     * 分组
     */
    groupBy?: GroupBy

}

export interface NewInboxCountData {

    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: NewInboxCountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
