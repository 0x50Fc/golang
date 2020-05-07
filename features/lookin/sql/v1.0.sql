CREATE TABLE IF NOT EXISTS `lookin_1_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_2_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_3_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_4_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_5_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_6_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_7_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_8_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_9_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_10_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_11_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_12_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_13_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_14_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_15_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
CREATE TABLE IF NOT EXISTS `lookin_16_lookin` (
	id BIGINT NOT NULL AUTO_INCREMENT	#ID
	,`tid` BIGINT DEFAULT 0	#[字段] 目标ID
	,`iid` BIGINT DEFAULT 0	#[字段] 项ID
	,`uid` BIGINT DEFAULT 0	#[字段] 用户ID
	,`fuid` BIGINT DEFAULT 0	#[字段] 好友ID
	,`fcode` VARCHAR(64) DEFAULT ''	#[字段] 好友推荐码
	,`flevel` INT DEFAULT 0	#[字段] 关系级别
	,`options` JSON DEFAULT NULL	#[字段] 其他数据
	,`ctime` BIGINT DEFAULT 0	#[字段] 创建时间
	, PRIMARY KEY(id) 
	,INDEX `tid` (`tid` ASC)	#[索引] 目标ID
	,INDEX `iid` (`iid` ASC)	#[索引] 项ID
	,INDEX `uid` (`uid` ASC)	#[索引] 用户ID
	,INDEX `fuid` (`fuid` ASC)	#[索引] 好友ID
	,INDEX `flevel` (`flevel` ASC)	#[索引] 关系级别
 ) AUTO_INCREMENT = 1;
