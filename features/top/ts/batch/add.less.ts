
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Top } from '../Top';

/**
 * 批量添加
 * @method POST
 */
interface Request {

    /**
     * 推荐表名
     */
    name: string

    /**
     * 其他数据 JSON
     * [ {tid : 1, rate : 1, options:  {}, keyword:''} ]
     */
    items?: string

}

interface Response extends BaseResponse {
    data?: Top[]
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
