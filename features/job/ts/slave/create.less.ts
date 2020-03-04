
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Slave } from '../Slave';

/**
 * 创建主机
 * @method POST
 */
interface Request {

    /**
     * 别名前缀 默认不限制
     */
    prefix?: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Slave
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
