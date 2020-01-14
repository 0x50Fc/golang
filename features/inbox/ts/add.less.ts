
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Inbox } from "./Inbox";

/**
 * 添加
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 类型
     */
    type: int64

    /**
     * 内容ID
     */
    mid: int64

    /**
     * 内容项ID
     */
    iid?: int64

    /**
     * 发布者ID
     */
    fuid: int64

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

    /**
     * 创建时间
     */
    ctime?: int32

}

interface Response extends BaseResponse {
    data?: Inbox
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
