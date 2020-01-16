
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 数量
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
    * 父级ID
    */
    pid?: int64

    /**
    * 类型
    */
    type?: int32

    /**
     * 扩展名
     */
    ext?: string

    /**
     * 路径前缀
     */
    prefix?: string

    /**
     * 搜索关键字
     */
    q?: string

}


export interface DocCountData {

    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: DocCountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
