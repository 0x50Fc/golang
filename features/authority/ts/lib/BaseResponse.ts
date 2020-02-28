import { int32 } from "./less";


export enum ErrCode {
    OK = 200
}

export interface BaseResponse {
    /**
     * 错误码
     */
    errno: int32
    /**
     * 错误信息
     */
    errmsg?: string
}
