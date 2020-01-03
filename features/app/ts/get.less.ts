
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { App } from "./App";

/**
 * 获取
 * @method GET
 */
interface Request {

    /**
    * ID
    */
    id: int64

    /**
    * 用户ID
    */
    uid?: int64
}

interface Response extends BaseResponse {
    data?: App
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
