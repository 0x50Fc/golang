
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Page } from "./Query";
import { Doc } from "./Doc";

/**
 * 查询
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

    
}


export interface DocQueryData {

    /**
     * 文档
     */
    items: Doc[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: DocQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
