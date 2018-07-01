drop table accounts if exists
drop table tasks if exists

create table accounts(
  id primary key,
  email varchar unique ,
  token varchar unique ,
  eth_address varchar(42) unique -- 区块链账户地址
);

create table tasks(
  id primary key,
  onwer_id int references accounts(id), --发布者
  template_id int references task_templates(id) --合约模型
  name varchar,
  description varchar,
  reward int, --赏金，单位wei
  address varchar(42) unique -- 区块链合约地址
);

-- 任务合约模板，用来创建合约
create table task_templates(
  id primary key,
  name varchar,
  description varchar,
  abi varchar, --abi
  constructor jsonb
  actions jsonb
)