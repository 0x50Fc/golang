import { int64, int32 } from "./lib/less";

/**
 * 评论的用户
 */
export class User {

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 发布的数量
     */
    count: int32 = 0

    /**
     * 最后发布时间
     */
    ctime: int64 = 0


}
