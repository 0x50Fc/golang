import { int64 } from "./lib/less";

/**
 * 在看
 * @type db
 */
export class Lookin {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 目标ID
     * @index ASC
     */
    tid: int64 = 0

    /**
     * 项ID
     * @index ASC
     */
    iid: int64 = 0

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 好友ID
     * @index ASC
     */
    fuid: int64 = 0

    /**
     * 好友推荐码
     */
    fcode: string = ""

    /**
     * 关系级别
     * @index ASC
     */
    flevel: int64 = 0

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
