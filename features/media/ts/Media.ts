import { int64 } from "./lib/less";

/**
 * 媒体
 * @type db
 */
export class Media {

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
     * 类型
     * @index ASC
     * @length 32
     */
    type: string = ''

    /**
     * 标题
     * @length 2048
     */
    title: string = ''

    /**
     * 关键字
     * @length 4096
     */
    keyword: string = ''

    /**
     * 存储路径
     * @length 2048
     */
    path: string = ''

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
