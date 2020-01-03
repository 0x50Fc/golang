
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Type } from "./Type";

/**
 * 删除
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

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
