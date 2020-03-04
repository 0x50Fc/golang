import { int64 } from "./lib/less";

/**
 * 应用
 * @type db
 */
export class App {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 别名
     * @length 128
     */
    alias: string = ''

    /**
     * 类型
     * @length 128
     */
    type: string = ''

    /**
     * 内容
     * @length -1
     */
    content: string = ''

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