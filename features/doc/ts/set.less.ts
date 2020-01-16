
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Doc } from "./Doc";

/**
 * 修改
 * @method POST
 */
interface Request {

    /**
    * ID
    */
    id: int64

    /**
    * 用户ID
    */
    uid: int64

    /**
    * 父级ID
    */
    pid?: int64

    /**
     * 标题
     */
    title?: string

    /**
     * 扩展名
     */
    ext?: string
    
    /**
     * 搜索关键字
     */
    keyword?: string

    /**
     * 其他数据 JSON
     */
    options?: string

    /**
     * 更新最近修改时间
     */
    mtime?: boolean

    /**
     * 更新最近访问时间
     */
    atime?: boolean

}

interface Response extends BaseResponse {
    data?: Doc
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
