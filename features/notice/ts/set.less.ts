
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Notice } from "./Notice";

/**
 * 修改
 * @method POST
 */
interface Request {

    /**
     * 分组ID
     */
    id: int64

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 通知类型
     */
    type?: int32

    /**
     * 消息来源ID
     */
    fid?: int64

    /**
     * 消息来源项ID
     */
    iid?: int64

    /**
     * 通知内容
     */
    body?: string 

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
