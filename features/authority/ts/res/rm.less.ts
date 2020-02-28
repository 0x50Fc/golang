
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Res } from '../Res';

/**
 * 删除资源
 * @method POST
 */
interface Request {

    /**
     * 资源ID
     */
    id: int64

}

interface Response extends BaseResponse {
    data?: Res
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
