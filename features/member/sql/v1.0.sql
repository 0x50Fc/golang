CREATE TABLE IF NOT EXISTS `member_member` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`bid` BIGINT DEFAULT 0	#[字段] 商户ID
	,`uid` BIGINT DEFAULT 0	#[字段] 成员ID
	,`title` VARCHAR(255) DEFAULT ''	#[字段] 备注名
	,`keyword` VARCHAR(2048) DEFAULT ''	#[字段] 搜索关键字
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `bid` (`bid` ASC)	#[索引] 商户ID
	,INDEX `uid` (`uid` ASC)	#[索引] 成员ID
 ) AUTO_INCREMENT = 1;
