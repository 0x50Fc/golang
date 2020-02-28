import { int64, int32 } from "./lib/less";

/**
 * 权限
 * @type db
 */
export class Authority {

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
     * 角色ID
     * @index ASC
     */
    roleId: int64 = 0

    /**
     * 资源ID
     * @index ASC
     */
    resId: int64 = 0

    /**
     * 其他选项 JSON 叠加
     * @length -1
     */
    options: any

}