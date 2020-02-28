
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Top } from './Top';

/**
 * 添加
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
     * 权重
     */
    rate: int32

    /**
     * 搜索关键字
     */
    keyword?: string

    /**
     * 时间默认当前时间
     */
    time?: int32

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
