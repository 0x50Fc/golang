
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Slave, SlaveState } from '../Slave';

/**
 * 修改主机
 * @method POST
 */
interface Request {

    /**
     * 主机ID
     */
    id: int64

    /**
     * 别名前缀 默认不限制
     */
    prefix?: string

    /**
     * 生产授权token
     */
    token?: boolean

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

    /**
     * 状态
     */
    state?: SlaveState

    /**
     * 超时时间
     */
    etime?: int64

}

interface Response extends BaseResponse {
    data?: Slave
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
