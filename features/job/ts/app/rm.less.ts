
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { App } from '../App';

/**
 * 删除应用
 * @method POST
 */
interface Request {

    /**
     * 应用ID
     */
    id: int64

}

interface Response extends BaseResponse {
    data?: App
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
