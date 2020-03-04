import { int64 } from "./lib/less";

/**
 * 从主机状态
 */
export enum SlaveState {
    /**
     * 等待启动
     */
    None,

    /**
     * 已启动
     */
    Running,
    /**
     * 已退出
     */
    Exit,

    /**
     * 已超时
     */
    Timeout
}

/**
 * Slave 主机
 * @type db
 */
export class Slave {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 别名前缀
     * @length 128
     */
    prefix: string = ''

    /**
     * 授权token
     * @index ASC
     * @length 32
     */
    token: string = ''

    /**
     * 主机状态
     * @index ASC
     */
    state: SlaveState = SlaveState.None

    /**
     * 超时时间
     * @index DESC
     */
    etime: int64 = 0

    /**
     * 其他数据
     * @length -1
     */
    options: any

    /**
     * 创建时间
     */
    ctime: int64 = 0

}