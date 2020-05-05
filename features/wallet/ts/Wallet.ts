import { int64 } from "./lib/less";

/**
 * 钱包
 * @type db
 */
export class Wallet {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 别名
     * @length 64
     * @index ASC
     */
    alias: string = ""

    /**
     * 可用余额
     */
    value: int64 = 0

    /**
     * 冻结余额
     */
    freezeValue: int64 = 0

    /**
     * 累计收入
     */
    inValue: int64 = 0

    /**
     * 累计支出
     */
    outValue: int64 = 0

    /**
     * 其他数据
     */
    options: any

    /**
     * 创建时间
     */
    ctime: int64 = 0

}