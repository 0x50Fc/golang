
import { BaseResponse, ErrCode } from "../../lib/BaseResponse"
import { int64, int32 } from "../../lib/less";
import { User } from '../../User';
import { FormId } from '../../FormId';

/**
 * 小程序 使用 formid
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * openid
     */
    openid: string

}

export interface AppFormIdUseData {
    formid: string
}


export interface Response extends BaseResponse {
    data?: AppFormIdUseData
    formid?: FormId
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
