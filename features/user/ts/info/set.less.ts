
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { User } from "../User";
import { int64 } from "../lib/less";
import { InfoType } from "./InfoType";
import { Info } from "./Info";

/**
 * 修改用户信息
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * key
     */
    key: string

    /**
     * 类型
     */
    type?: InfoType

    /**
     * 内容
     */
    value?: string

}

interface Response extends BaseResponse {
    info?: Info
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
