
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Notice } from "./Notice";

/**
 * 清理消息
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 类型, 多个逗号分割
     */
    type?: string

    /**
     * 消息来源ID , 多个逗号分割
     */
    fid?: string

    /**
     * 消息来源项ID , 多个逗号分割
     */
    iid?: string
    
}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
