
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Top } from './Top';

/**
 * 删除
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

}

interface Response extends BaseResponse {
    data?: Top
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
