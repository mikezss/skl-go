CREATE TABLE cmn_admingroup_tb (
  adminid varchar(50) NOT NULL,
  groupid varchar(50) NOT NULL,
  PRIMARY KEY (adminid,groupid)
);
CREATE TABLE cmn_adminorg_tb (
  adminid varchar(50) NOT NULL,
  orgid varchar(50) NOT NULL,
  PRIMARY KEY (adminid,orgid)
);
CREATE TABLE cmn_adminrole_tb (
  adminid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (adminid,roleid)
);

CREATE TABLE cmn_flowaction_tb (
  flowid varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  actionid varchar(50) NOT NULL DEFAULT '',
  nexttaskid varchar(50) DEFAULT NULL,
  backtotaskid varchar(50) DEFAULT NULL,
  taskstatus varchar(50) DEFAULT NULL,
  dispatcher varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowid,taskid,actionid)
);
CREATE TABLE cmn_flowsplit_tb (
  flowid varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  switcherid varchar(50) NOT NULL DEFAULT '',
  varyid varchar(50) DEFAULT NULL,
  functionname varchar(50) DEFAULT NULL,
  functionvalue varchar(50) DEFAULT NULL,
  nexttaskid varchar(50) DEFAULT NULL,
  taskstatus varchar(50) DEFAULT NULL,
  dispatcher varchar(50) DEFAULT NULL,
  taskname varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowid,taskid,switcherid)
);
CREATE TABLE cmn_flowtask_tb (
  flowid varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  taskname varchar(50) DEFAULT NULL,
  flowtype varchar(50) DEFAULT NULL,
  taskjump varchar(50) DEFAULT NULL,
  taskjumpinfer varchar(50) DEFAULT NULL,
  executorconcurrent varchar(50) DEFAULT NULL,
  samepersontaskid varchar(50) DEFAULT NULL,
  nobodytaskid varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowid,taskid)
);
CREATE TABLE cmn_flowuser_tb (
  flowid varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  executorid varchar(50) NOT NULL DEFAULT '',
  executortype varchar(50) DEFAULT NULL,
  executor varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowid,taskid,executorid)
);
CREATE TABLE cmn_flowvar_tb (
  flowid varchar(50) NOT NULL,
  varyid varchar(50) NOT NULL,
  varyname varchar(50) DEFAULT NULL,
  varyvalue varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowid,varyid)
);
CREATE TABLE cmn_grouplevel_tb (
  upperid varchar(50) NOT NULL,
  lowerid varchar(50) NOT NULL,
  PRIMARY KEY (upperid,lowerid)
);
CREATE TABLE cmn_grouprole_tb (
  groupid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (groupid,roleid)
);
CREATE TABLE cmn_orgleader_tb (
  orgid varchar(50) NOT NULL,
  userid varchar(50) NOT NULL,
  leadertype char(1) DEFAULT NULL,
  PRIMARY KEY (orgid,userid)
);
CREATE TABLE cmn_orglevel_tb (
  upperid varchar(50) NOT NULL,
  lowerid varchar(50) NOT NULL,
  PRIMARY KEY (upperid,lowerid)
);
CREATE TABLE cmn_orgrole_tb (
  orgid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (orgid,roleid)
);
CREATE TABLE cmn_roleprivilege_tb (
  roleid varchar(50) NOT NULL,
  modualid varchar(50) NOT NULL,
  PRIMARY KEY (roleid,modualid)
);
CREATE TABLE cmn_usergroup_tb (
  userid varchar(50) NOT NULL,
  groupid varchar(50) NOT NULL,
  expireddate date DEFAULT NULL,
  PRIMARY KEY (userid,groupid)
);
CREATE TABLE cmn_userrole_tb (
  userid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (userid,roleid)
);

CREATE TABLE fi_agent_tb (
  userid varchar(50) DEFAULT NULL,
  agent varchar(50) DEFAULT NULL,
  startdate date DEFAULT NULL,
  enddate date DEFAULT NULL
);
CREATE TABLE fi_countersign (
  tiid int(11) NOT NULL,
  userid varchar(50) NOT NULL,
  fiid int(11) NOT NULL DEFAULT '0',
  taskid varchar(255) NOT NULL DEFAULT '',
  taskstatus varchar(255) NOT NULL DEFAULT '0',
  PRIMARY KEY (tiid,userid)
);
CREATE TABLE fi_flowmantaskaction_tb (
  flowtemplateid varchar(50) NOT NULL,
  tasktype varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  action varchar(50) NOT NULL,
  jump varchar(50) NOT NULL,
  status varchar(50) DEFAULT NULL,
  nexttask varchar(50) DEFAULT NULL,
  backtask varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowtemplateid,taskid,action)
);
CREATE TABLE fi_flowswitchtaskaction_tb (
  flowtemplateid varchar(50) NOT NULL,
  tasktype varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  nos varchar(50) NOT NULL,
  conditions varchar(50) NOT NULL,
  functions varchar(50) NOT NULL,
  valuee varchar(50) NOT NULL,
  jump varchar(50) NOT NULL,
  statuss varchar(50) DEFAULT NULL,
  nexttask varchar(50) DEFAULT NULL,
  backtask varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowtemplateid,taskid,nos)
);
CREATE TABLE fi_flowtask_tb (
  flowtemplateid varchar(50) NOT NULL,
  tasktype varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  taskname varchar(50) NOT NULL,
  supportskip tinyint(1) NOT NULL DEFAULT '0',
  sendmessage tinyint(1) NOT NULL DEFAULT '0',
  concurrent varchar(50) NOT NULL,
  samepersontask varchar(50) DEFAULT NULL,
  nopersontask varchar(50) DEFAULT NULL,
  PRIMARY KEY (flowtemplateid,taskid)
);
CREATE TABLE fi_flowtaskexecuter_tb (
  flowtemplateid varchar(50) NOT NULL,
  tasktype varchar(50) NOT NULL,
  taskid varchar(50) NOT NULL,
  no varchar(50) NOT NULL,
  taskexecuter varchar(50) NOT NULL,
  expression varchar(50) NOT NULL,
  PRIMARY KEY (flowtemplateid,taskid,no)
);
CREATE TABLE fi_org_vary (
  orgid varchar(50) NOT NULL,
  vid varchar(50) NOT NULL,
  vvalue int(11) DEFAULT '0',
  PRIMARY KEY (orgid,vid)
);
CREATE TABLE fi_owner (
  tiid bigint(20) NOT NULL,
  owner varchar(50) NOT NULL,
  PRIMARY KEY (tiid,owner)
);
CREATE TABLE fi_var (
  fiid int(11) NOT NULL,
  vid varchar(50) NOT NULL,
  vvalue varchar(50) DEFAULT NULL,
  PRIMARY KEY (fiid,vid)
);
CREATE TABLE hr_leave_month (
  userid varchar(50) NOT NULL,
  years char(4) NOT NULL,
  remainsalaryleave double DEFAULT NULL,
  usedsalaryleave double DEFAULT NULL,
  remainwelfareleave double DEFAULT NULL,
  usedwelfareleave double DEFAULT NULL,
  remainsickleave double DEFAULT NULL,
  usedsickleave double DEFAULT NULL,
  lastyearremainleave double DEFAULT NULL,
  usedlastyearremainleave double DEFAULT NULL,
  PRIMARY KEY (userid,years)
);
CREATE TABLE hr_leave_year (
  userid varchar(50) NOT NULL,
  years char(4) NOT NULL,
  Currentyearsalaryleave double DEFAULT NULL,
  currentyearwelfareleave double DEFAULT NULL,
  currentyearsickleave double DEFAULT NULL,
  lastyearremainleave double DEFAULT NULL,
  lastyearwanderleave double DEFAULT NULL,
  PRIMARY KEY (userid,years)
);
CREATE TABLE hr_paidleave_month (
  userid varchar(50) NOT NULL,
  years char(4) NOT NULL,
  months varchar(2) NOT NULL,
  overtimeobtained double DEFAULT NULL,
  leaveused double DEFAULT NULL,
  leavedcanceled double DEFAULT NULL,
  invalid double DEFAULT NULL,
  limitdate date DEFAULT NULL,
  PRIMARY KEY (userid,years,months)
);
CREATE TABLE sequence (
  seqname varchar(50) NOT NULL,
  currentValue int(11) NOT NULL,
  increment int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (seqname)
);
INSERT INTO cmn_org_tb VALUES ('root', '', '', null, -1, null, null);
INSERT INTO cmn_user_tb VALUES ('devzss', 'devzss', null, 0, 666666, '', '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 0, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_userrole_tb VALUES ('devzss', 'super');
INSERT INTO cmn_role_tb VALUES ('root', 'root', '', -1, null);
INSERT INTO cmn_role_tb VALUES ('super', '超级管理员', 'root', 0, null);
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'passwordchange');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'systemmanage');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'administrator');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'companymanage');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'degree');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'employeetype');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'lang');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'modualmanage');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'orgmanage');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'orgtype');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'post');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'rolemanage');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'usergroup');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'usermanage');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'usertype');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'create-component');

INSERT INTO cmn_modual_tb VALUES ('root', 'root', '', null, null);
INSERT INTO cmn_modual_tb VALUES (0, 'managesystem', 'root', '', '');
INSERT INTO cmn_modual_tb VALUES ('administrator', '管理员', 'systemmanage', '/admin', '');
INSERT INTO cmn_modual_tb VALUES ('agent', '代理设置', 'flowmanagement', '/agent', '');
INSERT INTO cmn_modual_tb VALUES ('companymanage', '公司管理', 'systemmanage', '/company', '');
INSERT INTO cmn_modual_tb VALUES ('degree', '学历管理', 'systemmanage', '/degree', '');
INSERT INTO cmn_modual_tb VALUES ('donetask', '已办任务', 'flowmanagement', '/donetask', '');
INSERT INTO cmn_modual_tb VALUES ('employeetype', '职员类型', 'systemmanage', '/employeetype', '');
INSERT INTO cmn_modual_tb VALUES ('flow-monitor', '流程监控', 'flowmanagement', '/flow-monitor', '');
INSERT INTO cmn_modual_tb VALUES ('flowdefine', '流程定义', 'flowmanagement', '/flowdefine', '');
INSERT INTO cmn_modual_tb VALUES ('flowmanagement', '流程管理', 0, '', '');
INSERT INTO cmn_modual_tb VALUES ('floworgvary', '流程变量管理', 'flowmanagement', '/floworgvary', '');
INSERT INTO cmn_modual_tb VALUES ('flowstatus', '流程状态管理', 'flowmanagement', '/flowstatus', '');
INSERT INTO cmn_modual_tb VALUES ('flowtemplate', '流程模板', 'flowmanagement', '/flowtemplate', '');
INSERT INTO cmn_modual_tb VALUES ('lang', '多语言设置', 'systemmanage', '/lang', '');
INSERT INTO cmn_modual_tb VALUES ('modualmanage', '模块管理', 'systemmanage', '/modual', '');
INSERT INTO cmn_modual_tb VALUES ('myflow', '我的流程', 'flowmanagement', '/myflow', '');
INSERT INTO cmn_modual_tb VALUES ('orgmanage', '机构管理', 'systemmanage', '/org', '');
INSERT INTO cmn_modual_tb VALUES ('orgtype', '机构类型', 'systemmanage', '/orgtype', '');
INSERT INTO cmn_modual_tb VALUES ('orgvary', '机构变量', 'flowmanagement', '/orgvary', '');
INSERT INTO cmn_modual_tb VALUES ('passwordchange', '密码变更', 'commonmanage', '/passwordchange', '');
INSERT INTO cmn_modual_tb VALUES ('post', '职位管理', 'systemmanage', '/post', '');
INSERT INTO cmn_modual_tb VALUES ('rolemanage', '角色管理', 'systemmanage', '/role', '');
INSERT INTO cmn_modual_tb VALUES ('systemmanage', '系统管理', 0, '', '');
INSERT INTO cmn_modual_tb VALUES ('task-trace', '流程追踪', 'flowmanagement', '/task-trace', '');
INSERT INTO cmn_modual_tb VALUES ('todo', '待办任务', 'flowmanagement', '/todo', '');
INSERT INTO cmn_modual_tb VALUES ('transferflow', '转岗&离岗&转签', 'flowmanagement', '/transferflow', '');
INSERT INTO cmn_modual_tb VALUES ('usergroup', '用户组管理', 'systemmanage', '/usergroup', '');
INSERT INTO cmn_modual_tb VALUES ('userinfo', '员工信息', 'hrmanage', '/userinfo', '');
INSERT INTO cmn_modual_tb VALUES ('usermanage', '用户管理', 'systemmanage', '/usermanage', '');
INSERT INTO cmn_modual_tb VALUES ('usertype', '用户类型', 'systemmanage', '/usertype', '');


