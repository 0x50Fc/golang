import { int64 } from "../lib/less";

/**
 * 用户信息
 * @type db
 */
export class Info {

    /**
     * 用户信息ID
     */
    id: int64 = 0

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * key
     * @length 64
     * @index ASC
     */
    key: string = ""

    /**
     * 内容
     * @length -1
     */
    value: string = ""

}