import { int64, int32 } from './lib/less';

/**
 * 统计记录
 */
export class Addup {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 统计项ID
     * @index desc
     */
    iid: int64 = 0

    /**
     * 时间 唯一
     * @index desc
     */
    time: int64 = 0

}