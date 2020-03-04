
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';
import { MessageType } from '../MessageType';

/**
 * 获取小程序二维码
 * @method GET
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
     */
    scene: string

    /**
     * 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
     */
    page?: string

    /**
     * 二维码的宽度，单位 px，最小 280px，最大 1280px
     */
    width?: number

}

export interface WXAppQRData {

    /**
     * 类型
     */
    type: string

    /**
     * 内容
     */
    content: string
}

export interface Response extends BaseResponse {
    data?: WXAppQRData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
