package postgres

const (
	AddCarQuery          = `/*NO LOAD BALANCE*/ select * from operations.add_car($1, $2, $3, $4::integer, $5::integer) as id;`
	AddPeopleQuery = `/*NO LOAD BALANCE*/ select * from operations.add_people($1, $2, $3) as id;`
	GetCarQuery          = `select * from operations.get_car($1);`
	DeleteCarQuery       = `/*NO LOAD BALANCE*/ select * from operations.delete_car($1) as count;`
	UpdateCarQuery       = `/*NO LOAD BALANCE*/ select * from operations.update_car($1::integer, $2, $3, $4, $5::integer, $6::integer) as count;`
	GetCarListQuery         = `select * from operations.get_car_list($1::integer, $2::integer, $3, $4, $5, $6::integer);`
)
