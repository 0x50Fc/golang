import { int64 } from "./lib/less";

/**
 * 用户
 * @type db
 */
export class User {

    /**
     * 用户ID
     */
    id: int64 = 0

    /**
     * 用户名
     * @length 128
     * @unique ASC
     */
    name: string = ""

    /**
     * 昵称
     * @length 128
     * @index ASC
     */
    nick: string = ""

    /**
     * 密码
     * @length 32
     * @output false
     */
    password: string = ""

    /**
     * 创建时间
     */
    ctime: int64 = 0

}