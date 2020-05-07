
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Lookin } from "./Lookin";

/**
 * 添加在看好友
 * @method POST
 */
interface Request {

    /**
     * 目标ID
     */
    tid: int64

    /**
     * 项ID
     */
    iid?: int64

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 代码
     */
    fcode?: string

    /**
     * 好友ID
     */
    fuid?: int64

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface AddData {
    items: Lookin[]
}

interface Response extends BaseResponse {
    data?: AddData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
