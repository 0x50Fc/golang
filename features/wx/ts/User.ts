import { int64, int32 } from "./lib/less";
import { UserType } from './UserType';

export enum UserState {
    Subscribe = 0,
    UnSubscribe = 1
}

/**
 * 用户
 * @type db
 */
export class User {

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
     * openid
     * @index ASC
     * @length 128
     */
    openid: string = ''

    /**
     * unionid
     * @index ASC
     * @length 128
     */
    unionid: string = ''

    /**
     * access_token
     * @length 255
     */
    access_token: string = ''

    /**
     * refresh_token
     * @length 255
     */
    refresh_token: string = ''

    /**
     * session_key
     * @length 128
     */
    session_key: string = ''

    /**
     * 昵称
     * @length 255
     */
    nick: string = ""

    /**
     * 头像
     * @length 2048
     */
    logo: string = ""

    /**
     * 国家
     * @length 64
     */
    country: string = ""

    /**
     * 语言
     * @length 64
     */
    lang: string = ""

    /**
     * 省份
     * @length 64
     */
    province: string = ""

    /**
     * 城市
     * @length 64
     */
    city: string = ""

    /**
     * 性别
     */
    gender: int32 = 0

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
     * 过期时间
     */
    etime: int64 = 0

    /**
     * 关注状态
     */
    state: UserState = UserState.Subscribe

    /**
     * 最后绑定时间
     */
    mtime: int64 = 0
}
