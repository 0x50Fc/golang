import { int64 } from "./lib/less";
import { UserType } from './UserType';

/**
 * Token
 * @type db
 */
export class Token {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 类型
     * @index ASC
     */
    type: UserType = UserType.MP

    /**
     * appid
     * @index ASC
     * @length 64
     */
    appid: string = ''

    /**
     * access_token
     * @length 255
     * @output false
     */
    access_token: string = ''

    /**
     * 过期时间
     */
    etime: int64 = 0

}