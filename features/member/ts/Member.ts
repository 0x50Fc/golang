import { int64, int32 } from "./lib/less";

/**
 * 成员
 * @type db
 */
export class Member {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 商户ID
     * @index ASC
     */
    bid: int64 = 0

    /**
     * 成员ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 备注名
     * @length 255
     */
    title: string = ""

    /**
     * 搜索关键字
     * @length 2048
     */
    keyword: string = ""

    /**
     * 其他数据
     * @length -1
     */
    options: any

    /**
     * 创建时间
     */
    ctime: int64 = 0

}
