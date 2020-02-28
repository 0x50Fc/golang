
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int32, int64 } from "./lib/less";
import { Adv } from "./adv";

/**
 * 删除广告
 * @method POST
 */
interface Request {

    /**
     * 评论ID
     */
    id: int64

    /**
     * 频道
     */
    channel: string

    /**
     * 广告组位置
     */
    position: int32

}

interface Response extends BaseResponse {
    data?: Adv
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
