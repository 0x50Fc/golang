
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Ver } from "../Ver";

/**
 * 获取版本
 * @method GET
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

}

interface Response extends BaseResponse {
    data?: Ver
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
