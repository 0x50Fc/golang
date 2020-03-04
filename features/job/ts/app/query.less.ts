
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from '../Query';
import { App } from '../App';

/**
 * 查询应用
 * @method GET
 */
export interface Request {

    /**
     * 类型
     */
    type?: string

    /**
     * 别名前缀
     */
    prefix?: string

    /**
     * 别名
     */
    alias?: string
    
    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface AppQueryData {
    /**
     * 应用
     */
    items: App[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: AppQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
