
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Slave } from '../Slave';
import { Page } from '../Query';

/**
 * 查询主机
 * @method GET
 */
export interface Request {

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface SlaveQueryData {
    /**
     * 主机
     */
    items: Slave[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: SlaveQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
