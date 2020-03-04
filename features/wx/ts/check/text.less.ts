
import { BaseResponse, ErrCode } from "../lib/BaseResponse"

/**
 * 检查文本
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
    content: string

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
