import { int64 } from "./lib/less";
import { UserType } from './UserType';

/**
 * 开发平台
 * @type db
 */
export class Open {

    /**
     * ID
     */
    id: int64 = 0

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
     * access_token
     * @length 255
     * @output false
     */
    access_token: string = ''

    /**
     * refresh_token
     * @length 255
     * @output false
     */
    refresh_token: string = ''

    /**
     * 过期时间
     */
    etime: int64 = 0

    /**
     * 其他数据
     * @length -1
     */
    options: any


}