import { int64, int32 } from "./lib/less";

/**
 * Ver
 * @type db
 */
export class Ver {

    /**
    * ID
    */
    id: int64 = 0

    /**
    * 用户ID
    * @index ASC
    */
    appid: int64 = 0

    /**
     * 版本号
     * @index DESC
     */
    ver: int32 = 0

    /**
     * 应用信息 JSON
     * @length -1
     */
    info: any

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