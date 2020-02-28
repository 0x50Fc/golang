
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Res } from '../Res';

/**
 * 修改资源
 * @method POST
 */
interface Request {

    /**
     * 资源ID
     */
    id: int64

    /**
     * 资源路径
     */
    path?: string

    /**
     * 说明
     */
    title?: string

    /**
     * 其他选项 JSON 叠加
     */
    options?: string
}

interface Response extends BaseResponse {
    data?: Res
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
