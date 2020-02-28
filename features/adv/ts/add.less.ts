
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Adv } from "./Adv";

/**
 * 广告配置
 * @method GET
 */
export interface Request {

    /**
     * 标题
     */
    title: string

    /**
     * 频道
     * @index ASC
     */
    channel: string

    /**
     * 广告组位置
     */
    position: int32

    /**
     * 图片
     */
    pic: string

    /**
     * 描述
     */
    description: string

    /**
     * 跳转链接
     */
    link: string

    /**
     * 跳转类型
     */
    linktype: int32

    /**
     * 排序
     */
    sort: int32

    /**
     * 开始时间
     */
    starttime: int64

    /**
     * 结束时间
     */
    endtime: int64
}


export interface Response extends BaseResponse {
    data?: Adv
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}