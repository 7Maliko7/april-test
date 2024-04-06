create schema if not exists car;
create schema if not exists people;

create table if not exists people.list(
	people_id integer PRIMARY KEY,
	people_name	 varchar NOT NULL,
	people_surname	 varchar NOT NULL,
	people_patronymic	varchar
);

create table if not exists car.list(
	car_id integer PRIMARY KEY,
	car_regNum VARCHAR (10) NOT NULL,
	car_mark varchar NOT NULL,
	car_model varchar NOT null,
	car_year integer,
	car_owner_id integer,
	FOREIGN KEY (car_owner_id) REFERENCES people.list (people_id)
);

create sequence if not exists list_seq as integer start with 1 increment by 1;

create schema if not exists operations;

create or replace function operations.add_people(
    pPeople_name varchar,
    pPeople_surnume varchar,
    pPeople_patronymic	varchar
) returns integer
    language plpgsql as
$$
declare
    vResult integer;
begin
	vResult = nextval('public.list_seq'::regclass);
    insert
    into people.list(
     people_id,
	 people_name,
	 people_surname,
	 people_patronymic
    )
    values (vResult,
            pPeople_name,
            pPeople_surnume,
            pPeople_patronymic
           );
    return vResult;
end
$$;

create or replace function operations.add_car(
    pCar_regNum varchar,
	pCar_mark varchar,
	pCar_model varchar,
	pCar_year integer,
	pCar_owner_id integer
) returns integer
    language plpgsql as
$$
declare
    vResult integer;
begin
	vResult = nextval('public.list_seq'::regclass);
    insert
    into car.list(
    car_id,  
    car_regNum,
	car_mark,
	car_model,
	car_year,
	car_owner_id
    )
    values (vResult,
            pCar_regNum,
	        pCar_mark,
	        pCar_model,
	        pCar_year,
	        pCar_owner_id
           );
    return vResult;
end
$$;


create or replace function operations.get_car(pCar_regNum varchar)
    returns TABLE(car_regNum varchar, car_mark varchar, car_model varchar,car_year integer , car_owner_id integer)
    language plpgsql
as
$$
begin
    return query
    select *
    from car.list as cl
    where car_regNum = pCar_regNum;
end
$$;

create or replace function operations.get_car_list(pLimit integer, pOffset integer, pCar_regNum varchar, pCar_mark varchar, pCar_model varchar, pCar_year integer)
    returns TABLE(car_id integer, car_regNum varchar, car_mark varchar, car_model varchar,car_year integer , car_owner_id integer)
    language plpgsql
as
$$
begin
    return query
    select *
    from car.list as cl
    where cl.car_regNum = pCar_regNum and cl.car_mark = pCar_mark and cl.car_model = pCar_model and cl.car_year = pCar_year
    limit pLimit offset pOffset;
  
end
$$;

create or replace function operations.delete_car(pRegNum varchar) 
	returns smallint
    language plpgsql
    as $$
declare
	vCar_id integer;
    vResult smallint;
begin
	select car_id into vCar_id
    from car.list
    where car_regNum = pRegNum;

    delete
    from car.list 
    where car_id = vCar_id;


	with deleted as (
		delete
	    from car.list
	    where car_regNum = pRegNum
	    returning *
	)
	SELECT count(*) FROM deleted into vResult;

    return vResult;
end
$$;


create or replace function operations.update_car(pCar_id integer, pCar_regNum varchar(10), pCar_mark varchar, pCar_model varchar, pCar_year integer, pCar_owner_id integer) 
	returns smallint
    language plpgsql
    as $$
declare
    vResult smallint;
begin
	with updated as (
	    update car.list 
	    set 
	    car_regnum = pCar_regNum,
		car_mark = pCar_mark,
		car_model = pCar_model,
		car_year = pCar_year,
		car_owner_id = pCar_owner_id
	    where car_id = pCar_id
	   returning *           	       	                  
    )
    
    SELECT count(*) FROM updated into vResult;
    return vResult;
end
$$;

