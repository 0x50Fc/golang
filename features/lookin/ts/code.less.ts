
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Lookin } from "./Lookin";

/**
 * 生成推荐码
 * @method POST
 */
interface Request {

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

}

interface CodeData {
    code?: string
}

interface Response extends BaseResponse {
    data?: CodeData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
