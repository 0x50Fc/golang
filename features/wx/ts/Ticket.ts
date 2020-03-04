import { int64 } from "./lib/less";
import { UserType } from './UserType';

export enum TicketType {
    Jsapi = "jsapi",
    Card = "wx-card"
}

/**
 * Ticket
 * @type db
 */
export class Ticket {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 类型
     * @index ASC
     * @length 32
     */
    type: TicketType = TicketType.Jsapi

    /**
     * appid
     * @index ASC
     * @length 64
     */
    appid: string = ''

    /**
     * ticket
     * @length 255
     * @output false
     */
    ticket: string = ''

    /**
     * 过期时间
     */
    etime: int64 = 0

}