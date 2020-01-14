
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Inbox } from "./Inbox";

/**
 * 清理
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 类型 type1 | type2 | type3
     */
    type?: int64
    
    /**
     * 内容ID
     */
    mid?: int64

    /**
     * 发布者ID
     */
    fuid?: int64
    
    /**
     * 内容项ID
     */
    iid?: int64

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
