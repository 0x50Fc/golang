import { int64 } from "./lib/less";

export enum RelationType {
    Weak = 0,
    Strong = 1
}

/**
 * 关系表基类
 */
export class Relation {

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
     * 好友ID
     * @index ASC
     */
    fuid: int64 = 0

    /**
     * 类型
     */
    type: RelationType = RelationType.Weak

    /**
     * 创建时间
     */
    ctime: int64 = 0

}