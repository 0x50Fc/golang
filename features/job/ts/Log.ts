import { int64 } from "./lib/less";

/**
 * 日志类型
 */
export enum LogType {
    /**
     * 信息
     */
    Info,
    /**
     * 调试信息
     */
    Debug,
    /**
     * 警告
     */
    Warn,
    /**
     * 错误
     */
    Error
}

/**
 * 日志
 * @type db
 */
export class Log {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 工作ID
     * @index ASC
     */
    jobId: int64 = 0

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
     * 类型
     * @index ASC
     */
    type: LogType = LogType.Info

    /**
     * 日志内容
     * @length -1
     */
    body: string = ''

    /**
     * 创建时间
     */
    ctime: int64 = 0

}