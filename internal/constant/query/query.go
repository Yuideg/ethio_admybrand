package query
const(
	UserInsert=    `INSERT INTO "users" ("id","first_name","middle_name","last_name","phone","email","user_name","password","gender","address","created_at") VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	// UserInsert = `INSERT INTO users(first_name,password,first_name,last_name,gender,email,country,city,state,zipcode,street,latitude,longitude,phone,role_id) VALUES($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13,$14,$15);`
	UserSelectOne=`SELECT * FROM user_tb WHERE id = $1;`
	UserSelectAll=`SELECT * FROM user_tb;`
	UserUpdate=`UPDATE user_tb SET id=$1, username=$2,password=$3,first_name=$4,last_name=$5,gender=$6,email=$7,country=$8,city=$9,state=$10,zipcode=$11,street=$12,latitude=$13,longitude=$14,phone=$15,role_id=$16 WHERE id=$17;`
	UserDelete=`DELETE FROM user_tb WHERE id=$1;`
)