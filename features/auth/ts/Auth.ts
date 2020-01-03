import { int64, int32 } from "./lib/less";

export enum AuthType {
    JSON = "json",
    Text = "text"
}

/**
 * 验证
 * @type db
 */
export class Auth {

    /**
    * 用户ID
    */
    id: int64 = 0

    /**
     * 唯一键
     * @unique ASC
     * @length 128
     */
    key: string = ""

    /**
     * 值
     * @length -1
     */
    value: string = ""

    /**
     * 失效时间
     */
    etime: int64 = 0

}