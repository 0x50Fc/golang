import { int64, int32 } from "./lib/less";

/**
 * 广告
 * @type db
 */
export class Adv {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 频道
     * @index ASC
     */
    channel: string = ''

    /**
     * 标题
     * @index ASC
     */
    title: string = ''

    /**
     * 广告组位置
     * @index ASC
     */
    position: int32 = 0

    /**
     * 图片
     * @length 128
     */
    pic: string = ''

    /**
     * 描述
     * @length 512
     */
    description: string = ''

    /**
     * 跳转链接
     * @length 128
     */
    link: string = ''

    /**
     * 跳转类型
     */
    linktype: int32 = 0

    /**
     * 排序
     */
    sort: int32 = 0

    /**
     * 开始时间
     */
    starttime: int64 = 0

    /**
     * 结束时间
     */
    endtime: int64 = 0


    /**
     * 创建时间
     */
    ctime: int64 = 0

}