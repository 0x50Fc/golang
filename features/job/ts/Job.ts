import { int64, int32 } from "./lib/less";

/**
 * 工作状态
 */
export enum JobState {
    /**
     * 等待执行
     */
    None,
    /**
     * 执行中
     */
    Running,
    /**
     * 用户取消
     */
    Cancel,
    /**
     * 执行完成
     */
    Finish,
    /**
     * 错误中断
     */
    Error
}

/**
 * 工作
 * @type db
 */
export class Job {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 别名
     * @length 128
     * @index ASC
     */
    alias: string = ''

    /**
     * 类型
     * @index ASC
     */
    type: int32 = 0

    /**
     * 状态
     * @index ASC
     */
    state: JobState = JobState.None

    /**
     * 应用ID
     * @index ASC
     */
    appid: int64 = 0

    /**
     * 主机ID
     * @index ASC
     */
    sid: int64 = 0

    /**
     * 用户ID
     * @index ASC
     */
    uid: int64 = 0

    /**
     * 总任务数
     */
    maxCount: int32 = 0

    /**
     * 已执行任务数
     */
    count: int32 = 0

    /**
     * 错误任务数
     */
    errCount: int32 = 0

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
     * 开始时间
     * @index ASC
     */
    stime: int64 = 0

}