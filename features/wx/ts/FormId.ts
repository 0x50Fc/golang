import { int64 } from "./lib/less";
import { UserType } from './UserType';

/**
 * FormId
 * @type db
 */
export class FormId {

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
     * openid
     * @index ASC
     * @length 128
     */
    openid: string = ''

    /**
     * formid
     * @length 128
     */
    formid: string = ''

    /**
     * 过期时间
     * @index DESC
     */
    etime: int64 = 0

}