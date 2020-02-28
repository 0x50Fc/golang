import { int64, int32 } from "./lib/less";

/**
 * 通知
 * @type db
 */
export class Notice {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 通知类型
     * @index ASC
     */
    type: int32 = 0

    /**
     * 消息来源ID
     * @index ASC
     */
    fid: int64 = 0

    /**
     * 消息来源项ID
     * @index ASC
     */
    iid: int64 = 0

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 通知内容
     * @length -1
     */
    body: string = ""

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