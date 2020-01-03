
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
     * 超时时间(秒) type == url 使用
     */
    expires?: number
}

interface PostData {
    /**
     * 上传URL
     */
    url: string
    /**
     * 上传 Form Data
     */
    data: any
}

interface Response extends BaseResponse {
    data?: PostData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
