
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { App } from "./App";

/**
 * 修改
 * @method POST
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

    /**
      * 默认版本号
      */
    ver?: int32
    /**
     * 标题
     */
    title?: string

    /**
     * 其他数据
     */
    options?: string
}

interface Response extends BaseResponse {
    data?: App
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
