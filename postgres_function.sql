CREATE OR REPLACE FUNCTION currval(seq_name character varying)
  RETURNS bigint AS
$BODY$
DECLARE current INTEGER;
BEGIN 
     current=0;

     select currentValue into current from sequence where seqname=seq_name;

     return current;

END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

 
CREATE  FUNCTION getleadertitle(inputorgid VARCHAR(50),orgname VARCHAR(50)) RETURNS varchar(200)  as 
    $BODY$
    declare notfound INT DEFAULT 0;    --定义一个辅助变量用于判断
    declare userid varchar(50);
    declare username varchar(50);
    declare leadertitle varchar(200);
    
    declare results cursor for  
    select b.userid,c.username from cmn_org_tb a 
		left join cmn_orgleader_tb b on a.orgid=b.orgid
		inner join cmn_user_tb c on b.userid=c.userid
		where a.orgid=inputorgid;
     
 
BEGIN
	  
		
		 leadertitle =CONCAT('(',orgname);
		
		open results;
			if notfound = 1 then 
				return '';
            end if;
    	 LOOP
    		FETCH results INTO userid,username;
            
             leadertitle = CONCAT_ws(',',leadertitle,username);
    		
    	END LOOP  ;
    	   leadertitle =CONCAT(leadertitle,')');
    CLOSE results;
    RETURN leadertitle;
END;
$BODY$
LANGUAGE plpgsql VOLATILE
  COST 100;

 
CREATE  FUNCTION nextval(seq_name VARCHAR(50)) RETURNS bigint as 
$body$
BEGIN

    UPDATE sequence  

    SET currentValue = currentValue + increment  

    WHERE seqname = seq_name;    

    RETURN currval(seq_name);


END ;
$body$
LANGUAGE plpgsql VOLATILE
  COST 100;

CREATE FUNCTION setval(seq_name VARCHAR(50), value INTEGER) RETURNS bigint as 
$body$
BEGIN  
   UPDATE sequence   
   SET currentValue = value  
   WHERE seqname = seq_name;  
   RETURN currval(seq_name);  
END ;
$body$
LANGUAGE plpgsql VOLATILE
  COST 100;
  
CREATE AGGREGATE group_concat(anyelement)
(
    sfunc = array_append, -- 每行的操作函数，将本行append到数组里 
    stype = anyarray,  -- 聚集后返回数组类型 
    initcond = '{}'    -- 初始化空数组
);