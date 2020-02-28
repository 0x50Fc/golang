
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Top } from '../Top';

/**
 * 计算排名
 * @method POST
 */
interface Request {

    /**
     * 推荐表名
     */
    name: string

    /**
     * 限制数量
     */
    limit: int32

}

interface Response extends BaseResponse {
 
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
