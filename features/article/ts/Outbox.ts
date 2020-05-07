import { int64 } from "./lib/less";

/**
 * 发布的动态
 * @type db
 */
export class Outbox {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 发布者ID
     * @index
     */
    uid: int64 = 0

    /**
     * 动态ID
     * @index
     */
    mid: int64 = 0

    /**
     * 内容
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
