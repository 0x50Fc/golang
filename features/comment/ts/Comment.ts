import { int64 } from "./lib/less";

export enum CommentState {
    /**
     * 可用
     */
    None = 0,
    /**
     * 已回收
     */
    Recycle = 1
}

/**
 * 评论
 * @type db
 */
export class Comment {

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
     * 评论目标ID
     * @index ASC
     */
    eid: int64 = 0

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 内容
     * @length -1
     */
    body: string = ""

    /**
     * 其他选项 JSON 叠加
     * @length -1
     */
    options: any

    /**
     * 创建时间
     */
    ctime: int64 = 0

    /**
     * 回收状态
     * @index ASC
     */
    state: CommentState = CommentState.None

    /**
     * 一个评论下的所有回复的路径
     * @index ASC
     */
    path: string = ""
    
}
