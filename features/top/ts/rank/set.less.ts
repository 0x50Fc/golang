
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Top } from '../Top';

/**
 * 修改排名
 * @method POST
 */
interface Request {

    /**
     * 推荐表名
     */
    name: string

    /**
     * 目标
     */
    tid: int64

    /**
     * 排名 0 表示不指定排名
     */
    rank: int32

}

interface Response extends BaseResponse {
    data?: Top
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
