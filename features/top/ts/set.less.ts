
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Top } from './Top';

/**
 * 修改
 * @method POST
 */
interface Request {

    /**
     * 推荐表名
     */
    name: string

    /**
     * 目标
     */
    tid: int64

    /**
     * 搜索关键字
     */
    keyword?: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Top
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
