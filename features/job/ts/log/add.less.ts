
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { App } from '../App';
import { LogType, Log } from '../Log';

/**
 * 添加日志
 * @method POST
 */
interface Request {

    /**
     * 工作ID
     */
    jobId: int64

    /**
     * 应用ID
     */
    appid: int64

    /**
     * 主机ID
     */
    sid: int64

    /**
     * 类型
     */
    type: LogType

    /**
     * 日志内容
     */
    body: string

}

interface Response extends BaseResponse {
    data?: Log
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
