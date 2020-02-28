
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Notice } from "./Notice";

/**
 * 创建
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 通知类型 默认 0
     */
    type?: int32

    /**
     * 消息来源ID 默认 0
     */
    fid?: int64

    /**
     * 消息来源项ID 默认 0
     */
    iid?: int64

    /**
     * 通知内容
     */
    body: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Notice
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
