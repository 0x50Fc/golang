import { int64, int32 } from "./lib/less";

/**
 * 资源
 * @type db
 */
export class Res {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 资源路径
     * @unique ASC
     * @length 128
     */
    path: string = ''

    /**
     * 说明
     * @length 255
     */
    title: string = ''

    /**
     * 其他选项 JSON 叠加
     * @length -1
     */
    options: any

}