INSERT INTO fi_template_tb VALUES ('expense', '费用报销', '费用报销', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('plan', '企画申请', '企画申请', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('pay', '支付申请', '支付申请', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('loan', '借款申请', '借款申请', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('loantravel', '出差申请', '出差申请', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('common', '通用流程模板', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('common2', '通用流程模板无财务', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('meeting', '会议费申请', '会议费申请', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('travel', '差旅费报销', '费用报销', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('payaccount', '支付账户申请', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('att', '请假申请', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('attcancel', '请假取消申请', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('overtime', '加班申请', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('resign', '补签考勤申请', '通用流程模板', 'flowinstid', 'flowstatus');
INSERT INTO fi_template_tb VALUES ('meetingroomapply', '会议室申请', '通用流程模板', 'flowinstid', 'flowstatus');

INSERT INTO fi_templateitem_tb VALUES (5, 'expense', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (6, 'expense', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (7, 'expense', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (11, 'plan', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (12, 'plan', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (13, 'plan', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (17, 'pay', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (18, 'pay', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (19, 'pay', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (29, 'loan', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (30, 'loan', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (31, 'loan', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (38, 'loantravel', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (39, 'loantravel', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (40, 'loantravel', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (41, 'loantravel', 'scopeflag', '国内/国外', 2, '');
INSERT INTO fi_templateitem_tb VALUES (42, 'common', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (43, 'common', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (44, 'common', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (45, 'common2', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (46, 'common2', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (47, 'common2', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (48, 'meeting', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (49, 'meeting', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (50, 'meeting', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (51, 'travel', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (52, 'travel', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (53, 'travel', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (54, 'payaccount', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (55, 'payaccount', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (56, 'payaccount', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (57, 'att', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (58, 'att', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (59, 'att', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (60, 'attcancel', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (61, 'attcancel', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (62, 'attcancel', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (70, 'overtime', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (71, 'overtime', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (72, 'overtime', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (73, 'overtime', 'exceeddays', '超过日期', 2, '');
INSERT INTO fi_templateitem_tb VALUES (74, 'overtime', 'totalovertimedays', '总加班日数', 2, '');
INSERT INTO fi_templateitem_tb VALUES (75, 'resign', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (76, 'resign', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (77, 'resign', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (83, 'meetingroomapply', 'money', '金额', 2, '');
INSERT INTO fi_templateitem_tb VALUES (84, 'meetingroomapply', 'orglevel', '机构级别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (85, 'meetingroomapply', 'vicemanagertype', '副总类别', 2, '');
INSERT INTO fi_templateitem_tb VALUES (86, 'meetingroomapply', 'schedule', '是否需要审批', 2, '');

INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 30, '部长审批', 0, 0, '', 32, 32);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 35, '本部长审批', 0, 0, '', 40, 40);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 50, '总经理审批', 0, 0, '', 55, 55);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 60, '财务课长审批', 0, 0, '', '', 20);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 70, '财务部长审批', 0, 0, '', '', 75);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 75, '出纳审批', 0, 0, '', '', 80);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 80, '出纳发款', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'switch', 47, '判断副总提交', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 42, '营业副总审批', 0, 0, '', 47, 47);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'switch', 40, '判断副总审批', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('expense', 'switch', 32, '判断机构级别', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 45, '售后副总审批', 0, 0, '', 47, 47);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 55, '财务担当审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 65, '财务副部长审批', 0, 0, '', 70, 70);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 30, '部长审批', 0, 0, '', 32, 32);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'switch', 32, '判断机构级别', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 35, '本部长审批', 0, 0, '', 40, 40);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'switch', 40, '判断副总审批', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 42, '营业副总审批', 0, 0, '', 47, 47);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 45, '售后副总审批', 0, 0, '', 47, 47);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'switch', 47, '判断副总提交', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 50, '总经理审批', 0, 0, '', 55, 55);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 55, '财务担当审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 60, '财务课长审批', 0, 0, '', '', 20);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 65, '财务副部长审批', 0, 0, '', 70, 70);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 70, '财务部长审批', 0, 0, '', '', 75);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 75, '出纳审批', 0, 0, '', '', 80);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 80, '出纳发款', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 30, '部长审批', 0, 0, '', 32, 32);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 80, '出纳发款', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'switch', 40, '判断副总审批', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'switch', 47, '判断副总提交', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 50, '总经理审批', 0, 0, '', 55, 55);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 55, '财务担当审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 60, '财务课长审批', 0, 0, '', '', 20);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 65, '财务副部长审批', 0, 0, '', 70, 70);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 70, '财务部长审批', 0, 0, '', '', 75);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 75, '出纳审批', 0, 0, '', '', 80);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 80, '出纳发款', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'switch', 33, '判断金额2', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'switch', 32, '判断金额1', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 30, '部长审批', 0, 0, '', 34, 34);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'switch', 34, '判断机构级别', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 42, '营业副总审批', 0, 0, '', 46, 46);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 45, '售后副总审批', 0, 0, '', 46, 46);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'switch', 46, '判断金额3', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 37, '本部长2审批', 0, 0, '', 33, 33);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 35, '本部长审批', 0, 0, '', 40, 40);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 50, '总经理审批', 0, 0, '', 55, 55);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 55, '财务担当审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 60, '财务课长审批', 0, 0, '', '', 20);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 70, '财务部长审批', 0, 0, '', '', 75);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 75, '出纳审批', 0, 0, '', '', 80);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 80, '出纳发款', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 30, '部长审批', 0, 0, '', 35, 35);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 35, '本部长审批', 0, 0, '', 45, 45);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 45, '副总审批', 0, 0, '', 50, 50);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 65, '财务副部长审批', 0, 0, '', '', 70);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 30, '部长审批', 0, 0, '', 32, 80);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('loan', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('expense', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'switch', 40, '判断金额', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 50, '综合企画课长审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 60, '综合企画部长审批', 0, 0, '', '', 70);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 70, '本部长审批', 0, 0, '', '', 80);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'switch', 85, '判断金额2', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 80, '副总审批', 0, 0, '', '', 85);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 90, '总经理审批', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 30, '部长审批', 0, 0, '', 32, 40);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('plan', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'switch', 32, '判断机构级别', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 35, '本部长审批', 0, 0, '', 5, 45);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 45, '副总审批', 0, 0, '', 47, 50);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'switch', 50, '判断国内/国外', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('meeting', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('pay', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('common', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 30, '部长审批', 0, 0, '', 35, 35);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 35, '本部长审批', 0, 0, '', 45, 45);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 45, '副总审批', 0, 0, '', 50, 50);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('common2', 'man', 50, '总经理审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('loantravel', 'man', 60, '总经理审批', 0, 0, '', 47, 999);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 30, '部长审批', 0, 0, '', 32, 32);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'switch', 32, '判断机构级别', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 35, '本部长审批', 0, 0, '', 40, 40);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'switch', 40, '判断副总审批', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 42, '营业副总审批', 0, 0, '', 47, 47);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 45, '售后副总审批', 0, 0, '', 47, 47);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'switch', 47, '判断副总提交', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 50, '总经理审批', 0, 0, '', 55, 55);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 55, '财务担当审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 60, '财务课长审批', 0, 0, '', '', 20);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 65, '财务副部长审批', 0, 0, '', 70, 70);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 70, '财务部长审批', 0, 0, '', '', 75);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 75, '出纳审批', 0, 0, '', '', 80);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 80, '出纳发款', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('travel', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 10, '代理课长审批', 0, 0, '', 15, 15);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 15, '课长审批', 0, 0, '', 20, 20);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 20, '副部长审批', 0, 0, '', 25, 25);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 25, '代理部长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 5, '副课长审批', 0, 0, 0, 10, 10);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 65, '财务副部长审批', 0, 0, '', '', 70);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 30, '部长审批', 0, 0, '', 35, 40);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 40, '财务担当审批', 0, 0, '', '', 50);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 50, '出纳审批', 0, 0, '', '', 60);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 70, '财务部长审批', 0, 0, '', '', 999);
INSERT INTO fi_flowtask_tb VALUES ('payaccount', 'man', 60, '财务课长审批', 0, 0, '', '', 65);
INSERT INTO fi_flowtask_tb VALUES ('att', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('att', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('att', 'man', 15, '课长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('att', 'man', 30, '部长审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('attcancel', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('attcancel', 'man', 15, '课长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('attcancel', 'man', 30, '部长审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('attcancel', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'man', 30, '部长审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'switch', 5, '判断当月加班天数', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'man', 50, '总经理审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'switch', 10, '判断加班日期', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('overtime', 'man', 40, '副总审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('resign', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('resign', 'man', 15, '课长审批', 0, 0, '', 30, 30);
INSERT INTO fi_flowtask_tb VALUES ('resign', 'man', 30, '部长审批', 0, 0, '', 999, 999);
INSERT INTO fi_flowtask_tb VALUES ('resign', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('meetingroomapply', 'man', 1, '启动', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('meetingroomapply', 'man', 999, '结束', 0, 0, 0, '', '');
INSERT INTO fi_flowtask_tb VALUES ('meetingroomapply', 'switch', 5, '判断是否需要审批', 0, 0, '', '', '');
INSERT INTO fi_flowtask_tb VALUES ('meetingroomapply', 'man', 10, '会议室管理者审批', 0, 0, '', '', 999);

INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 60, 1, 3, 'finance_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 70, 1, 3, 'finance_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 75, 1, 3, 'chuna1');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 80, 1, 3, 'chuna2');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 42, 1, 3, 'business_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 45, 1, 3, 'aftersale_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 55, 1, 3, 'finance_employee');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 65, 1, 3, 'finance_viceminister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 42, 1, 3, 'business_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 45, 1, 3, 'aftersale_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 55, 1, 3, 'finance_employee');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 60, 1, 3, 'finance_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 65, 1, 3, 'finance_viceminister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 70, 1, 3, 'finance_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 75, 1, 3, 'chuna1');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 80, 1, 3, 'chuna2');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 80, 1, 3, 'chuna2');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 55, 1, 3, 'finance_employee');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 60, 1, 3, 'finance_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 65, 1, 3, 'finance_viceminister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 70, 1, 3, 'finance_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 75, 1, 3, 'chuna1');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 80, 1, 3, 'chuna2');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 42, 1, 3, 'business_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 45, 1, 3, 'aftersale_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 37, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 55, 1, 3, 'finance_employee');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 60, 1, 3, 'finance_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 70, 1, 3, 'finance_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 75, 1, 3, 'chuna1');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 80, 1, 3, 'chuna2');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 45, 1, 7, 1);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 65, 1, 3, 'finance_viceminister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loan', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('expense', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 50, 1, 3, 'comprehensiveplan_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 60, 1, 3, 'comprehensiveplan_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 70, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 80, 1, 7, 1);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 90, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('plan', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 45, 1, 7, 1);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meeting', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('pay', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 45, 1, 7, 1);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('common2', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('loantravel', 'man', 60, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 35, 1, 6, 2);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 42, 1, 3, 'business_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 45, 1, 3, 'aftersale_vicegm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 55, 1, 3, 'finance_employee');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 60, 1, 3, 'finance_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 65, 1, 3, 'finance_viceminister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 70, 1, 3, 'finance_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 75, 1, 3, 'chuna1');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('travel', 'man', 80, 1, 3, 'chuna2');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 10, 1, 6, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 20, 1, 7, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 25, 1, 6, 4);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 5, 1, 7, 6);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 65, 1, 3, 'finance_viceminister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 40, 1, 3, 'finance_employee');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 50, 1, 3, 'chuna1');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 70, 1, 3, 'finance_minister');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('payaccount', 'man', 60, 1, 3, 'finance_director');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('att', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('att', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('attcancel', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('attcancel', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('overtime', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('overtime', 'man', 50, 1, 3, 'gm');
INSERT INTO fi_flowtaskexecuter_tb VALUES ('overtime', 'man', 40, 1, 7, 1);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('resign', 'man', 15, 1, 6, 5);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('resign', 'man', 30, 1, 6, 3);
INSERT INTO fi_flowtaskexecuter_tb VALUES ('meetingroomapply', 'man', 10, 1, 3, 'meetingroomverify');

INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 47, 1, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 47, 2, 'orglevel', 2, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 47, 3, 'orglevel', 3, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 40, 1, 'vicemanagertype', 4, 1, 'forward', '', 45, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 40, 2, 'vicemanagertype', 4, 2, 'forward', '', 42, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 32, 1, 'orglevel', 4, 0, 'forward', '总经理提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 32, 2, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 32, 3, 'orglevel', 4, 2, 'forward', '本部长提交', 40, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 32, 4, 'orglevel', 4, 3, 'forward', '部长提交', 35, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('expense', 'switch', 32, 5, 'orglevel', 0, 4, 'forward', '部长以下提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 32, 1, 'orglevel', 4, 0, 'forward', '总经理提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 32, 2, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 32, 3, 'orglevel', 4, 2, 'forward', '本部长提交', 40, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 32, 4, 'orglevel', 4, 3, 'forward', '部长提交', 35, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 32, 5, 'orglevel', 0, 4, 'forward', '部长以下提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 40, 1, 'vicemanagertype', 4, 1, 'forward', '', 45, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 40, 2, 'vicemanagertype', 4, 2, 'forward', '', 42, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 47, 1, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 47, 2, 'orglevel', 2, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('pay', 'switch', 47, 3, 'orglevel', 3, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 40, 1, 'vicemanagertype', 4, 1, 'forward', '', 45, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 40, 2, 'vicemanagertype', 4, 2, 'forward', '', 42, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 47, 1, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 47, 2, 'orglevel', 2, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 47, 3, 'orglevel', 3, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 33, 1, 'money', 0, 5000, 'forward', '', 40, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 33, 2, 'money', 3, 5000, 'forward', '', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 32, 1, 'money', 0, 5000, 'forward', '', 35, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 32, 2, 'money', 3, 5000, 'forward', '', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 34, 1, 'orglevel', 4, 0, 'forward', '总经理提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 34, 2, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 34, 3, 'orglevel', 4, 2, 'forward', '本部长提交', 40, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 34, 4, 'orglevel', 4, 3, 'forward', '部长提交', 37, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 34, 5, 'orglevel', 0, 4, 'forward', '部长以下提交', 32, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 46, 1, 'money', 0, 10000, 'forward', '', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meeting', 'switch', 46, 2, 'money', 3, 10000, 'forward', '', 47, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('plan', 'switch', 40, 1, 'money', 0, 100000, 'forward', '', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('plan', 'switch', 40, 2, 'money', 3, 100000, 'forward', '', 70, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('plan', 'switch', 85, 1, 'money', 0, 2000000, 'forward', '', 90, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('plan', 'switch', 85, 2, 'money', 3, 2000000, 'forward', '', 999, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 32, 1, 'orglevel', 4, 0, 'forward', '总经理提交', 999, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 32, 2, 'orglevel', 4, 1, 'forward', '副总提交', 60, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 32, 3, 'orglevel', 4, 2, 'forward', '本部长提交', 45, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 32, 4, 'orglevel', 4, 3, 'forward', '部长提交', 35, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 32, 5, 'orglevel', 0, 4, 'forward', '部长以下提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 50, 1, 'scopeflag', 4, 0, 'forward', '国内出差', 999, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('loantravel', 'switch', 50, 2, 'scopeflag', 4, 1, 'forward', '海外出差', 60, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 32, 1, 'orglevel', 4, 0, 'forward', '总经理提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 32, 2, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 32, 3, 'orglevel', 4, 2, 'forward', '本部长提交', 40, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 32, 4, 'orglevel', 4, 3, 'forward', '部长提交', 35, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 32, 5, 'orglevel', 0, 4, 'forward', '部长以下提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 40, 1, 'vicemanagertype', 4, 1, 'forward', '', 45, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 40, 2, 'vicemanagertype', 4, 2, 'forward', '', 42, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 47, 1, 'orglevel', 4, 1, 'forward', '副总提交', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 47, 2, 'orglevel', 2, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('travel', 'switch', 47, 3, 'orglevel', 3, 1, 'forward', '非副总提交', 55, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('overtime', 'switch', 5, 1, 'totalovertimedays', 1, 4.5, 'forward', '', 10, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('overtime', 'switch', 5, 2, 'totalovertimedays', 2, 4.5, 'forward', '', 50, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('overtime', 'switch', 10, 1, 'exceeddays', 1, 7, 'forward', '', 30, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('overtime', 'switch', 10, 2, 'exceeddays', 2, 7, 'forward', '', 40, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meetingroomapply', 'switch', 5, 1, 'schedule', 4, 0, 'forward', '不需要审批', 999, '');
INSERT INTO fi_flowswitchtaskaction_tb VALUES ('meetingroomapply', 'switch', 5, 2, 'schedule', 4, 1, 'forward', '需要审批', 10, '');

INSERT INTO fi_flowstatus VALUES (0, '审批中');
INSERT INTO fi_flowstatus VALUES (1, '已通过');
INSERT INTO fi_flowstatus VALUES (2, '已中止');
INSERT INTO fi_flowstatus VALUES (4, '已取消');

INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 30, 'next', 'forward', '部长审批通过', 32, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 35, 'next', 'forward', '部长审批通过', 40, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 35, 'return', 'backward', '部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 35, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 50, 'next', 'forward', '总经理审批同意', 55, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 60, 'next', 'forward', '财务课长审批通过', 65, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 60, 'return', 'backward', '财务课长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 60, 'stop', 'stop', '财务课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 70, 'next', 'forward', '财务部长审批通过', 75, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 70, 'return', 'backward', '财务部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 70, 'stop', 'stop', '财务部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 75, 'next', 'forward', '出纳审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 75, 'return', 'backward', '出纳审批驳回', 1, 75);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 75, 'stop', 'stop', '出纳审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 80, 'next', 'forward', '出纳发款', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 80, 'return', 'backward', '出纳发款驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 80, 'stop', 'stop', '出纳发款中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 42, 'next', 'forward', '营业副总审批通过', 47, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 42, 'return', 'backward', '营业副总审批驳回', 1, 42);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 42, 'stop', 'stop', '营业副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 45, 'next', 'forward', '售后副总审批通过', 47, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 45, 'return', 'backward', '售后副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 45, 'stop', 'stop', '售后副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 55, 'next', 'forward', '财务担当审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 55, 'return', 'backward', '财务担当审批驳回', 1, 55);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 55, 'stop', 'stop', '财务担当审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 65, 'next', 'forward', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 65, 'return', 'backward', '财务副部长审批', 1, 65);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 65, 'stop', 'stop', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 30, 'next', 'forward', '部长审批通过', 32, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 35, 'next', 'forward', '部长审批通过', 40, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 35, 'return', 'backward', '部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 35, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 42, 'next', 'forward', '营业副总审批通过', 47, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 42, 'return', 'backward', '营业副总审批驳回', 1, 42);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 42, 'stop', 'stop', '营业副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 45, 'next', 'forward', '售后副总审批通过', 47, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 45, 'return', 'backward', '售后副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 45, 'stop', 'stop', '售后副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 50, 'next', 'forward', '总经理审批同意', 55, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 55, 'next', 'forward', '财务担当审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 55, 'return', 'backward', '财务担当审批驳回', 1, 55);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 55, 'stop', 'stop', '财务担当审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 60, 'next', 'forward', '财务课长审批通过', 65, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 60, 'return', 'backward', '财务课长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 60, 'stop', 'stop', '财务课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 65, 'next', 'forward', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 65, 'return', 'backward', '财务副部长审批', 1, 65);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 65, 'stop', 'stop', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 70, 'next', 'forward', '财务部长审批通过', 75, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 70, 'return', 'backward', '财务部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 70, 'stop', 'stop', '财务部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 75, 'next', 'forward', '出纳审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 75, 'return', 'backward', '出纳审批驳回', 1, 75);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 75, 'stop', 'stop', '出纳审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 80, 'next', 'forward', '出纳发款', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 80, 'return', 'backward', '出纳发款驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 80, 'stop', 'stop', '出纳发款中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 30, 'next', 'forward', '部长审批通过', 32, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 80, 'next', 'forward', '出纳发款', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 80, 'return', 'backward', '出纳发款驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 80, 'stop', 'stop', '出纳发款中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 50, 'next', 'forward', '总经理审批同意', 55, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 55, 'next', 'forward', '财务担当审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 55, 'return', 'backward', '财务担当审批驳回', 1, 55);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 55, 'stop', 'stop', '财务担当审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 60, 'next', 'forward', '财务课长审批通过', 65, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 60, 'return', 'backward', '财务课长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 60, 'stop', 'stop', '财务课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 65, 'next', 'forward', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 65, 'return', 'backward', '财务副部长审批', 1, 65);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 65, 'stop', 'stop', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 70, 'next', 'forward', '财务部长审批通过', 75, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 70, 'return', 'backward', '财务部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 70, 'stop', 'stop', '财务部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 75, 'next', 'forward', '出纳审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 75, 'return', 'backward', '出纳审批驳回', 1, 75);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 75, 'stop', 'stop', '出纳审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 80, 'next', 'forward', '出纳发款', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 80, 'return', 'backward', '出纳发款驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 80, 'stop', 'stop', '出纳发款中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 30, 'next', 'forward', '部长审批通过', 34, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 42, 'next', 'forward', '营业副总审批通过', 46, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 42, 'return', 'backward', '营业副总审批驳回', 1, 42);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 42, 'stop', 'stop', '营业副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 45, 'next', 'forward', '售后副总审批通过', 46, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 45, 'return', 'backward', '售后副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 45, 'stop', 'stop', '售后副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 37, 'next', 'forward', '本部长审批通过', 33, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 37, 'return', 'backward', '本部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 37, 'stop', 'stop', '本部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 35, 'next', 'forward', '本部长审批通过', 40, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 35, 'return', 'backward', '本部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 35, 'stop', 'stop', '本部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 50, 'next', 'forward', '总经理审批同意', 55, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 55, 'next', 'forward', '财务担当审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 55, 'return', 'backward', '财务担当审批驳回', 1, 55);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 55, 'stop', 'stop', '财务担当审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 60, 'next', 'forward', '财务课长审批通过', 65, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 60, 'return', 'backward', '财务课长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 60, 'stop', 'stop', '财务课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 70, 'next', 'forward', '财务部长审批通过', 75, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 70, 'return', 'backward', '财务部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 70, 'stop', 'stop', '财务部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 75, 'next', 'forward', '出纳审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 75, 'return', 'backward', '出纳审批驳回', 1, 75);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 75, 'stop', 'stop', '出纳审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 80, 'next', 'forward', '出纳发款', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 80, 'return', 'backward', '出纳发款驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 80, 'stop', 'stop', '出纳发款中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 30, 'next', 'forward', '部长审批通过', 35, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 35, 'next', 'forward', '部长审批通过', 45, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 35, 'return', 'backward', '部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 35, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 45, 'next', 'forward', '副总审批通过', 50, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 45, 'return', 'backward', '副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 45, 'stop', 'stop', '副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 65, 'next', 'forward', '财务副部长审批', 70, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 65, 'return', 'backward', '财务副部长审批', 1, 65);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 65, 'stop', 'stop', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 30, 'next', 'forward', '部长审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loan', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('expense', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 50, 'next', 'forward', '综合企画课长审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 50, 'return', 'backward', '综合企画课长审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 50, 'stop', 'stop', '综合企画课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 60, 'next', 'forward', '综合企画部长审批通过', 70, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 60, 'return', 'backward', '综合企画部长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 60, 'stop', 'stop', '综合企画部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 70, 'next', 'forward', '本部长审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 70, 'return', 'backward', '本部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 70, 'stop', 'stop', '本部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 80, 'next', 'forward', '副总审批通过', 85, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 80, 'return', 'backward', '副总审批驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 80, 'stop', 'stop', '副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 90, 'next', 'forward', '总经理审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 90, 'return', 'backward', '总经理审批驳回', 1, 90);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 90, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 30, 'next', 'forward', '部长审批通过', 40, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('plan', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 35, 'next', 'forward', '部长审批通过', 45, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 35, 'return', 'backward', '部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 35, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 45, 'next', 'forward', '副总审批通过', 50, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 45, 'return', 'backward', '副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 45, 'stop', 'stop', '副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meeting', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('pay', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 30, 'next', 'forward', '部长审批通过', 35, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 35, 'next', 'forward', '部长审批通过', 45, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 35, 'return', 'backward', '部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 35, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 45, 'next', 'forward', '副总审批通过', 50, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 45, 'return', 'backward', '副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 45, 'stop', 'stop', '副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 50, 'next', 'forward', '总经理审批同意', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('common2', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 60, 'next', 'forward', '总经理审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 60, 'return', 'backward', '总经理审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('loantravel', 'man', 60, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 30, 'next', 'forward', '部长审批通过', 32, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 35, 'next', 'forward', '部长审批通过', 40, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 35, 'return', 'backward', '部长审批驳回', 1, 35);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 35, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 42, 'next', 'forward', '营业副总审批通过', 47, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 42, 'return', 'backward', '营业副总审批驳回', 1, 42);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 42, 'stop', 'stop', '营业副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 45, 'next', 'forward', '售后副总审批通过', 47, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 45, 'return', 'backward', '售后副总审批驳回', 1, 45);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 45, 'stop', 'stop', '售后副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 50, 'next', 'forward', '总经理审批同意', 55, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 55, 'next', 'forward', '财务担当审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 55, 'return', 'backward', '财务担当审批驳回', 1, 55);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 55, 'stop', 'stop', '财务担当审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 60, 'next', 'forward', '财务课长审批通过', 65, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 60, 'return', 'backward', '财务课长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 60, 'stop', 'stop', '财务课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 65, 'next', 'forward', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 65, 'return', 'backward', '财务副部长审批', 1, 65);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 65, 'stop', 'stop', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 70, 'next', 'forward', '财务部长审批通过', 75, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 70, 'return', 'backward', '财务部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 70, 'stop', 'stop', '财务部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 75, 'next', 'forward', '出纳审批通过', 80, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 75, 'return', 'backward', '出纳审批驳回', 1, 75);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 75, 'stop', 'stop', '出纳审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 80, 'next', 'forward', '出纳发款', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 80, 'return', 'backward', '出纳发款驳回', 1, 80);
INSERT INTO fi_flowmantaskaction_tb VALUES ('travel', 'man', 80, 'stop', 'stop', '出纳发款中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 10, 'next', 'forward', '代理课长审批通过', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 10, 'return', 'backward', '代理课长审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 10, 'stop', 'stop', '代理课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 15, 'next', 'forward', '课长审批通过', 20, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 20, 'next', 'forward', '副部长审批通过', 25, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 20, 'return', 'backward', '副部长审批驳回', 1, 20);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 20, 'stop', 'stop', '副部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 25, 'next', 'forward', '代理部长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 25, 'return', 'backward', '代理部长审批驳回', 1, 25);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 25, 'stop', 'stop', '代理部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 5, 'next', 'forward', '副课长审批通过', 10, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 5, 'return', 'backward', '副课长审批驳回', 1, 5);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 5, 'stop', 'stop', '副课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 65, 'next', 'forward', '财务副部长审批', 70, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 65, 'return', 'backward', '财务副部长审批', 1, 65);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 65, 'stop', 'stop', '财务副部长审批', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 30, 'next', 'forward', '部长审批通过', 40, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 40, 'next', 'forward', '财务担当审批通过', 50, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 40, 'return', 'backward', '财务担当审批驳回', 1, 55);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 40, 'stop', 'stop', '财务担当审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 50, 'next', 'forward', '出纳审批通过', 60, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 50, 'return', 'backward', '出纳审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 50, 'stop', 'stop', '出纳审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 70, 'next', 'forward', '财务部长审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 70, 'return', 'backward', '财务部长审批驳回', 1, 70);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 70, 'stop', 'stop', '财务部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 60, 'next', 'forward', '财务课长审批通过', 65, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 60, 'return', 'backward', '财务课长审批驳回', 1, 60);
INSERT INTO fi_flowmantaskaction_tb VALUES ('payaccount', 'man', 60, 'stop', 'stop', '财务课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 1, 'submit', 'forward', '待审批', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 15, 'next', 'forward', '课长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 30, 'next', 'forward', '部长审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('att', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 1, 'submit', 'forward', '待审批', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 15, 'next', 'forward', '课长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 30, 'next', 'forward', '部长审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('attcancel', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 30, 'next', 'forward', '部长审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 50, 'next', 'forward', '总经理审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 50, 'return', 'backward', '总经理审批驳回', 1, 50);
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 50, 'stop', 'stop', '总经理审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 40, 'next', 'forward', '副总审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 40, 'return', 'backward', '副总审批驳回', 1, 40);
INSERT INTO fi_flowmantaskaction_tb VALUES ('overtime', 'man', 40, 'stop', 'stop', '副总审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 1, 'submit', 'forward', '待审批', 15, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 15, 'next', 'forward', '课长审批通过', 30, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 15, 'return', 'backward', '课长审批驳回', 1, 15);
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 15, 'stop', 'stop', '课长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 30, 'next', 'forward', '部长审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 30, 'return', 'backward', '部长审批驳回', 1, 30);
INSERT INTO fi_flowmantaskaction_tb VALUES ('resign', 'man', 30, 'stop', 'stop', '部长审批中止', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meetingroomapply', 'man', 1, 'save', 'pause', '', '', '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meetingroomapply', 'man', 1, 'submit', 'forward', '待审批', 5, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meetingroomapply', 'man', 10, 'next', 'forward', '会议室管理者审批通过', 999, '');
INSERT INTO fi_flowmantaskaction_tb VALUES ('meetingroomapply', 'man', 10, 'return', 'backward', '会议室管理者审批驳回', 1, 10);
INSERT INTO fi_flowmantaskaction_tb VALUES ('meetingroomapply', 'man', 10, 'stop', 'stop', '会议室管理者审批中止', '', '');


INSERT INTO cmn_post_tb VALUES (1, '部长');
INSERT INTO cmn_post_tb VALUES (2, '副部长');
INSERT INTO cmn_post_tb VALUES (3, '课长');
INSERT INTO cmn_post_tb VALUES (4, '副课长');
INSERT INTO cmn_post_tb VALUES (5, '代理部长');
INSERT INTO cmn_post_tb VALUES (6, '代理课长');
INSERT INTO cmn_post_tb VALUES (7, '本部长');
INSERT INTO cmn_post_tb VALUES (8, '副总');
INSERT INTO cmn_post_tb VALUES (9, '总经理');



INSERT INTO cmn_group_tb VALUES ('root', '', 'root', -1, null);
INSERT INTO cmn_group_tb VALUES ('chuna', 'root', '出纳', -1, '');
INSERT INTO cmn_group_tb VALUES ('finance_employee', 'root', '财务担当', -1, '');
INSERT INTO cmn_group_tb VALUES ('finance_viceminister', 'root', '财务副部长', -1, '');
INSERT INTO cmn_group_tb VALUES ('finance_minister', 'root', '财务部长', -1, '');
INSERT INTO cmn_group_tb VALUES ('finance_vicedirector', 'root', '财务副课长', -1, '');
INSERT INTO cmn_group_tb VALUES ('finance_director', 'root', '财务课长', -1, '');
INSERT INTO cmn_group_tb VALUES ('gm', 'root', '总经理', -1, '');
INSERT INTO cmn_group_tb VALUES ('chuna1', 'root', '出纳审核组', -1, '');
INSERT INTO cmn_group_tb VALUES ('chuna2', 'root', '出纳发款组', -1, '');
INSERT INTO cmn_group_tb VALUES ('human_director', 'root', '人事课长', -1, '');
INSERT INTO cmn_group_tb VALUES ('human_minister', 'root', '人事部长', -1, '');
INSERT INTO cmn_group_tb VALUES ('admin_minister', 'root', '行政部长', -1, '');
INSERT INTO cmn_group_tb VALUES ('manage_vicegm', 'root', '管理副总', -1, '');
INSERT INTO cmn_group_tb VALUES ('business_vicegm', 'root', '营业副总', -1, '');
INSERT INTO cmn_group_tb VALUES ('businessplan_vicegm', 'root', '营业支援副总', -1, '');
INSERT INTO cmn_group_tb VALUES ('aftersale_vicegm', 'root', '售后部品副总', -1, '');
INSERT INTO cmn_group_tb VALUES ('comprehensiveplan_director', 'root', '综合企画课长', -1, '');
INSERT INTO cmn_group_tb VALUES ('comprehensiveplan_minister', 'root', ' 综合企划部部长组', -1, '');
INSERT INTO cmn_group_tb VALUES ('contract_remind', 'root', '合同提醒', -1, '');
INSERT INTO cmn_group_tb VALUES ('vehicle_director', 'root', '车辆总括课课长', -1, '');
INSERT INTO cmn_group_tb VALUES ('vehicle_minister', 'root', '车辆总括课部长', -1, '');
INSERT INTO cmn_group_tb VALUES ('meetingroomverify', 'root', '会议室申请审批组', -1, '');


INSERT INTO cmn_lang_tb VALUES ('Account', '账号', '账号', 'account', '账号');
INSERT INTO cmn_lang_tb VALUES ('Accountid', '账户编号', '账户编号', 'account id', '账户编号');
INSERT INTO cmn_lang_tb VALUES ('Accounttype', '账户类别', '账户类别', 'account type', '账户类别');
INSERT INTO cmn_lang_tb VALUES ('Action', '动作', '动作', 'action', '动作');
INSERT INTO cmn_lang_tb VALUES ('Address', '地址', '地址', 'address', '地址');
INSERT INTO cmn_lang_tb VALUES ('Agent', '代理人', '代理人', 'agent', '代理人');
INSERT INTO cmn_lang_tb VALUES ('All', '全部', '全部', 'all', '全部');
INSERT INTO cmn_lang_tb VALUES ('Amend', '上午上班结束时间', '上午上班结束时间', 'am end', '上午上班结束时间');
INSERT INTO cmn_lang_tb VALUES ('Amount', '金额', '金额', 'amount', '金额');
INSERT INTO cmn_lang_tb VALUES ('Amstart', '上午上班开始时间', '上午上班开始时间', 'am start', '上午上班开始时间');
INSERT INTO cmn_lang_tb VALUES ('Attachment', '附件', '附件', 'attachment', '附件');
INSERT INTO cmn_lang_tb VALUES ('Authorization', '授权', '授权', 'authorization', '授权');
INSERT INTO cmn_lang_tb VALUES ('Backtask', '后退返回', '后退返回', 'back task', '后退返回');
INSERT INTO cmn_lang_tb VALUES ('Backward', '后退', '后退', 'backward', '后退');
INSERT INTO cmn_lang_tb VALUES ('Balance', '余额', '余额', 'balance money', '余额');
INSERT INTO cmn_lang_tb VALUES ('Bankid', '银行', '银行', 'bank id', '银行');
INSERT INTO cmn_lang_tb VALUES ('Bankname', '银行名称', '银行名称', 'bank name', '银行名称');
INSERT INTO cmn_lang_tb VALUES ('Bankofdeposit', '开户行', '开户行', 'band of deposit', '开户行');
INSERT INTO cmn_lang_tb VALUES ('Beginmonth', '起始月份', '起始月份', 'begin month', '起始月份');
INSERT INTO cmn_lang_tb VALUES ('Birthday', '生日', '生日', 'birthday', '生日');
INSERT INTO cmn_lang_tb VALUES ('Caller', '申请人', '申请人', 'caller', '申请人');
INSERT INTO cmn_lang_tb VALUES ('Cancel', '取消', '取消', 'cancel', '取消');
INSERT INTO cmn_lang_tb VALUES ('Carborrow_qualification', '借车资格', '借车资格', 'car borrow qualification', '借车资格');
INSERT INTO cmn_lang_tb VALUES ('China', '中国', '中国', 'china', '中国');
INSERT INTO cmn_lang_tb VALUES ('Chinese', '简体中文', '简体中文', 'chinese', '简体中文');
INSERT INTO cmn_lang_tb VALUES ('City', '城镇', '城镇', 'city', '城镇');
INSERT INTO cmn_lang_tb VALUES ('Computer', '计算机（配置）', '计算机（配置）', 'computer', '计算机（配置）');
INSERT INTO cmn_lang_tb VALUES ('Computer_cert', '计算机证书', '计算机证书', 'computer cert', '计算机证书');
INSERT INTO cmn_lang_tb VALUES ('Computer_level', '计算机水平', '计算机水平', 'computer level', '计算机水平');
INSERT INTO cmn_lang_tb VALUES ('Concurrent', '执行人并发', '执行人并发', 'concurrent', '执行人并发');
INSERT INTO cmn_lang_tb VALUES ('Conditions', '条件', '条件', 'conditions', '条件');
INSERT INTO cmn_lang_tb VALUES ('Contact_person', '紧急联系人', '紧急联系人', 'contact person', '紧急联系人');
INSERT INTO cmn_lang_tb VALUES ('Contact_way', '紧急联系方式', '紧急联系方式', 'contact way', '紧急联系方式');
INSERT INTO cmn_lang_tb VALUES ('Content', '内容', '内容', 'content', '内容');
INSERT INTO cmn_lang_tb VALUES ('Controldisplayname', '控件显示名称', '控件显示名称', 'control display name', '控件显示名称');
INSERT INTO cmn_lang_tb VALUES ('Controlname', '控件名称', '控件名称', 'control name', '控件名称');
INSERT INTO cmn_lang_tb VALUES ('Controltype', '控件类型', '控件类型', 'control type', '控件类型');
INSERT INTO cmn_lang_tb VALUES ('Countersign', '会签', '会签', 'countersign', '会签');
INSERT INTO cmn_lang_tb VALUES ('Country', '国籍', '国籍', 'country', '国籍');
INSERT INTO cmn_lang_tb VALUES ('Create', '创建', '创建', 'Create', '创建');
INSERT INTO cmn_lang_tb VALUES ('Currencyid', '币种ID', '币种ID', 'currency id', '币种ID');
INSERT INTO cmn_lang_tb VALUES ('Currencyname', '币种名称', '币种名称', 'currency name', '币种名称');
INSERT INTO cmn_lang_tb VALUES ('Currencysymbol', '币种符号', '币种符号', 'currency symbol', '币种符号');
INSERT INTO cmn_lang_tb VALUES ('Datefrom', '开始日期', '开始日期', 'start date', '开始日期');
INSERT INTO cmn_lang_tb VALUES ('Dateto', '结束日期', '结束日期', 'end date', '结束日期');
INSERT INTO cmn_lang_tb VALUES ('Defaultvalue', '缺省值', '缺省值', 'default value', '缺省值');
INSERT INTO cmn_lang_tb VALUES ('Degree', '学历编号', '学历编号', 'degree', '学历编号');
INSERT INTO cmn_lang_tb VALUES ('Degreename', '职级名称', '职级名称', 'degree name', '职级名称');
INSERT INTO cmn_lang_tb VALUES ('Delete', '删除', '删除', 'Delete', '删除');
INSERT INTO cmn_lang_tb VALUES ('Edit', '编辑', '编辑', 'edit', '编辑');
INSERT INTO cmn_lang_tb VALUES ('Editor', '审批者', '审批者', 'editor', '审批者');
INSERT INTO cmn_lang_tb VALUES ('Email', '邮箱', '邮箱', 'email', '邮箱');
INSERT INTO cmn_lang_tb VALUES ('Emergency', '紧急度', '紧急度', 'emergency', '紧急度');
INSERT INTO cmn_lang_tb VALUES ('Employeeid', '唯一号', '唯一号', 'employee id', '唯一号');
INSERT INTO cmn_lang_tb VALUES ('End', '结束', '结束', 'end', '结束');
INSERT INTO cmn_lang_tb VALUES ('Enddate', '结束日期', '结束日期', 'end date', '结束日期');
INSERT INTO cmn_lang_tb VALUES ('Endmonth', '截止月份', '截止月份', 'end month', '截止月份');
INSERT INTO cmn_lang_tb VALUES ('English', '英语', '英语', 'english', '英语');
INSERT INTO cmn_lang_tb VALUES ('English_cert', '英语证书', '英语证书', 'english cert', '英语证书');
INSERT INTO cmn_lang_tb VALUES ('English_level', '英语水平', '英语水平', 'english level', '英语水平');
INSERT INTO cmn_lang_tb VALUES ('English_name', '英文名称', '英文名称', 'english name', '英文名称');
INSERT INTO cmn_lang_tb VALUES ('Entrydate', '入职日期', '入职日期', 'entrydate', '入职日期');
INSERT INTO cmn_lang_tb VALUES ('Expired', '已过期', '已过期', 'expired', '已过期');
INSERT INTO cmn_lang_tb VALUES ('Expireddate', '有效期', '有效期', 'expired date', '有效期');
INSERT INTO cmn_lang_tb VALUES ('Expression', '表达式', '表达式', 'expression', '表达式');
INSERT INTO cmn_lang_tb VALUES ('Ext', '分机号码', '分机号码', 'ext', '分机号码');
INSERT INTO cmn_lang_tb VALUES ('Female', '女', '女', 'female', '女');
INSERT INTO cmn_lang_tb VALUES ('Fiid', '流程编号', '流程编号', 'flow id', '流程编号');
INSERT INTO cmn_lang_tb VALUES ('Filesize', '文件大小', '文件大小', 'file size', '文件大小');
INSERT INTO cmn_lang_tb VALUES ('Filetype', '文件类型', '文件类型', 'file type', '文件类型');
INSERT INTO cmn_lang_tb VALUES ('Flowcontent', '流程内容', '流程内容', 'flow content', '流程内容');
INSERT INTO cmn_lang_tb VALUES ('Flowfinishtime', '流程结束时间', '流程结束时间', 'flow finish time', '流程结束时间');
INSERT INTO cmn_lang_tb VALUES ('Flowid', '流程类型', '流程类型', 'flow id', '流程类型');
INSERT INTO cmn_lang_tb VALUES ('Flowinstid', '流程编号', '流程编号', 'flow instid', '流程编号');
INSERT INTO cmn_lang_tb VALUES ('Flowinstidcol', '流程编号字段', '流程编号字段', 'flow fiid column', '流程编号字段');
INSERT INTO cmn_lang_tb VALUES ('Flowname', '流程名称', '流程名称', 'flow name', '流程名称');
INSERT INTO cmn_lang_tb VALUES ('Flowstarttime', '流程开始时间', '流程开始时间', 'flow start time', '流程开始时间');
INSERT INTO cmn_lang_tb VALUES ('Flowstatus', '流程状态', '流程状态', 'flow status', '流程状态');
INSERT INTO cmn_lang_tb VALUES ('Flowstatuscol', '流程状态字段', '流程状态字段', 'flow status column', '流程状态字段');
INSERT INTO cmn_lang_tb VALUES ('Flowstatusname', '流程状态名称', '流程状态名称', 'flow status name', '流程状态名称');
INSERT INTO cmn_lang_tb VALUES ('Flowtemplateid', '流程模板编号', '流程模板编号', 'flow template id', '流程模板编号');
INSERT INTO cmn_lang_tb VALUES ('Flowtemplatename', '流程模板名称', '流程模板名称', 'flow template name', '流程模板名称');
INSERT INTO cmn_lang_tb VALUES ('Forward', '前进', '前进', 'forward', '前进');
INSERT INTO cmn_lang_tb VALUES ('Friday', '星期五', '星期五', 'friday', '星期五');
INSERT INTO cmn_lang_tb VALUES ('Functions', '函数', '函数', 'function', '函数');
INSERT INTO cmn_lang_tb VALUES ('Good', '良好', '良好', 'good', '良好');
INSERT INTO cmn_lang_tb VALUES ('Grade1', '一级', '一级', 'grade one', '一级');
INSERT INTO cmn_lang_tb VALUES ('Grade2', '二级', '二级', 'grade two', '二级');
INSERT INTO cmn_lang_tb VALUES ('Grade3', '三级', '三级', 'grade three', '三级');
INSERT INTO cmn_lang_tb VALUES ('Grade4', '四级', '四级', 'grade four', '四级');
INSERT INTO cmn_lang_tb VALUES ('Grade6', '六级', '六级', 'grade six', '六级');
INSERT INTO cmn_lang_tb VALUES ('Grade8', '八级', '八级', 'grade eight', '八级');
INSERT INTO cmn_lang_tb VALUES ('Groupid', '组编号', '组编号', 'group id', '组编号');
INSERT INTO cmn_lang_tb VALUES ('Grouplevel', '组级别', '组级别', 'group level', '组级别');
INSERT INTO cmn_lang_tb VALUES ('Groupname', '组名称', '组名称', 'group name', '组名称');
INSERT INTO cmn_lang_tb VALUES ('Guard_card', '门禁卡', '门禁卡', 'guard card', '门禁卡');
INSERT INTO cmn_lang_tb VALUES ('Holiday', '节假日', '节假日', 'holiday', '节假日');
INSERT INTO cmn_lang_tb VALUES ('Idcard', '身份证号码', '身份证号码', 'id card', '身份证号码');
INSERT INTO cmn_lang_tb VALUES ('Isleader', '主管', '主管', 'is leader', '主管');
INSERT INTO cmn_lang_tb VALUES ('Ismultiple', '多选', '多选', 'is multiple', '多选');
INSERT INTO cmn_lang_tb VALUES ('Japanese', '日语', '日语', 'japanese', '日语');
INSERT INTO cmn_lang_tb VALUES ('Japanese_cert', '日语证书', '日语证书', 'japaness cert', '日语证书');
INSERT INTO cmn_lang_tb VALUES ('Japanese_level', '日语水平', '日语水平', 'japanese level', '日语水平');
INSERT INTO cmn_lang_tb VALUES ('Jump', '调度', '调度', 'jump', '调度');
INSERT INTO cmn_lang_tb VALUES ('Key_user', '钥匙', '钥匙', 'key user', '钥匙');
INSERT INTO cmn_lang_tb VALUES ('Langid', '语言ID', '语言ID', 'lang id', '语言ID');
INSERT INTO cmn_lang_tb VALUES ('Leadertype', '主管类型', '主管类型', 'leader type', '主管类型');
INSERT INTO cmn_lang_tb VALUES ('Limit', '限制', '限制', 'limit', '限制');
INSERT INTO cmn_lang_tb VALUES ('Limitfileqty', '限制文件数量', '限制文件数量', 'limit file quantity', '限制文件数量');
INSERT INTO cmn_lang_tb VALUES ('Loginip', '登录IP', '登录IP', 'login ip', '登录IP');
INSERT INTO cmn_lang_tb VALUES ('Major', '专业', '专业', 'major', '专业');
INSERT INTO cmn_lang_tb VALUES ('Male', '男', '男', 'male', '男');
INSERT INTO cmn_lang_tb VALUES ('Manager', '主管', '主管', 'manager', '主管');
INSERT INTO cmn_lang_tb VALUES ('Marital_status', '婚姻状况', '婚姻状况', 'marital status', '婚姻状况');
INSERT INTO cmn_lang_tb VALUES ('Marry', '已婚', '已婚', 'marry', '已婚');
INSERT INTO cmn_lang_tb VALUES ('Maxvalue', '最大值', '最大值', 'max value', '最大值');
INSERT INTO cmn_lang_tb VALUES ('Minvalue', '最小值', '最小值', 'min value', '最小值');
INSERT INTO cmn_lang_tb VALUES ('Mobile', '手机号码', '手机号码', 'mobile', '手机号码');
INSERT INTO cmn_lang_tb VALUES ('Monday', '星期一', '星期一', 'monday', '星期一');
INSERT INTO cmn_lang_tb VALUES ('NO', '序号', '序号', 'no', '序号');
INSERT INTO cmn_lang_tb VALUES ('Native_place', '籍贯', '籍贯', 'native place', '籍贯');
INSERT INTO cmn_lang_tb VALUES ('Next', '同意', '同意', 'next', '同意');
INSERT INTO cmn_lang_tb VALUES ('Nexttask', '下一任务', '下一任务', 'next task', '下一任务');
INSERT INTO cmn_lang_tb VALUES ('Nextyearleave', '可转次年休假', '可转次年休假', 'next year leave', '可转次年休假');
INSERT INTO cmn_lang_tb VALUES ('No', '否', '否', 'no', '否');
INSERT INTO cmn_lang_tb VALUES ('Nopersontask', '无人转向', '无人转向', 'no person task', '无人转向');
INSERT INTO cmn_lang_tb VALUES ('Normal', '一般', '一般', 'normal', '一般');
INSERT INTO cmn_lang_tb VALUES ('Nos', '序号', '序号', 'no', '序号');
INSERT INTO cmn_lang_tb VALUES ('NotMarry', '未婚', '未婚', 'not marry', '未婚');
INSERT INTO cmn_lang_tb VALUES ('Notcity', '非城镇', '非城镇', 'not city', '非城镇');
INSERT INTO cmn_lang_tb VALUES ('Notexpired', '未过期', '未过期', 'not expired', '未过期');
INSERT INTO cmn_lang_tb VALUES ('Nothing', '无', '无', 'nothing', '无');
INSERT INTO cmn_lang_tb VALUES ('Notpassed', '未通过', '未通过', 'not passed', '未通过');
INSERT INTO cmn_lang_tb VALUES ('Offdutytimebutton', '下班打卡', '下班打卡', 'off duty time', '下班打卡');
INSERT INTO cmn_lang_tb VALUES ('Ondutytimebutton', '上班打卡', '上班打卡', 'on duty time', '上班打卡');
INSERT INTO cmn_lang_tb VALUES ('Opnion', '审批意见', '审批意见', 'opnion', '审批意见');
INSERT INTO cmn_lang_tb VALUES ('Orgid', '机构编号', '机构编号', 'organization id', '机构编号');
INSERT INTO cmn_lang_tb VALUES ('Orglevel', '机构级别', '机构级别', 'organization level', '机构级别');
INSERT INTO cmn_lang_tb VALUES ('Orgname', '机构名称', '机构名称', 'organization name', '机构名称');
INSERT INTO cmn_lang_tb VALUES ('Orgtype', '机构类别', '机构类别', 'organization type', '机构类别');
INSERT INTO cmn_lang_tb VALUES ('Orgtypename', '机构类别名称', '机构类别名称', 'organization type name', '机构类别名称');
INSERT INTO cmn_lang_tb VALUES ('Outtype', '请假类型', '请假类型', 'out type', '请假类型');
INSERT INTO cmn_lang_tb VALUES ('Outtypeid', '考勤类型编号', '考勤类型编号', 'out type id', '考勤类型编号');
INSERT INTO cmn_lang_tb VALUES ('Outtypename', '考勤类型名称', '考勤类型名称', 'out type name', '考勤类型名称');
INSERT INTO cmn_lang_tb VALUES ('Parentid', '父编号', '父编号', 'parent id', '父编号');
INSERT INTO cmn_lang_tb VALUES ('Pass', '通过', '通过', 'pass', '通过');
INSERT INTO cmn_lang_tb VALUES ('Passed', '通过', '通过', 'pass', '通过');
INSERT INTO cmn_lang_tb VALUES ('Pause', '暂停', '暂停', 'pause', '暂停');
INSERT INTO cmn_lang_tb VALUES ('Paytype', '支付类型', '支付类型', 'pay type', '支付类型');
INSERT INTO cmn_lang_tb VALUES ('Paytypename', '支付类型名称', '支付类型名称', 'pay type name', '支付类型名称');
INSERT INTO cmn_lang_tb VALUES ('Planid', '企画书ID', '企画书ID', 'plan id', '企画书ID');
INSERT INTO cmn_lang_tb VALUES ('Pmend', '下午上班结束时间', '下午上班结束时间', 'pm end', '下午上班结束时间');
INSERT INTO cmn_lang_tb VALUES ('Pmstart', '下午上班开始时间', '下午上班开始时间', 'pm start', '下午上班开始时间');
INSERT INTO cmn_lang_tb VALUES ('Post', '职位管理', '职位管理', 'post management', '职位管理');
INSERT INTO cmn_lang_tb VALUES ('Postcode', '邮编', '邮编', 'post code', '邮编');
INSERT INTO cmn_lang_tb VALUES ('Postid', '职位编号', '职位编号', 'postid', '职位编号');
INSERT INTO cmn_lang_tb VALUES ('Postname', '职位名称', '职位名称', 'postname', '职位名称');
INSERT INTO cmn_lang_tb VALUES ('Professional_title', '职称', '职称', 'professional title', '职称');
INSERT INTO cmn_lang_tb VALUES ('Pssword', '密码', '密码', 'password', '密码');
INSERT INTO cmn_lang_tb VALUES ('Rank', '职级', '职级', 'rank', '职级');
INSERT INTO cmn_lang_tb VALUES ('Rate', '税率', '税率', 'rate', '税率');
INSERT INTO cmn_lang_tb VALUES ('Remark', '备注', '备注', 'remark', '备注');
INSERT INTO cmn_lang_tb VALUES ('Remarks', '备注', '备注', 'remark', '备注');
INSERT INTO cmn_lang_tb VALUES ('Resetpassword', '重置密码', '重置密码', 'reset password', '重置密码');
INSERT INTO cmn_lang_tb VALUES ('Residence_addres', '户口所在地', '户口所在地', 'residence address', '户口所在地');
INSERT INTO cmn_lang_tb VALUES ('Residence_type', '户口类型', '户口类型', 'residence type', '户口类型');
INSERT INTO cmn_lang_tb VALUES ('Restart', '重启', '重启', 'restart', '重启');
INSERT INTO cmn_lang_tb VALUES ('Return', '驳回', '驳回', 'return', '驳回');
INSERT INTO cmn_lang_tb VALUES ('Rows', '行数', '行数', 'rows', '行数');
INSERT INTO cmn_lang_tb VALUES ('Salaryleave', '有给休假', '有给休假', 'salary leave', '有给休假');
INSERT INTO cmn_lang_tb VALUES ('Samepersontask', '同人转向', '同人转向', 'same person task', '同人转向');
INSERT INTO cmn_lang_tb VALUES ('Saturday', '星期六', '星期六', 'saturday', '星期六');
INSERT INTO cmn_lang_tb VALUES ('Save', '保存', '保存', 'save', '保存');
INSERT INTO cmn_lang_tb VALUES ('School', '毕业院校', '毕业院校', 'school', '毕业院校');
INSERT INTO cmn_lang_tb VALUES ('Sendmessage', '发送消息', '所属部门', 'Sendmessage', '所属部门');
INSERT INTO cmn_lang_tb VALUES ('Sex', '性别', '性别', 'sex', '性别');
INSERT INTO cmn_lang_tb VALUES ('Skip', '跳过', '跳过', 'skip', '跳过');
INSERT INTO cmn_lang_tb VALUES ('Speciality', '特长', '特长', 'speciality', '特长');
INSERT INTO cmn_lang_tb VALUES ('Speciality_cert', '特长证书', '特长证书', 'speciality cert', '特长证书');
INSERT INTO cmn_lang_tb VALUES ('Split', '分流', '分流', 'split flow', '分流');
INSERT INTO cmn_lang_tb VALUES ('Start', '开始', '开始', 'start', '开始');
INSERT INTO cmn_lang_tb VALUES ('Startdate', '开始日期', '开始日期', 'start date', '开始日期');
INSERT INTO cmn_lang_tb VALUES ('Status', '办理状态', '办理状态', 'status', '办理状态');
INSERT INTO cmn_lang_tb VALUES ('Statuss', '办理状态', '办理状态', 'status', '办理状态');
INSERT INTO cmn_lang_tb VALUES ('Stepvalue', '增幅', '增幅', 'step value', '增幅');
INSERT INTO cmn_lang_tb VALUES ('Stop', '中止', '中止', 'stop', '中止');
INSERT INTO cmn_lang_tb VALUES ('Submit', '提交', '提交', 'submit', '提交');
INSERT INTO cmn_lang_tb VALUES ('Sunday', '星期日', '星期日', 'sunday', '星期日');
INSERT INTO cmn_lang_tb VALUES ('Supporskip', '支持跳转', '支持跳转', 'support skip', '支持跳转');
INSERT INTO cmn_lang_tb VALUES ('Supportskip', '支持跳转', '支持跳转', 'Supportskip', '支持跳转');
INSERT INTO cmn_lang_tb VALUES ('Taskexecuter', '执行人', '执行人', 'task executer', '执行人');
INSERT INTO cmn_lang_tb VALUES ('Taskfinishtime', '任务完成时间', '任务完成时间', 'task finish time', '任务完成时间');
INSERT INTO cmn_lang_tb VALUES ('Taskid', '任务ID', '任务ID', 'task id', '任务ID');
INSERT INTO cmn_lang_tb VALUES ('Taskname', '任务名称', '任务名称', 'task name', '任务名称');
INSERT INTO cmn_lang_tb VALUES ('Taskstarttime', '任务开始时间', '任务开始时间', 'task start time', '任务开始时间');
INSERT INTO cmn_lang_tb VALUES ('Taskstatus', '任务状态', '任务状态', 'task status', '任务状态');
INSERT INTO cmn_lang_tb VALUES ('Tchinese', '繁体中文', '繁体中文', 'tchinese', '繁体中文');
INSERT INTO cmn_lang_tb VALUES ('Telphone', '电话号码', '电话号码', 'telphone', '电话号码');
INSERT INTO cmn_lang_tb VALUES ('Thursday', '星期四', '星期四', 'thursday', '星期四');
INSERT INTO cmn_lang_tb VALUES ('Tickets', '发票张数', '发票张数', 'tickets', '发票张数');
INSERT INTO cmn_lang_tb VALUES ('Totalleave', '合计休假', '合计休假', 'total leave', '合计休假');
INSERT INTO cmn_lang_tb VALUES ('Transfer', '转换', '转换', 'transfer', '转换');
INSERT INTO cmn_lang_tb VALUES ('Transferpost', '转岗', '转岗', 'transfer post', '转岗');
INSERT INTO cmn_lang_tb VALUES ('Transfersign', '转签', '转签', 'transfer sign', '转签');
INSERT INTO cmn_lang_tb VALUES ('Transfertype', '转换类型', '转换类型', 'transfer type', '转换类型');
INSERT INTO cmn_lang_tb VALUES ('Transferuserid', '转签、转岗审批人/离职接班人', '转签、转岗审批人/离职接班人', 'transfer useid', '转签、转岗审批人/离职接班人');
INSERT INTO cmn_lang_tb VALUES ('Tuesday', '星期二', '星期二', 'tuesday', '星期二');
INSERT INTO cmn_lang_tb VALUES ('Upload', '上传', '上传', 'Upload', '上传');
INSERT INTO cmn_lang_tb VALUES ('Userid', '用户编号', '用户编号', 'user id', '用户编号');
INSERT INTO cmn_lang_tb VALUES ('Userlevel', '用户级别', '用户级别', 'user level', '用户级别');
INSERT INTO cmn_lang_tb VALUES ('Username', '用户名称', '用户名称', 'user name', '用户名称');
INSERT INTO cmn_lang_tb VALUES ('Usertype', '用户类型', '用户类型', 'user type', '用户类型');
INSERT INTO cmn_lang_tb VALUES ('Valuee', '值', '值', 'value', '值');
INSERT INTO cmn_lang_tb VALUES ('Vary', '变量', '变量', 'vary', '变量');
INSERT INTO cmn_lang_tb VALUES ('Varyname', '变量名称', '变量名称', 'vary name', '变量名称');
INSERT INTO cmn_lang_tb VALUES ('Varytype', '变量类型', '变量类型', 'vary type', '变量类型');
INSERT INTO cmn_lang_tb VALUES ('Varyvalue', '变量值', '变量值', 'vary value', '变量值');
INSERT INTO cmn_lang_tb VALUES ('Verygood', '精通', '精通', 'vary good', '精通');
INSERT INTO cmn_lang_tb VALUES ('ViceManager', '副主管', '副主管', '', '副主管');
INSERT INTO cmn_lang_tb VALUES ('Vicemanager', '副主管', '副主管', 'vice manager', '副主管');
INSERT INTO cmn_lang_tb VALUES ('Vid', '变量', '变量', 'vary', '变量');
INSERT INTO cmn_lang_tb VALUES ('View', '查看', '查看', 'view', '查看');
INSERT INTO cmn_lang_tb VALUES ('Vname', '变量名称', '变量名称', 'vary name', '变量名称');
INSERT INTO cmn_lang_tb VALUES ('Vvalue', '变量值', '变量值', 'vary value', '变量值');
INSERT INTO cmn_lang_tb VALUES ('Wednesday', '星期三', '星期三', 'wednesday', '星期三');
INSERT INTO cmn_lang_tb VALUES ('Weekend', '双休日', '双休日', 'week end', '双休日');
INSERT INTO cmn_lang_tb VALUES ('Welfareleave', '福利休假', '福利休假', 'welfare leave', '福利休假');
INSERT INTO cmn_lang_tb VALUES ('Work_card', '工作牌号', '工作牌号', 'work card', '工作牌号');
INSERT INTO cmn_lang_tb VALUES ('Work_date', '参加工作时间', '参加工作时间', 'work date', '参加工作时间');
INSERT INTO cmn_lang_tb VALUES ('Workday', '工作日', '工作日', 'workday', '工作日');
INSERT INTO cmn_lang_tb VALUES ('Years', '入职年数', '入职年数', 'years', '入职年数');
INSERT INTO cmn_lang_tb VALUES ('Yes', '是', '是', 'yes', '是');
INSERT INTO cmn_lang_tb VALUES ('cancel', '已取消', '已取消', 'cancel', '已取消');
INSERT INTO cmn_lang_tb VALUES ('export', '导出', '导出', 'export', '导出');
INSERT INTO cmn_lang_tb VALUES ('isleader', '是否主管', '是否主管', 'isleader', '是否主管');
INSERT INTO cmn_lang_tb VALUES ('leave', '离职', '离职', 'leave', '离职');
INSERT INTO cmn_lang_tb VALUES ('loadjson', '读取json文件', '读取json文件', 'load json file', '读取json文件');
INSERT INTO cmn_lang_tb VALUES ('next', '待办', '待办', 'to do', '待办');
INSERT INTO cmn_lang_tb VALUES ('orgvary', '机构变量', '机构变量', 'organization vary', '机构变量');
INSERT INTO cmn_lang_tb VALUES ('ownerdepartment', '所属部门', '所属部门', 'ownerdepartment', '所属部门');
INSERT INTO cmn_lang_tb VALUES ('passed', '已通过', '已通过', 'passed', '已通过');
INSERT INTO cmn_lang_tb VALUES ('query', '查询', '查询', 'query', '查询');
INSERT INTO cmn_lang_tb VALUES ('relatedusers', '关联用户', '关联用户', 'related users', '关联用户');
INSERT INTO cmn_lang_tb VALUES ('restart', '重启', '重启', 'restart', '重启');
INSERT INTO cmn_lang_tb VALUES ('save', '保存', '保存', 'save', '保存');
INSERT INTO cmn_lang_tb VALUES ('skip', '跳过', '跳过', 'skip', '跳过');
INSERT INTO cmn_lang_tb VALUES ('stop', '已中止', '已中止', 'stop', '已中止');
INSERT INTO cmn_lang_tb VALUES ('submit', '提交', '提交', 'submit', '提交');
INSERT INTO cmn_lang_tb VALUES ('updatejson', '更新json文件', '更新json文件', 'update json file', '更新json文件');
INSERT INTO cmn_lang_tb VALUES ('upload', '上传', '上传', 'upload', '上传');
INSERT INTO cmn_lang_tb VALUES ('Attdate', '考勤日期', '考勤日期', 'Attdate', '考勤日期');
INSERT INTO cmn_lang_tb VALUES ('Palondutytime', '补签上班时间', '补签上班时间', 'Palondutytime', '补签上班时间');
INSERT INTO cmn_lang_tb VALUES ('Paloffdutytime', '补签下班时间', '补签下班时间', 'Paloffdutytime', '补签下班时间');
INSERT INTO cmn_lang_tb VALUES ('Palreason', '补签原因', '补签下班时间', 'Palreason', '补签下班时间');
INSERT INTO cmn_lang_tb VALUES ('Remainleave', '剩余年假', '剩余年假', 'Remainleave', '剩余年假');
INSERT INTO cmn_lang_tb VALUES ('Usedleave', '已用年假', '已用年假', 'Usedleave', '已用年假');
INSERT INTO cmn_lang_tb VALUES ('Year', '年度', '年度', 'Year', '年度');
INSERT INTO cmn_lang_tb VALUES ('Outdays', '请假天数', '请假天数', 'Outdays', '请假天数');
INSERT INTO cmn_lang_tb VALUES ('Comment', '注释', '注释', 'Comment', '注释');
INSERT INTO cmn_lang_tb VALUES ('Reason', '原因', '原因', 'Reason', '原因');
INSERT INTO cmn_lang_tb VALUES ('Linkman', '联系人', '联系人', 'Linkman', '联系人');
INSERT INTO cmn_lang_tb VALUES ('Calltimefrom', '申请时间FROM', '申请时间FROM', 'Calltimefrom', '申请时间FROM');
INSERT INTO cmn_lang_tb VALUES ('Calltimeto', '申请时间TO', '申请时间TO', 'Calltimeto', '申请时间TO');
INSERT INTO cmn_lang_tb VALUES ('Calltime', '申请时间', '申请时间', 'Calltime', '申请时间');
INSERT INTO cmn_lang_tb VALUES ('Overtimereason', '加班原因', '加班原因', 'Overtimereason', '加班原因');
INSERT INTO cmn_lang_tb VALUES ('Overtimedate', '加班日期', '加班日期', 'Overtimedate', '加班日期');
INSERT INTO cmn_lang_tb VALUES ('Istravel', '是否出差', '是否出差', 'Istravel', '是否出差');
INSERT INTO cmn_lang_tb VALUES ('Leave', '请假', '请假', 'Leave', '请假');
INSERT INTO cmn_lang_tb VALUES ('Ondutytime', '上班时间', '上班时间', 'Ondutytime', '上班时间');
INSERT INTO cmn_lang_tb VALUES ('Offdutytime', '下班时间', '下班时间', 'Offdutytime', '下班时间');
INSERT INTO cmn_lang_tb VALUES ('Attmonth', '考勤月份', '考勤月份', 'Attmonth', '考勤月份');
INSERT INTO cmn_lang_tb VALUES ('Attexception', '考勤异常', '考勤异常', 'Attexception', '考勤异常');
INSERT INTO cmn_lang_tb VALUES ('Earlyleave', '早退', '早退', 'Earlyleave', '早退');
INSERT INTO cmn_lang_tb VALUES ('Late', '迟到', '迟到', 'Late', '迟到');
INSERT INTO cmn_lang_tb VALUES ('Dayovertime', '平日加班时间', '迟到', 'Dayovertime', '迟到');
INSERT INTO cmn_lang_tb VALUES ('Nightovertime', '深夜加班时间', '迟到', 'Nightovertime', '迟到');
INSERT INTO cmn_lang_tb VALUES ('Weekendovertime', '周末加班时间', '迟到', 'Weekendovertime', '迟到');
INSERT INTO cmn_lang_tb VALUES ('Lastweekendovertime', '上月周末加班结余', '上月周末加班结余', 'Lastweekendovertime', '上月周末加班结余');
INSERT INTO cmn_lang_tb VALUES ('Holidayovertime', '节日加班时间', '节日加班时间', 'Holidayovertime', '节日加班时间');
INSERT INTO cmn_lang_tb VALUES ('Overtimefee', '午餐费', '节日加班时间', 'Overtimefee', '节日加班时间');
INSERT INTO cmn_lang_tb VALUES ('Sickleave', '病假', '病假', 'Sickleave', '病假');
INSERT INTO cmn_lang_tb VALUES ('Casualleave', '特别假', '特别假', 'Casualleave', '特别假');
INSERT INTO cmn_lang_tb VALUES ('Lastwelfareleave', '可转次年假', '可转次年假', 'Lastwelfareleave', '可转次年假');
INSERT INTO cmn_lang_tb VALUES ('Reverseleave', '倒休假', '倒休假', 'Reverseleave', '倒休假');
INSERT INTO cmn_lang_tb VALUES ('Medicalleave', 'Medical', 'Medical', 'Medicalleave', 'Medical');
INSERT INTO cmn_lang_tb VALUES ('Marriageleave', '婚假', '婚假', 'Marriageleave', '婚假');
INSERT INTO cmn_lang_tb VALUES ('Checkleave', '事假', '事假', 'Checkleave', '事假');
INSERT INTO cmn_lang_tb VALUES ('Maternityleave', '产假', '产假', 'Maternityleave', '产假');
INSERT INTO cmn_lang_tb VALUES ('Paternityleave', '陪产假', '陪产假', 'Paternityleave', '陪产假');
INSERT INTO cmn_lang_tb VALUES ('Planmaternityleave', '产检假', '产检假', 'Planmaternityleave', '产检假');
INSERT INTO cmn_lang_tb VALUES ('Funeralleave', '丧假', '产假', 'Funeralleave', '产假');
INSERT INTO cmn_lang_tb VALUES ('Officialholiday', '国家法定假', '国家法定假', 'Officialholiday', '国家法定假');
INSERT INTO cmn_lang_tb VALUES ('Hasticket', '是否有发票', '国家法定假', 'Hasticket', '国家法定假');
INSERT INTO cmn_lang_tb VALUES ('Ticketcomment', '发票注释', '发票注释', 'Ticketcomment', '发票注释');
INSERT INTO cmn_lang_tb VALUES ('Attchment', '附件', '发票注释', 'Attchment', '发票注释');
INSERT INTO cmn_lang_tb VALUES ('Claimantfile', '索赔文件', '索赔文件', 'Claimantfile', '索赔文件');
INSERT INTO cmn_lang_tb VALUES ('Payee', '支付对象', '支付对象', 'Payee', '支付对象');
INSERT INTO cmn_lang_tb VALUES ('Bankaccount', '银行账户', '银行账户', 'Bankaccount', '银行账户');
INSERT INTO cmn_lang_tb VALUES ('Bank', '银行', '银行', 'Bank', '银行');
INSERT INTO cmn_lang_tb VALUES ('Branchno', '分行', '分行', 'Branchno', '分行');
INSERT INTO cmn_lang_tb VALUES ('Taxcode', '税号', '税号', 'Taxcode', '税号');
INSERT INTO cmn_lang_tb VALUES ('Address_tel', '地址电话', '地址电话', 'Address_tel', '地址电话');
INSERT INTO cmn_lang_tb VALUES ('Modify', '更改', '更改', 'Modify', '更改');
INSERT INTO cmn_lang_tb VALUES ('Aim', '目的', '目的', 'Aim', '目的');
INSERT INTO cmn_lang_tb VALUES ('Budgetary', '预算措施', '预算措施', 'Budgetary', '预算措施');
INSERT INTO cmn_lang_tb VALUES ('Dotime', '出差日期', '出差日期', 'Dotime', '出差日期');
INSERT INTO cmn_lang_tb VALUES ('Others', '其它', '其它', 'Others', '其它');
INSERT INTO cmn_lang_tb VALUES ('Travelarea', '出差区域', '出差区域', 'Travelarea', '出差区域');
INSERT INTO cmn_lang_tb VALUES ('Customercompany', '对方公司', '对方公司', 'Customercompany', '对方公司');
INSERT INTO cmn_lang_tb VALUES ('Customer', '客户', '客户', 'Customer', '客户');
INSERT INTO cmn_lang_tb VALUES ('Itemcode', '费用编号', '费用编号', 'Itemcode', '费用编号');
INSERT INTO cmn_lang_tb VALUES ('Itemdescription', '费用说明', '费用说明', 'Itemdescription', '费用说明');
INSERT INTO cmn_lang_tb VALUES ('Travelid', '出差申请编号', '出差申请编号', 'Travelid', '出差申请编号');
INSERT INTO cmn_lang_tb VALUES ('Destination', '目的地', '费用说明', 'Destination', '费用说明');
INSERT INTO cmn_lang_tb VALUES ('Dayamount', '每天补助', '每天补助', 'Dayamount', '每天补助');
INSERT INTO cmn_lang_tb VALUES ('Dayamount_description', '每天补助说明', '每天补助说明', 'Dayamount_description', '每天补助说明');
INSERT INTO cmn_lang_tb VALUES ('Itemdate', '项目日期', '项目日期', 'Itemdate', '项目日期');
INSERT INTO cmn_lang_tb VALUES ('Itemtype', '项目类型', '项目日期', 'Itemtype', '项目日期');
INSERT INTO cmn_lang_tb VALUES ('Flytno', '航班号', '航班号', 'Flytno', '航班号');
INSERT INTO cmn_lang_tb VALUES ('Pamount', '预付', '预付', 'Pamount', '预付');
INSERT INTO cmn_lang_tb VALUES ('Camount', '垫付', '垫付', 'Camount', '垫付');
INSERT INTO cmn_lang_tb VALUES ('Line', '线路', '线路', 'Line', '线路');
INSERT INTO cmn_lang_tb VALUES ('Name', '名称', '名称', 'Name', '名称');
INSERT INTO cmn_lang_tb VALUES ('Displayno', '显示顺序', '显示顺序', 'Displayno', '显示顺序');
INSERT INTO cmn_lang_tb VALUES ('Accountype', '账户类型', '账户类型', 'Accountype', '账户类型');
INSERT INTO cmn_lang_tb VALUES ('Openbank', '开户行', '开户行', 'Openbank', '开户行');
INSERT INTO cmn_lang_tb VALUES ('Modualid', '模块编号', '模块编号', 'Modualid', '模块编号');
INSERT INTO cmn_lang_tb VALUES ('Modualname', '模块名称', '模块名称', 'Modualname', '模块名称');
INSERT INTO cmn_lang_tb VALUES ('Tablename', '表名', '表名', 'Tablename', '表名');
INSERT INTO cmn_lang_tb VALUES ('Roleid', '角色ID', '角色ID', 'Roleid', '角色ID');
INSERT INTO cmn_lang_tb VALUES ('Rolename', '角色名称', '角色名称', 'Rolename', '角色名称');
INSERT INTO cmn_lang_tb VALUES ('Rolelevel', '角色级别', '角色级别', 'Rolelevel', '角色级别');
INSERT INTO cmn_lang_tb VALUES ('Userfile', '用户文件', '用户文件', 'Userfile', '用户文件');
INSERT INTO cmn_lang_tb VALUES ('Componentname', '组件名称', '组件名称', 'Componentname', '组件名称');
INSERT INTO cmn_lang_tb VALUES ('Title', '标题', '标题', 'Title', '标题');
INSERT INTO cmn_lang_tb VALUES ('Buttons', '按钮', '按钮', 'Buttons', '按钮');
INSERT INTO cmn_lang_tb VALUES ('Style', '页面风格', '页面风格', 'Style', '页面风格');
INSERT INTO cmn_lang_tb VALUES ('Gutter', '间隔', '间隔', 'Gutter', '间隔');
INSERT INTO cmn_lang_tb VALUES ('Colcount', '列数', '列数', 'Colcount', '列数');
INSERT INTO cmn_lang_tb VALUES ('Componentlevel', '组件级别', '组件级别', 'Componentlevel', '组件级别');
INSERT INTO cmn_lang_tb VALUES ('Godirectory', 'golang项目路径', 'golang项目路径', 'Godirectory', 'golang项目路径');
INSERT INTO cmn_lang_tb VALUES ('Ngdirectory', 'angular项目路径', 'angular项目路径', 'Ngdirectory', 'angular项目路径');
INSERT INTO cmn_lang_tb VALUES ('Islimit', '是否限制', '是否限制', 'Islimit', '是否限制');
INSERT INTO cmn_lang_tb VALUES ('Minvalues', '最小值', '最小值', 'Minvalues', '最小值');
INSERT INTO cmn_lang_tb VALUES ('Maxvalues', '最大值', '最大值', 'Maxvalues', '最大值');
INSERT INTO cmn_lang_tb VALUES ('Icon', '图标', '图标', 'Icon', '图标');
INSERT INTO cmn_lang_tb VALUES ('Password', '密码', '密码', 'Password', '密码');
INSERT INTO cmn_lang_tb VALUES ('Seq', '显示顺序', '显示顺序', 'Seq', '显示顺序');
INSERT INTO cmn_lang_tb VALUES ('Hotel', '住宿地址及名称', '住宿地址及名称', 'Hotel', '住宿地址及名称');
INSERT INTO cmn_lang_tb VALUES ('Flytime', '航班时间', '航班时间', 'Flytime', '航班时间');
INSERT INTO cmn_lang_tb VALUES ('Vehicle', '交通工具', '交通工具', 'Vehicle', '交通工具');
INSERT INTO cmn_lang_tb VALUES ('Idate', '出差日期', '出差日期', 'Idate', '出差日期');
INSERT INTO cmn_lang_tb VALUES ('Medical', '体检', '体检', 'Medical', '体检');
INSERT INTO cmn_lang_tb VALUES ('Paycontent', '支付内容', '支付内容', 'Paycontent', '支付内容');
INSERT INTO cmn_lang_tb VALUES ('Ticketnum', '发票数', '发票数', 'Ticketnum', '发票数');
INSERT INTO cmn_lang_tb VALUES ('expensequerytitle', '费用报销申请查询', '费用报销申请查询', 'expensequerytitle', '费用报销申请查询');
INSERT INTO cmn_lang_tb VALUES ('aftersignattformtitle', '补签考勤', '补签考勤', 'aftersign att', '补签考勤');
INSERT INTO cmn_lang_tb VALUES ('attoutquerytitle', '请假申请查询', '请假申请查询', 'attout query', '请假申请查询');
INSERT INTO cmn_lang_tb VALUES ('attoutformtitle', '请假申请', '请假申请', 'att out apply', '请假申请');
INSERT INTO cmn_lang_tb VALUES ('attoutquerylisttitle', '请假申请查询列表', '请假申请查询列表', 'att out query list', '请假申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('overtimeformtitle', '加班申请', '加班申请', 'overtimeformtitle', '加班申请');
INSERT INTO cmn_lang_tb VALUES ('transfer flow management', '流程转换', '流程转换', 'transfer flow management', '流程转换');
INSERT INTO cmn_lang_tb VALUES ('attformtitle', '请假申请', '请假申请', 'attformtitle', '请假申请');
INSERT INTO cmn_lang_tb VALUES ('attquerytitle', '请假申请查询', '请假申请查询', 'attquerytitle', '请假申请查询');
INSERT INTO cmn_lang_tb VALUES ('attquerylisttitle', '请假申请查询列表', '请假申请查询列表', 'attquerylisttitle', '请假申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('claimantformtitle', '索赔申请', '索赔申请', 'claimantformtitle', '索赔申请');
INSERT INTO cmn_lang_tb VALUES ('claimantimportformtitle', '索赔申请导入', '索赔申请导入', 'claimantimportformtitle', '索赔申请导入');
INSERT INTO cmn_lang_tb VALUES ('claimantquerytitle', '索赔申请查询', '索赔申请查询', 'claimantquerytitle', '索赔申请查询');
INSERT INTO cmn_lang_tb VALUES ('claimantquerylisttitle', '索赔申请查询列表', '索赔申请查询列表', 'claimantquerylisttitle', '索赔申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('payaccountformtitle', '支付账户申请', '支付账户申请', 'payaccountformtitle', '支付账户申请');
INSERT INTO cmn_lang_tb VALUES ('payaccountquerytitle', '支付账户申请查询', '支付账户申请查询', 'payaccountquerytitle', '支付账户申请查询');
INSERT INTO cmn_lang_tb VALUES ('payaccountquerylisttitle', '支付账户申请查询列表', '支付账户申请查询列表', 'payaccountquerylisttitle', '支付账户申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('payformtitle', '支付申请', '支付申请', 'payformtitle', '支付申请');
INSERT INTO cmn_lang_tb VALUES ('payquerytitle', '支付申请查询', '支付申请查询', 'payquerytitle', '支付申请查询');
INSERT INTO cmn_lang_tb VALUES ('payquerylisttitle', '支付申请查询列表', '支付申请查询列表', 'payquerylisttitle', '支付申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('expensequerylisttitle', '费用报销申请查询列表', '费用报销申请查询列表', 'expensequerylisttitle', '费用报销申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('meetingformtitle', '会议费申请', '会议费申请', 'meetingformtitle', '会议费申请');
INSERT INTO cmn_lang_tb VALUES ('meetingquerytitle', '会议费申请查询', '会议费申请查询', 'meetingquerytitle', '会议费申请查询');
INSERT INTO cmn_lang_tb VALUES ('meetingquerylisttitle', '会议费申请查询列表', '会议费申请查询列表', 'meeting query list', '会议费申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('travelformtitle', '差旅费报销申请', '差旅费报销申请', 'travel form', '差旅费报销申请');
INSERT INTO cmn_lang_tb VALUES ('travelquerytitle', '差旅费报销申请查询', '差旅费报销申请查询', 'travel query', '差旅费报销申请查询');
INSERT INTO cmn_lang_tb VALUES ('travelquerylisttitle', '差旅费报销申请查询列表', '差旅费报销申请查询列表', 'travel query list', '差旅费报销申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('loanformtitle', '借款申请', '借款申请', 'loan', '借款申请');
INSERT INTO cmn_lang_tb VALUES ('loanquerytitle', '借款申请查询', '借款申请查询', 'loan query', '借款申请查询');
INSERT INTO cmn_lang_tb VALUES ('loanquerylisttitle', '借款申请查询列表', '借款申请查询列表', 'loan query list', '借款申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('loantravelformtitle', '出差申请', '出差申请', 'loantravel form', '出差申请');
INSERT INTO cmn_lang_tb VALUES ('loantravelquerytitle', '出差申请查询', '出差申请查询', 'loantravel query', '出差申请查询');
INSERT INTO cmn_lang_tb VALUES ('loantravelquerylisttitle', '出差申请查询列表', '出差申请查询列表', 'loantravel query list', '出差申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('planformtitle', '企画申请', '企画申请', 'plan form', '企画申请');
INSERT INTO cmn_lang_tb VALUES ('planquerytitle', '企画申请查询', '企画申请查询', 'plan query', '企画申请查询');
INSERT INTO cmn_lang_tb VALUES ('planquerylisttitle', '企画申请查询列表', '企画申请查询列表', 'plan query list', '企画申请查询列表');
INSERT INTO cmn_lang_tb VALUES ('accounttitle', '银行账户', '银行账户', 'account', '银行账户');
INSERT INTO cmn_lang_tb VALUES ('meetingfeetypetitle', '会议费类型', '会议费类型', 'meeting fee type', '会议费类型');
INSERT INTO cmn_lang_tb VALUES ('langsetting', '多语言', '会议费类型', 'lang setting', '会议费类型');
INSERT INTO cmn_lang_tb VALUES ('annualrule', '年假规则', '年假规则', 'annual rule', '年假规则');
INSERT INTO cmn_lang_tb VALUES ('attsetup', '考勤设置', '考勤设置', 'att setup', '考勤设置');
INSERT INTO cmn_lang_tb VALUES ('holidaysetting', '节假日设置', '节假日设置', 'holidaysetting', '节假日设置');
INSERT INTO cmn_lang_tb VALUES ('outtype', '考勤类型', '考勤类型', 'outtype', '考勤类型');
INSERT INTO cmn_lang_tb VALUES ('usersearch', '用户检索', '用户检索', 'usersearch', '用户检索');
INSERT INTO cmn_lang_tb VALUES ('userlist', '用户列表', '用户列表', 'userlist', '用户列表');
INSERT INTO cmn_lang_tb VALUES ('donetasksearch', '已办任务检索', '已办任务检索', 'donetasksearch', '已办任务检索');
INSERT INTO cmn_lang_tb VALUES ('donetasklist', '已办任务列表', '已办任务列表', 'donetasklist', '已办任务列表');
INSERT INTO cmn_lang_tb VALUES ('flowmonitor', '流程监控', '流程监控', 'flowmonitor', '流程监控');
INSERT INTO cmn_lang_tb VALUES ('flowmonitorlist', '流程监控列表', '流程监控列表', 'flowmonitorlist', '流程监控列表');
INSERT INTO cmn_lang_tb VALUES ('myflow', '我的流程', '我的流程', 'myflow', '我的流程');
INSERT INTO cmn_lang_tb VALUES ('myflowlist', '我的流程列表', '我的流程列表', 'myflowlist', '我的流程列表');
INSERT INTO cmn_lang_tb VALUES ('todotask', '待办任务', '待办任务', 'todotask', '待办任务');
INSERT INTO cmn_lang_tb VALUES ('todotasklist', '待办任务列表', '待办任务列表', 'todotasklist', '待办任务列表');
INSERT INTO cmn_lang_tb VALUES ('mantask', '人工任务', '人工任务', 'mantask', '人工任务');
INSERT INTO cmn_lang_tb VALUES ('action', '动作', '动作', 'action', '动作');
INSERT INTO cmn_lang_tb VALUES ('executor', '执行人', '执行人', 'executor', '执行人');
INSERT INTO cmn_lang_tb VALUES ('floworgvary', '流程机构变量', '流程机构变量', 'floworgvary', '流程机构变量');
INSERT INTO cmn_lang_tb VALUES ('flowstatus', '流程状态', '流程状态', 'flowstatus', '流程状态');
INSERT INTO cmn_lang_tb VALUES ('flowtemplate', '流程模板', '流程状态', 'flowtemplate', '流程状态');
INSERT INTO cmn_lang_tb VALUES ('varies', '变量', '变量', 'varies', '变量');
INSERT INTO cmn_lang_tb VALUES ('transferflowmanagement', '流程转换管理', '流程转换管理', 'transferflowmanagement', '流程转换管理');
INSERT INTO cmn_lang_tb VALUES ('expense', '费用报销申请', '费用报销申请', 'expense', '费用报销申请');
INSERT INTO cmn_lang_tb VALUES ('claimantformlisttitle', '索赔费用列表', '费用报销申请', 'claimantformlisttitle', '费用报销申请');
INSERT INTO cmn_lang_tb VALUES ('meetinglisttitle', '会议费列表', '会议费列表', 'meetinglisttitle', '会议费列表');
INSERT INTO cmn_lang_tb VALUES ('customerlisttitle', '客户列表', '客户列表', 'customerlisttitle', '客户列表');
INSERT INTO cmn_lang_tb VALUES ('detaillisttitle', '会议费明细', '会议费明细', 'detaillisttitle', '会议费明细');
INSERT INTO cmn_lang_tb VALUES ('feelisttitle', '费用列表', '费用列表', 'feelisttitle', '费用列表');
INSERT INTO cmn_lang_tb VALUES ('travellisttitle', '差旅费明细列表', '差旅费明细列表', 'travellisttitle', '差旅费明细列表');
INSERT INTO cmn_lang_tb VALUES ('loanlisttitle', '借款明细列表', '借款明细列表', 'loanlisttitle', '借款明细列表');
INSERT INTO cmn_lang_tb VALUES ('loantravellisttitle', '出差明细列表', '出差明细列表', 'loantravellisttitle', '出差明细列表');
INSERT INTO cmn_lang_tb VALUES ('planlisttitle', '企画明细列表', '企画明细列表', 'planlisttitle', '企画明细列表');
INSERT INTO cmn_lang_tb VALUES ('banksetting', '银行设置', '银行设置', 'banksetting', '银行设置');
INSERT INTO cmn_lang_tb VALUES ('currencysetting', '币种设置', '币种设置', 'currencysetting', '币种设置');
INSERT INTO cmn_lang_tb VALUES ('paytypetitle', '支付类型', '支付类型', 'paytypetitle', '支付类型');
INSERT INTO cmn_lang_tb VALUES ('degreetitle', '学历', '学历', 'degreetitle', '学历');
INSERT INTO cmn_lang_tb VALUES ('orgtypetitle', '机构类型', '机构类型', 'orgtypetitle', '机构类型');
INSERT INTO cmn_lang_tb VALUES ('posttitle', '职位', '职位', 'posttitle', '职位');
INSERT INTO cmn_lang_tb VALUES ('usermanagement', '用户管理', '用户管理', 'usermanagement', '用户管理');
INSERT INTO cmn_lang_tb VALUES ('userimport', '用户导入', '用户导入', 'userimport', '用户导入');
INSERT INTO cmn_lang_tb VALUES ('userinfo', '用户信息', '用户信息', 'userinfo', '用户信息');
INSERT INTO cmn_lang_tb VALUES ('passwordchange', '密码变更', '密码变更', 'passwordchange', '密码变更');
INSERT INTO cmn_lang_tb VALUES ('agentsetting', '代理设置', '代理设置', 'agentsetting', '代理设置');
INSERT INTO cmn_lang_tb VALUES ('switchtask', '分支任务', '分支任务', 'switchtask', '分支任务');
INSERT INTO cmn_lang_tb VALUES ('switch', '分支', '分支', 'switch', '分支');
INSERT INTO cmn_lang_tb VALUES ('modualmanagement', '模块管理', '模块管理', 'modualmanagement', '模块管理');
INSERT INTO cmn_lang_tb VALUES ('orgmanagement', '机构管理', '机构管理', 'orgmanagement', '机构管理');
INSERT INTO cmn_lang_tb VALUES ('orgleader', '机构主管', '机构主管', 'orgleader', '机构主管');
INSERT INTO cmn_lang_tb VALUES ('rolemanagement', '角色管理', '角色管理', 'rolemanagement', '角色管理');
INSERT INTO cmn_lang_tb VALUES ('usergroupmanagement', '用户组管理', '用户组管理', 'usergroupmanagement', '用户组管理');
INSERT INTO cmn_lang_tb VALUES ('Langname', '语言名称', '语言名称', 'Langname', '语言名称');
INSERT INTO cmn_lang_tb VALUES ('user', '用户', '用户', 'user', '用户');
INSERT INTO cmn_lang_tb VALUES ('langsearch', '多语言检索', 'langsearch', 'langsearch', 'langsearch');
INSERT INTO cmn_lang_tb VALUES ('Copy', '复制', '复制', 'Copy', '复制');
INSERT INTO cmn_lang_tb VALUES ('designelmentstitle', '设计画面元素', '设计画面元素', 'designelmentstitle', '设计画面元素');
INSERT INTO cmn_lang_tb VALUES ('designcomponent', '组件设计', '组件设计', 'design component', '组件设计');
INSERT INTO cmn_lang_tb VALUES ('preview', '预览', '预览', 'preview', '预览');
INSERT INTO cmn_lang_tb VALUES ('paylisttitle', '支付明细列表', '支付明细列表', 'paylisttitle', '支付明细列表');
INSERT INTO cmn_lang_tb VALUES ('donetrace', '已办跟踪', '已办跟踪', 'donetrace', '已办跟踪');
INSERT INTO cmn_lang_tb VALUES ('todotrace', '待办跟踪', '待办跟踪', 'todotrace', '待办跟踪');
INSERT INTO cmn_lang_tb VALUES ('print', '打印', '打印', 'print', '打印');
INSERT INTO cmn_lang_tb VALUES ('Opinion', '审批意见', '审批意见', 'Opinion', '审批意见');
INSERT INTO cmn_lang_tb VALUES ('Now', '当前时间', '当前时间', 'Now', '当前时间');
INSERT INTO cmn_lang_tb VALUES ('meetingroomtitle', '会议室管理', '会议室管理', 'meetingroomtitle', '会议室管理');
INSERT INTO cmn_lang_tb VALUES ('Roomid', '会议室编号', '会议室编号', 'Roomid', '会议室编号');
INSERT INTO cmn_lang_tb VALUES ('Roomname', '会议室名称', '会议室名称', 'Roomname', '会议室名称');
INSERT INTO cmn_lang_tb VALUES ('Layer', '所在楼层', '所在楼层', 'Layer', '所在楼层');
INSERT INTO cmn_lang_tb VALUES ('Persons', '容纳人数', '容纳人数', 'Persons', '容纳人数');
INSERT INTO cmn_lang_tb VALUES ('Isvalid', '是否可用', '是否可用', 'Isvalid', '是否可用');
INSERT INTO cmn_lang_tb VALUES ('Equipment', '设备', '设备', 'Equipment', '设备');
INSERT INTO cmn_lang_tb VALUES ('Schedule', '能否预约', '能否预约', 'Schedule', '能否预约');
INSERT INTO cmn_lang_tb VALUES ('supplytypetitle', '物品类型管理', '物品类型管理', 'supplytypetitle', '物品类型管理');
INSERT INTO cmn_lang_tb VALUES ('Mtcode', '物品类型编号', '物品类型编号', 'Mtcode', '物品类型编号');
INSERT INTO cmn_lang_tb VALUES ('Mtname', '物品类型名称', '物品类型名称', 'Mtname', '物品类型名称');
INSERT INTO cmn_lang_tb VALUES ('equipmenttitle', '会议室设备管理', '会议室设备管理', 'equipmenttitle', '会议室设备管理');
INSERT INTO cmn_lang_tb VALUES ('Equipmentcode', '设备编号', '设备编号', 'Equipmentcode', '设备编号');
INSERT INTO cmn_lang_tb VALUES ('Equipmentname', '设备名称', '设备名称', 'Equipmentname', '设备名称');
INSERT INTO cmn_lang_tb VALUES ('unittitle', '单位管理', '单位管理', 'unittitle', '单位管理');
INSERT INTO cmn_lang_tb VALUES ('Unitcode', '单位编号', '单位编号', 'Unitcode', '单位编号');
INSERT INTO cmn_lang_tb VALUES ('Unitname', '单位名称', '单位名称', 'Unitname', '单位名称');
INSERT INTO cmn_lang_tb VALUES ('meetingroomapplyformtitle', '会议室申请', '会议室申请', 'meetingroomapplyformtitle', '会议室申请');
INSERT INTO cmn_lang_tb VALUES ('Meetingtopic', '会议主题', '会议主题', 'Meetingtopic', '会议主题');
INSERT INTO cmn_lang_tb VALUES ('Meetingdate', '会议日期', '会议日期', 'Meetingdate', '会议日期');
INSERT INTO cmn_lang_tb VALUES ('Meetingstarttime', '会议开始时间', '会议开始时间', 'Meetingstarttime', '会议开始时间');
INSERT INTO cmn_lang_tb VALUES ('Meetingendtime', '会议结束时间', '会议结束时间', 'Meetingendtime', '会议结束时间');
INSERT INTO cmn_lang_tb VALUES ('Meetingroom', '会议室', '会议室', 'Meetingroom', '会议室');
INSERT INTO cmn_lang_tb VALUES ('Meetingpersons', '参会人员', '参会人员', 'Meetingpersons', '参会人员');
INSERT INTO cmn_lang_tb VALUES ('Notice', '是否提醒', '是否提醒', 'Notice', '是否提醒');
INSERT INTO cmn_lang_tb VALUES ('Meetingcontent', '会议内容', '会议内容', 'Meetingcontent', '会议内容');

insert into sequence(seqname,currentValue,increment) values('fiid_sequence',201800001,1);
insert into sequence(seqname,currentValue,increment) values('tiid_sequence',1,1);

INSERT INTO cmn_org_tb VALUES (1, '售后管理', 0, '', 1, '', '');
INSERT INTO cmn_org_tb VALUES (11, '售后本部', 1, '', 2, '', '');
INSERT INTO cmn_org_tb VALUES (11101, '售后部代理', 111, '', 4, '', '');
INSERT INTO cmn_org_tb VALUES (11111, '售后课代理', 1111, '', 6, '', '');
INSERT INTO cmn_org_tb VALUES (111, '售后部', 11, '', 3, '', '');
INSERT INTO cmn_org_tb VALUES (2111, '营业管理部代理', 211, '', 4, '', '');
INSERT INTO cmn_org_tb VALUES (211111, '营业管理部课代理', 21111, '', 6, '', '');
INSERT INTO cmn_org_tb VALUES (211, '营业管理部', 21, '', 3, '', '');
INSERT INTO cmn_org_tb VALUES (21, '营业管理本部', 2, '', 2, '', '');
INSERT INTO cmn_org_tb VALUES (21111, '营业管理课', 2111, '', 5, '', '');
INSERT INTO cmn_org_tb VALUES (1111, '售后课', 11101, '', 5, '', '');
INSERT INTO cmn_org_tb VALUES (31, '管理本部', 3, '', 2, '', '');
INSERT INTO cmn_org_tb VALUES (3111, '财务部代理', 311, '', 4, '', '');
INSERT INTO cmn_org_tb VALUES (311111, '财务课代理', 31111, '', 6, '', '');
INSERT INTO cmn_org_tb VALUES (0, '***公司', 'root', '', 0, '', '');
INSERT INTO cmn_org_tb VALUES (311, '财务部', 31, '', 3, '', '');
INSERT INTO cmn_org_tb VALUES (31111, '财务课', 3111, '', 5, '', '');
INSERT INTO cmn_org_tb VALUES (3, '管理部', 0, '', 1, '', '');
INSERT INTO cmn_org_tb VALUES (11112, '索赔课代理', 1112, '', 6, '', '');
INSERT INTO cmn_org_tb VALUES (1112, '索赔课', 11101, '', 5, '', '');
INSERT INTO cmn_org_tb VALUES (2, '营业管理', 0, '', 1, '', '');

INSERT INTO cmn_user_tb VALUES ('hhui', 'hhui', '', 11111, 666666, 0, '2020-09-20 05:11:25+00:00', null, '', null, '', '', '', '', null, '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', null, '', null, '', '', 0, 0, '', '', '', '', '2007-01-01 02:11:04.051+00:00', '', '', '', '', null, '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '');
INSERT INTO cmn_user_tb VALUES ('wjun', 'wjun', null, 1111, 666666, 0, '2020-09-20 05:11:25', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('whyong', 'whyong', null, 111, 666666, 0, '2019-09-29 08:59:18', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 1, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('lxf', 'lxf', null, 1, 666666, 0, '2018-09-19 07:04:51', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 1, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('zhangs', '张三', null, 311, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 1, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('lisi', '李四', null, 31111, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 3, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('zss', '张总经理', null, 0, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 9, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('gfz', '管副总', null, 3, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 8, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('yfz', '营副总', null, 2, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 8, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('lx', '李出纳发', null, 311111, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 0, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('lxq', '冷出纳审', null, 311111, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 0, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('zyd', '张财担当', null, 311111, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 0, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('zyhong', 'zyhong', null, 1112, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 3, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('sp1', 'sp1', null, 11112, 666666, 0, '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 0, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_user_tb VALUES ('zhangl', 'zhangl', null, 11111, 666666, 0, '2022-11-20 08:00:47', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, 3, 1, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);

CREATE TABLE dev_component_tb (
    componentname varchar(255) NOT NULL ,
    parentid varchar(255) NOT NULL DEFAULT '' ,
    title varchar(255) NOT NULL DEFAULT '' ,
    buttons varchar(255),
    style varchar(255),
    gutter varchar(255),
    colcount varchar(255), 
    componentlevel varchar(255), 
    godirectory varchar(255), 
    ngdirectory varchar(255),
  PRIMARY KEY (componentname));

-- ----------------------------
-- Records of dev_component_tb
-- ----------------------------
INSERT INTO  dev_component_tb VALUES ('root', -1, '***项目', '', '', '', '', 0, 'D:\goproject\src\skl-api', 'D:\angular\skl');

INSERT INTO skl_enum_tb VALUES ('orgtype', '机构类型');
INSERT INTO skl_enum_tb VALUES ('usertype', '用户类型');
INSERT INTO skl_enum_tb VALUES ('employeetype', '职员类型');
INSERT INTO skl_enum_tb VALUES ('degree', '学历管理');
INSERT INTO skl_enum_tb VALUES ('post', '职位管理');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('orgtype', 1, '公司');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('orgtype', 2, '部门');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('usertype', 1, '超级用户');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('usertype', 2, '普通用户');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('usertype', 3, '系统管理员');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('employeetype', 1, '本埠员工');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('employeetype', 2, '外埠员工');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('employeetype', 3, '协力公司');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('degree', 10, '博士后');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('degree', 20, '博士');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('degree', 30, '硕士');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('degree', 40, '学士');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('degree', 50, '大专');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('degree', 60, '高中');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 1, '总经理');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 2, '副总经理');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 3, '本部长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 4, '部长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 5, '代理部长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 6, '副部长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 7, '课长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 8, '主查');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 9, '代理课长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 10, '副课长');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 11, '总管');
INSERT INTO skl_enumitem_tb(enumcode,value,label) VALUES ('post', 12, '一般员工');