
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { App } from '../App';

/**
 * 创建应用
 * @method POST
 */
interface Request {

    /**
     * 别名
     */
    alias: string

    /**
     * 类型
     */
    type: string

    /**
     * 内容
     */
    content: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: App
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
