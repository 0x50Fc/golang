
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 发送邮件
 * @method POST
 */
interface Request {

    /**
     * 收件人
     */
    to: string

    /**
     * 标题
     */
    subject: string

    /**
     * 内容
     */
    body: string

    /**
     * 内容类型
     */
    contentType: string
}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
