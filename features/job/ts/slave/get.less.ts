
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Slave, SlaveState } from '../Slave';

/**
 * 获取主机
 * @method GET
 */
interface Request {

    /**
     * 主机ID
     */
    id: int64

}

interface Response extends BaseResponse {
    data?: Slave
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
