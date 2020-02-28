
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Media } from "./Media";

/**
 * 删除
 * @method POST
 */
interface Request {

    /**
     * 存储表名
     */
    name?: string

    /**
     * 存储分区
     */
    region?: int32

    /**
     * 媒体ID
     */
    id: int64

    /**
     * 用户ID
     */
    uid?: int64
}

interface Response extends BaseResponse {
    data?: Media
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
