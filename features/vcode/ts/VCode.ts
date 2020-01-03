import { int64 } from "./lib/less";

/**
 * 验证码
 * @type db
 */
export class VCode {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * Key
     * @length 128
     * @unique ASC
     */
    key: string = ""

    /**
     * 数字验证码 最大 12位
     * @length 12
     */
    code: string = ""

    /**
     * 32位 HASH
     * @length 32
     */
    hash: string = ""

    /**
     * 过期时间
     */
    etime: int64 = 0

}