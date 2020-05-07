
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Member } from "./Member";

/**
 * 修改成员信息
 * @method POST
 */
interface Request {

    /**
     * 商户ID
     */
    bid: int64

    /**
     * 成员ID
     */
    uid: int64

    /**
     * 备注名
     */
    title?: string

    /**
     * 搜索关键字
     */
    keyword?: string

    /**
     * 其他数据 JSON
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Member
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
