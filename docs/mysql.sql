create database goadmin default character set utf8mb4;
DROP TABLE IF EXISTS `go_role`;
CREATE TABLE `go_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '名字',
  `status` int(1) unsigned DEFAULT '0' COMMENT '状态值',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `go_role` VALUES ('1', '运维部','0', null, null, '0');
INSERT INTO `go_role` VALUES ('2', '开发部','0', null, null, '0');
INSERT INTO `go_role` VALUES ('3', '测试部','0', null, null, '0');


DROP TABLE IF EXISTS `go_user`;
CREATE TABLE `go_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号',
  `status` int(1) unsigned DEFAULT '0' COMMENT '状态值',
  `roleid` int(10) unsigned NOT NULL COMMENT '角色ID',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`),
  FOREIGN KEY (roleid) REFERENCES go_role(id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户管理';

INSERT INTO `go_user` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '15936007872', '0', '1',NULL,NULL,'0');
INSERT INTO `go_user` VALUES ('2', 'devlop', 'e10adc3949ba59abbe56e057f20f883e', NULL, '0', '2',NULL,NULL,'0');
INSERT INTO `go_user` VALUES ('3', 'test', 'e10adc3949ba59abbe56e057f20f883e', NULL, '0', '3',NULL,NULL,'0');



DROP TABLE IF EXISTS `go_menu`;
CREATE TABLE `go_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '路由名',
  `path` varchar(50) DEFAULT '' COMMENT '路由地址',
  `title` varchar(50) DEFAULT '' COMMENT '侧边栏显示名字',
  `icon` varchar(50) DEFAULT '' COMMENT '侧边栏显示图标',
  `componenturl` varchar(50) DEFAULT '' COMMENT '组件路径',
  `parentid` int(10) unsigned DEFAULT NULL COMMENT '父ID',
  `type` varchar(50) NOT NULL COMMENT '菜单类父ID型',
   `method` varchar(50)  DEFAULT NULL COMMENT '接口方法',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户管理';

INSERT INTO `go_menu` VALUES ('1','Sys','/sys','系统管理','sys','Layout',NULL,'M',NULL,NULL,NULL,'0');
INSERT INTO `go_menu` VALUES ('2','User','user','用户管理','user','user','1','C',NULL,NULL,NULL,'0');
INSERT INTO `go_menu` VALUES ('3','Role','role','角色管理','role','role','1','C',NULL,NULL,NULL,'0');
INSERT INTO `go_menu` VALUES ('4','Menu','menu','菜单管理','menu','menu','1','C',NULL,NULL,NULL,'0');
INSERT INTO `go_menu` VALUES ('5', '', '/api/v1/users', '查询用户列表', '', '', '2', 'J', 'GET', '1597997193', '0', NULL);
INSERT INTO `go_menu` VALUES ('6','','/api/v1/user/:id','查询单个用户','','','2','J','GET','1597997328','0',NULL);
INSERT INTO `go_menu` VALUES ('7','','/api/v1/users','添加用户','','','2','J','POST','1597997405','0',NULL);
INSERT INTO `go_menu` VALUES ('8','','/api/v1/users/:id','修改用户','','','2','J','PUT','1597997716','1598254875',NULL);
INSERT INTO `go_menu` VALUES ('9','','/api/v1/user/:id','重置密码','','','2','J','PUT','1597997769','0',NULL);
INSERT INTO `go_menu` VALUES ('10','','/api/v1/users/:id','删除用户','','','2','J','DELETE','1597997793','0',NULL);
INSERT INTO `go_menu` VALUES ('11','','/api/v1/roles','查询角色列表','','','3','J','GET','1598254670','1598254980',NULL);
INSERT INTO `go_menu` VALUES ('12','','/api/v1/role/:id','查询单个角色','','','3','J','GET','1598254742','0',NULL);
INSERT INTO `go_menu` VALUES ('13','','/api/v1/roles','添加角色','','','3','J','POST','1598254798','0',NULL);
INSERT INTO `go_menu` VALUES ('14','','/api/v1/roles/:id','修改角色','','','3','J','PUT','1598254828','0',NULL);
INSERT INTO `go_menu` VALUES ('15','','/api/v1/roles/:id','删除角色','','','3','J','DELETE','1598254861','0',NULL);
INSERT INTO `go_menu` VALUES ('16','','/api/v1/menus','查询菜单列表','','','4','J','GET','1598255018','0',NULL);
INSERT INTO `go_menu` VALUES ('17','','/api/v1/menu/:id','查看单个菜单','','','4','J','GET','1598255059','0',NULL);
INSERT INTO `go_menu` VALUES ('18','','/api/v1/menus','添加菜单','','','4','J','POST','1598255130','0',NULL);
INSERT INTO `go_menu` VALUES ('19','','/api/v1/menus/:id','修改菜单','','','4','J','PUT','1598255167','0',NULL);
INSERT INTO `go_menu` VALUES ('20','','/api/v1/menus/:id','删除菜单','','','4','J','DELETE','1598255202','0',NULL);


DROP TABLE IF EXISTS `go_role_menu`;
CREATE TABLE `go_role_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `roleid` int(10) unsigned NOT NULL COMMENT '角色ID',
  `menuid` int(10) unsigned NOT NULL COMMENT '菜单ID',
  `created_on` int(11) unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int(11) unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`),
  FOREIGN KEY (roleid) REFERENCES go_role(id),
  FOREIGN KEY (menuid) REFERENCES go_menu(id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色菜单关联';

INSERT INTO `go_role_menu` VALUES ('1','1','1','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('2','1','2','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('3','1','3','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('4','1','4','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('5','1','5','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('6','1','6','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('7','1','7','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('8','1','8','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('9','1','9','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('10','1','10','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('11','1','11','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('12','1','12','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('13','1','13','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('14','1','14','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('15','1','15','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('16','1','16','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('17','1','17','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('18','1','18','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('19','1','19','0','0',NULL);
INSERT INTO `go_role_menu` VALUES ('20','1','20','0','0',NULL);


DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL COMMENT '类型',
  `v0` varchar(100) DEFAULT NULL COMMENT '部门',
  `v1` varchar(100) DEFAULT NULL COMMENT 'path',
  `v2` varchar(100) DEFAULT NULL COMMENT 'method',
  `v3` varchar(100) DEFAULT NULL COMMENT '',
  `v4` varchar(100) DEFAULT NULL COMMENT '',
  `v5` varchar(100) DEFAULT NULL COMMENT ''
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='casbin_rule接口权限';

INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/users','GET',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/user/:id','GET',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/users','POST',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/users/:id','PUT',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/user/:id','PUT',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/users/:id','DELETE',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/roles','GET',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/role/:id','GET',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/roles','POST',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/roles/:id','PUT',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/roles/:id','DELETE',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/menus','GET',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/menu/:id','GET',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/menus','POST',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/menus/:id','PUT',NULL,NULL,NULL);
INSERT INTO `casbin_rule` VALUES ('p','运维部','/api/v1/menus/:id','DELETE',NULL,NULL,NULL);
