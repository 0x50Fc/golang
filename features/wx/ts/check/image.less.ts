
import { BaseResponse, ErrCode } from "../lib/BaseResponse"

/**
 * 检查图片
 * @method POST
 */
interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * 消息
     */
    url: string

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
