
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Media } from "./Media";

/**
 * 添加
 * @method POST
 */
interface Request {

    /**
     * 存储表名
     */
    name?: string

    /**
     * 存储分区
     */
    region?: int32

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 类型
     */
    type?: string

    /**
     * 标题
     */
    title?: string

    /**
     * 关键字
     */
    keyword?: string

    /**
     * 存储路径
     */
    path?: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Media
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
