import { int64 } from "./lib/less";

/**
 * 收件箱
 * @type db
 */
export class Inbox {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 收件类型
     * @index desc
     */
    type: int64 = 0

    /**
     * 接受者ID
     * @index desc
     */
    uid: int64 = 0

    /**
     * 发布者ID
     * @index desc
     */
    fuid: int64 = 0

    /**
     * 内容ID
     * @index desc
     */
    mid: int64 = 0

    /**
     * 内容项ID
     * @index desc
     */
    iid: int64 = 0

    /**
    * 其他数据
    * @length -1
    */
    options: any

    /**
     * 创建时间
     * @index desc
     */
    ctime: int64 = 0

}
