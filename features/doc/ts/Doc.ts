import { int64, int32 } from "./lib/less";

export enum DocType {
    File = 1, Dir = 2
}

/**
 * 文档
 * @type db
 */
export class Doc {

    /**
    * ID
    */
    id: int64 = 0

    /**
    * 父级ID
    * @index ASC
    */
    pid: int64 = 0

    /**
     * 标题
     * @length 2048
     */
    title: string = ""

    /**
    * 用户ID
    * @index ASC
    */
    uid: int64 = 0

    /**
    * 类型
    * @index ASC
    */
    type: DocType = 0

    /**
     * 路径
     * @length 2048
     */
    path: string = ""

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
     * 扩展名
     * @index DESC
     * @length 32
     */
    ext: string = ''

    /**
     * 最近修改时间
     * @index DESC
     */
    mtime: int64 = 0

    /**
     * 最近访问时间
     * @index DESC
     */
    atime: int64 = 0

    /**
     * 创建时间
     */
    ctime: int64 = 0

}