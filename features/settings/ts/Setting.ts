import { int64, int32 } from "./lib/less";

/**
 * 系统设置
 * @type db
 */
export class Setting {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 设置名
     * @index ASC
     * @length 128
     */
    name: string = ''

    /**
     * 其他选项 JSON 叠加
     * @length -1
     */
    options: any

}