import { int64, int32 } from './lib/less';

/**
 * Top
 * @type db
 */
export class Top {

    /**
     * ID
     */
    id: int64 = 0

    /**
     * 目标ID
     * @index ASC
     */
    tid: int64 = 0

    /**
     * 搜索关键字
     * @length 2048
     */
    keyword: string = ''

    /**
     * 序号 降序
     * @index DESC
     */
    sid: int64 = 0

    /**
     * 排名
     * @index ASC
     */
    rank: int32 = 0

    /**
     * 固定排名位置
     */
    fixed: int32 = 0;

    /**
     * 其他数据
     * @length -1
     */
    options: any

}