
import { BaseResponse, ErrCode } from "../../../lib/BaseResponse"
import { int64, int32 } from "../../../lib/less";
import { Authority } from '../../../Authority';

/**
 * 角色设置资源
 * @method POST
 */
interface Request {

    /**
     * 角色名
     */
    role: string

    /**
     * 资源路径, 多个逗号分割
     */
    res: string


}


interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
