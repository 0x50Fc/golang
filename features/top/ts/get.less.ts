
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Top } from './Top';

/**
 * 获取推荐项
 * @method GET
 */
interface Request {

    /**
     * 推荐表名
     */
    name: string

    /**
     * ID
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
