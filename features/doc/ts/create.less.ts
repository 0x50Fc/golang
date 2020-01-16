
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { DocType, Doc } from "./Doc";

/**
 * 创建
 * @method POST
 */
interface Request {

    /**
    * 父级ID
    */
    pid?: int64

    /**
     * 标题
     */
    title: string

    /**
    * 用户ID
    */
    uid: int64

    /**
    * 类型
    */
    type?: DocType

    /**
     * 扩展名
     */
    ext?: string

    /**
     * 搜索关键字
     */
    keyword?: string

    /**
     * 其他数据  JSON
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Doc
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
