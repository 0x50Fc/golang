
import { BaseResponse, ErrCode } from "../../../lib/BaseResponse"
import { int64, int32 } from "../../../lib/less";
import { Open } from '../../../Open';

/**
 * 开发平台 公众号授权 更新 Ticket
 * @method GET
 */
export interface Request {

    /**
     * Ticket
     */
    ticket: string

}

export interface Response extends BaseResponse {
    data?: Open
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
