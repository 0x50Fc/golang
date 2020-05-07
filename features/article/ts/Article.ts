import { int64 } from "./lib/less";

export enum ArticleState {
    
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
 * 文章
 * @type db
 */
export class Article {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 发布者ID
     * @index
     */
    uid: int64 = 0

    /**
     * 内容
     * @length -1
     */
    body: string = ""

    /**
     * 其他数据
     * @length -1
     */
    options: any

    /**
     * 创建时间
     */
    ctime: int64 = 0

    /**
     * 状态
     * @index ASC
     */
    state: ArticleState = ArticleState.None
    
}