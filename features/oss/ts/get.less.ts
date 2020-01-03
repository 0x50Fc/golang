
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Type } from "./Type";

/**
 * 获取
 * @method GET
 */
interface Request {

    /**
     * 配置名称
     */
    name?: string

    /**
     * Key
     */
    key: string

    /**
     * 类型 默认 Type.Url
     */
    type?: Type

    /**
     * 超时时间(秒) 公开读不设置
     */
    expires?: number
}

interface Response extends BaseResponse {
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
