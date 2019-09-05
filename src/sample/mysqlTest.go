package sample

/**
 * import时调用package的init函数，自动将db driver注册到sql包中：
 * 		sql.Register("mysql", &MySQLDriver())
 * 调用sql.Open() 返回db 对象Cononnection
 *
 * 1. sql.Open() 返回sql.DB类型db。该对象是线程安全的，且内部已经包含了线程池
 *    open 函数并没有创建连接，它只是验证参数是否合法。然后开启一个单独goroutines去监听是否需要建立新的连接，当有请求建立新连接时就创建新连接。
 *    注意：open函数应该被调用一次，通常是没必要close的。
 * 		db.setMaxOpenConns()
 * 		db.Exec()
 * 		db.Query() / db.QueryRow()
 *
 * 2. Statement 执行计划
 * 		db.Prepare()
 * 		stm.Exec() / stm.Query()
 * 		stm.Close()
 *
 * 2. 事务。Begin()返回Tx对象。TX就和指定的连接绑定在一起了。一旦事务提交或者回滚，该事务绑定的连接就还给DB的连接池。
 * 		db.Begin()
 * 		tx.Exec() / Query()
 * 		tx.Commit()
 * 		tx.Rolllback()
 * 		tx.Stmt() //用于将一个已存在的statement和tx绑定在一起
 *
 * 3. sql.Rows
 * 		Next()
 * 		Scan()
 * 		Err()
 *
 * 使用一个事务或者Statement，多次db操作后，一次性提交性能最好
 *
 * */

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //为了执行包的Init
)

const sqlStr = "select gatewayId,policyCheckSum from recent_changes_policy;"

// MysqlTest : test mysql driver
func MysqlTest() {
	db, err := sql.Open("mysql", "admin:cloudmailscan#1.@tcp(10.206.66.55:3306)/ceagentdb?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var (
		id       string
		checksum string
	)

	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer rows.Close()

	/*
		stm, _ := db.Prepare(sqlStr)
		rows, _ := stm.Query()
		...
		rows.Close()
		stm.Close()
	*/

	/*
		tx, _ := db.Begin()
		rows, _ := tx.Query(sqlStr)
		tx.Commit()
		rows.Close()
	*/

	// 必须把rows中数据读完，或者显式调用Close()方法。否则 defer rows.Close()之前，连接会一直hold
	for rows.Next() {
		err := rows.Scan(&id, &checksum)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Printf("%s , %s\n", id, checksum)
	}
	err = rows.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
