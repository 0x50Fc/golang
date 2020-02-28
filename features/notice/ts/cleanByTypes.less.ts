
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";

/**
 * 清理消息
 * @method POST
 */
interface Request {


    /**
     * 类型, 多个逗号分割
     */
    type: string

    /**
     * 消息来源ID , 多个逗号分割
     */
    fid: int64

    /**
     * 消息来源项ID , 多个逗号分割
     */
    iid: int64


}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
