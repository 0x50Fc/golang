
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Type } from "./Type";

/**
 * 上传
 * @method POST
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
     * 类型 默认 url
     */
    type?: Type

    /**
     * 内容
     * 当 Type.Text || Type.Base64 时使用
     * 当 Type.URL 时 JSON.stringify({
     *  "url":"",
     *  "header":{}
     * })
     */
    content?: string

    /**
     * 超时时间(秒) type == url 使用
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
