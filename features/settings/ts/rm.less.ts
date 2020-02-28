
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Setting } from './Setting';

/**
 * 删除
 * @method POST
 */
interface Request {

    /**
     * 名称
     */
    name: string

}

interface Response extends BaseResponse {
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
