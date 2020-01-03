
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Ver } from "../Ver";


/**
 * 删除版本
 * @method POST
 */
interface Request {

    /**
    * appid
    */
    appid: int64

    /**
     * 版本号
     */
    ver:int32

    /**
     * INFO
     */
    info?: string

    /**
     * 其他数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Ver
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
