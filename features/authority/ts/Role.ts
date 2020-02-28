import { int64, int32 } from "./lib/less";


/**
 * 角色
 * @type db
 */
export class Role {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 角色名
     * @unique ASC
     * @length 64
     */
    name: string = ''

    /**
     * 角色说明
     * @length 255
     */
    title: string = ''

    /**
     * 其他选项 JSON 叠加
     * @length -1
     */
    options: any

}