import { int64 } from "./lib/less";

export enum FeedStatus {
    
    /**
     * 新动态
     */
    None = 0,
    
    /**
     * 已分发完成
     */
    Completed = 1,
}

export enum FeedState {
    
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
 * 动态
 * @type db
 */
export class Feed {

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
     * 状态
     * @index ASC
     */
    status: FeedStatus = FeedStatus.None

    /**
     * 创建时间
     */
    ctime: int64 = 0

    /**
     * 回收状态
     * @index ASC
     */
    state: FeedState = FeedState.None
    
}