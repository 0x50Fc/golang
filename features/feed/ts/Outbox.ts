import { int64 } from "./lib/less";
import { FeedState } from '../../../market-less/despatch/top/app/kk/feed/ObjectSet';

export enum OutboxStatus {
    
    /**
     * 草稿
     */
    None = 0,
   
    /**
     * 已发送
     */
    Sended = 1,
}

/**
 * 发布的动态
 * @type db
 */
export class Outbox {

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
     * 动态ID
     * @index
     */
    mid: int64 = 0

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
     * @index DESC
     */
    status: OutboxStatus = OutboxStatus.None

    /**
     * 创建时间
     */
    ctime: int64 = 0
    
}
