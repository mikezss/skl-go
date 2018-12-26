DELIMITER $$
CREATE  FUNCTION currval(seq_name VARCHAR(50)) RETURNS int(11)
BEGIN

     DECLARE current INTEGER;

     SET current=0;

     select currentValue into current from sequence where seqname=seq_name;

     return current;

END$$
DELIMITER ;

DELIMITER $$
CREATE  FUNCTION getleadertitle(inputorgid VARCHAR(50),orgname VARCHAR(50)) RETURNS varchar(200) CHARSET utf8
BEGIN
	DECLARE notfound INT DEFAULT 0;    #定义一个辅助变量用于判断
    declare userid varchar(50);
    declare username varchar(50);
    declare leadertitle varchar(200);
    
		declare results cursor for  
    select b.userid,c.username from cmn_org_tb a 
		left join cmn_orgleader_tb b on a.orgid=b.orgid
		inner join cmn_user_tb c on b.userid=c.userid
		where a.orgid=inputorgid;
	DECLARE CONTINUE HANDLER FOR NOT FOUND SET notfound = 1;  #定义declare continue handler
		
		SET leadertitle =CONCAT('(',orgname);
		
		open results;
			if notfound = 1 then 
				return '';
            end if;
    	out_loop:LOOP
    		FETCH results INTO userid,username;
            
            SET leadertitle = CONCAT_ws(',',leadertitle,username);
    		
    	END LOOP out_loop;
    	  SET leadertitle =CONCAT(leadertitle,')');
    CLOSE results;
    RETURN leadertitle;
END$$
DELIMITER ;

DELIMITER $$
CREATE  FUNCTION nextval(seq_name VARCHAR(50)) RETURNS int(11)
BEGIN

    UPDATE sequence  

    SET currentValue = currentValue + increment  

    WHERE seqname = seq_name;    

    RETURN currval(seq_name);


END$$
DELIMITER ;

DELIMITER $$
CREATE FUNCTION setval(seq_name VARCHAR(50), value INTEGER) RETURNS int(11)
BEGIN  
   UPDATE sequence   
   SET currentValue = value  
   WHERE seqname = seq_name;  
   RETURN currval(seq_name);  
END$$
DELIMITER ;
