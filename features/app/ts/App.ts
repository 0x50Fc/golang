import { int64, int32 } from "./lib/less";

/**
 * App
 * @type db
 */
export class App {

    /**
    * ID
    */
    id: int64 = 0

    /**
     * 标题
     * @length 255
     */
    title: string = ""

    /**
    * 用户ID
    */
    uid: int64 = 0

    /**
     * 默认版本号
     */
    ver: int32 = 0

    /**
     * 最新版本号
     */
    lastVer: int32 = 0

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