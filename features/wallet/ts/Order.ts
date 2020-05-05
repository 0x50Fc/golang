import { int64 } from "./lib/less";

export enum OrderState {
    None = 0,
    Freeze = 1,
    OK = 2,
    Canceled = 3,
    Rollback = 4
}

/**
 * 订单
 * @type db
 */
export class Order {

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
     * 状态
     * @index ASC
     */
    state: OrderState = OrderState.None

    /**
     * 其他数据
     */
    options: any

    /**
     * 是否冻结
     */
    freeze: boolean = false

    /**
     * 创建时间
     */
    ctime: int64 = 0

}

/**
 * 订单明细
 * @type db
 */
export class OrderItem {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 订单ID
     * @index ASC
     */
    orderId: int64 = 0

    /**
     * 用户ID
     */
    uid: int64 = 0

    /**
     * 钱包ID
     */
    walletId: int64 = 0

    /**
     * 金额 正值收入/负值支出
     */
    value: int64 = 0

    /**
     * 状态
     * @index ASC
     */
    state: OrderState = OrderState.None
}
