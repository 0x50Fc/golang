
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { CountData } from '../Query';

/**
 * 排名最新数量
 * @method GET
 */
export interface Request {

    /**
     * 推荐表名
     */
    name: string
    
    /**
     * 搜索关键字
     */
    q?: string

    /**
     * 顶部ID
     */
    topId: int64

}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
