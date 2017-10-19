package main

/*First test to obtain "Master Status" from Remote Host*/

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//'mount' object that will hold MySQL status info
	type mount struct {
		Slave_IO_State                string
		Master_Host                   string
		Master_User                   string
		Master_Port                   string
		Connect_Retry                 string
		Master_Log_File               string
		Read_Master_Log_Pos           string
		Relay_Log_File                string
		Relay_Log_Pos                 string
		Relay_Master_Log_File         string
		Slave_IO_Running              string
		Slave_SQL_Running             string
		Replicate_Do_DB               string
		Replicate_Ignore_DB           string
		Replicate_Do_Table            string
		Replicate_Ignore_Table        string
		Replicate_Wild_Do_Table       string
		Replicate_Wild_Ignore_Table   string
		Last_Errno                    string
		Last_Error                    string
		Skip_Counter                  string
		Exec_Master_Log_Pos           string
		Relay_Log_Space               string
		Until_Condition               string
		Until_Log_File                string
		Until_Log_Pos                 string
		Master_SSL_Allowed            string
		Master_SSL_CA_File            string
		Master_SSL_CA_Path            string
		Master_SSL_Cert               string
		Master_SSL_Cipher             string
		Master_SSL_Key                string
		Seconds_Behind_Master         string
		Master_SSL_Verify_Server_Cert string
		Last_IO_Errno                 string
		Last_IO_Error                 string
		Last_SQL_Errno                string
		Last_SQL_Error                string
		Replicate_Ignore_Server_Ids   string
		Master_Server_Id              string
		Master_UUID                   string
		Master_Info_File              string
		SQL_Delay                     string
		SQL_Remaining_Delay           string
		Slave_SQL_Running_State       string
		Master_Retry_Count            string
		Master_Bind                   string
		Last_IO_Error_Timestamp       string
		Last_SQL_Error_Timestamp      string
		Master_SSL_Crl                string
		Master_SSL_Crlpath            string
		Retrieved_Gtid_Set            string
		Executed_Gtid_Set             string
		Auto_Position                 string
		Replicate_Rewrite_DB          string
		Channel_Name                  string
		Master_TLS_Version            string
	}

	//Create (m)aster Status mount object to mount status info in
	sStatus := new(mount)

	//Connect to MySQL
	db, err := sql.Open("mysql", "XXXX:XXXX@tcp(XXXXXXX:3306)/XXXXTesting")

	if err != nil {
		log.Println(err)
	}

	//Query MySQL and check for errors - defer the closing of the connection until rows have been gathered
	rows, err := db.Query("SHOW SLAVE STATUS")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	///Mount values into 'mount' using 'rows.Scan()'
	for rows.Next() {
		err := rows.Scan(&sStatus.Slave_IO_State, &sStatus.Master_Host, &sStatus.Master_User, &sStatus.Master_Port, &sStatus.Connect_Retry, &sStatus.Master_Log_File, &sStatus.Read_Master_Log_Pos, &sStatus.Relay_Log_File, &sStatus.Relay_Log_Pos, &sStatus.Relay_Master_Log_File, &sStatus.Slave_IO_Running, &sStatus.Slave_SQL_Running, &sStatus.Replicate_Do_DB, &sStatus.Replicate_Ignore_DB, &sStatus.Replicate_Do_Table, &sStatus.Replicate_Ignore_Table, &sStatus.Replicate_Wild_Do_Table, &sStatus.Replicate_Wild_Ignore_Table, &sStatus.Last_Errno, &sStatus.Last_Error, &sStatus.Skip_Counter, &sStatus.Exec_Master_Log_Pos, &sStatus.Relay_Log_Space, &sStatus.Until_Condition, &sStatus.Until_Log_File, &sStatus.Until_Log_Pos, &sStatus.Master_SSL_Allowed, &sStatus.Master_SSL_CA_File, &sStatus.Master_SSL_CA_Path, &sStatus.Master_SSL_Cert, &sStatus.Master_SSL_Cipher, &sStatus.Master_SSL_Key, &sStatus.Seconds_Behind_Master, &sStatus.Master_SSL_Verify_Server_Cert, &sStatus.Last_IO_Errno, &sStatus.Last_IO_Error, &sStatus.Last_SQL_Errno, &sStatus.Last_SQL_Error, &sStatus.Replicate_Ignore_Server_Ids, &sStatus.Master_Server_Id, &sStatus.Master_UUID, &sStatus.Master_Info_File, &sStatus.SQL_Delay, &sStatus.SQL_Remaining_Delay, &sStatus.Slave_SQL_Running_State, &sStatus.Master_Retry_Count, &sStatus.Master_Bind, &sStatus.Last_IO_Error_Timestamp, &sStatus.Last_SQL_Error_Timestamp, &sStatus.Master_SSL_Crl, &sStatus.Master_SSL_Crlpath, &sStatus.Retrieved_Gtid_Set, &sStatus.Executed_Gtid_Set, &sStatus.Auto_Position, &sStatus.Replicate_Rewrite_DB, &sStatus.Channel_Name, &sStatus.Master_TLS_Version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(sStatus.Master_Host)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
